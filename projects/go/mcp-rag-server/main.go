// MCP RAG 知识库检索服务（Streamable HTTP 模式）
//
// 在基础计算器之上，新增 knowledge_base_search 工具，
// 实现 RAG（检索增强生成）的标准管线：
//
//	用户提问 → Embedding 向量化 → 向量相似度检索 → 返回 Top K 相关文档
//
// 向量化支持两种实现，通过环境变量 EMBEDDER 切换：
//   - simple（默认）: 基于字符 bigram 的本地向量化，零依赖，开箱即用
//   - ollama: 调用 Ollama 的 nomic-embed-text 模型，语义理解更准确
//
// 客户端配置示例：
//
//	{
//	  "mcpServers": {
//	    "rag-server": { "url": "http://localhost:8080/mcp" }
//	  }
//	}
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ============================================================================
// 第一部分：知识库数据结构
// ============================================================================

// KnowledgeDoc 表示知识库中的一篇文档。
type KnowledgeDoc struct {
	ID      int    // 文档唯一标识
	Title   string // 文档标题
	Content string // 文档正文
	Source  string // 来源说明（如 "员工手册 v3.2"）
}

// ============================================================================
// 第二部分：向量化接口与实现
// ============================================================================

// Embedder 定义向量化接口。
// 输入文本，输出一个固定（或可变）维度的浮点数向量。
type Embedder interface {
	Embed(ctx context.Context, text string) ([]float64, error)
	Dim() int // 向量维度
}

// SimpleEmbedder 基于字符 bigram 的本地向量化实现。
//
// 原理：
//  1. 从所有文档中提取字符 bigram，构建词汇表（vocabulary）
//  2. 每篇文档表示为一个词频向量（BoW），维度 = 词汇表大小
//  3. 向量做 L2 归一化，使得余弦相似度退化为向量点积
//
// 这种方法的优点是完全本地运行，无需任何外部依赖；
// 缺点是缺乏语义理解——"电脑"和"计算机"会被视为完全不同的 bigram。
type SimpleEmbedder struct {
	vocab map[string]int // bigram → 向量维度索引
}

// NewSimpleEmbedder 从已有文档集构建词汇表。
func NewSimpleEmbedder(docs []KnowledgeDoc) *SimpleEmbedder {
	e := &SimpleEmbedder{vocab: make(map[string]int)}
	for _, doc := range docs {
		for _, bg := range extractBigrams(doc.Content) {
			if _, ok := e.vocab[bg]; !ok {
				e.vocab[bg] = len(e.vocab)
			}
		}
		for _, bg := range extractBigrams(doc.Title) {
			if _, ok := e.vocab[bg]; !ok {
				e.vocab[bg] = len(e.vocab)
			}
		}
	}
	slog.Info("simple embedder built", "vocab_size", len(e.vocab))
	return e
}

func (e *SimpleEmbedder) Dim() int { return len(e.vocab) }

// Embed 将文本转为词频向量并做 L2 归一化。
func (e *SimpleEmbedder) Embed(ctx context.Context, text string) ([]float64, error) {
	vec := make([]float64, len(e.vocab))
	bigrams := extractBigrams(text)
	for _, bg := range bigrams {
		if idx, ok := e.vocab[bg]; ok {
			vec[idx]++
		}
	}
	// L2 归一化
	norm := 0.0
	for _, v := range vec {
		norm += v * v
	}
	if norm > 0 {
		norm = math.Sqrt(norm)
		for i := range vec {
			vec[i] /= norm
		}
	}
	return vec, nil
}

// extractBigrams 从中文文本中提取字符级别的 bigram。
// 忽略空白和标点符号，只保留汉字组成的 bigram。
func extractBigrams(text string) []string {
	runes := make([]rune, 0, len([]rune(text)))
	for _, r := range text {
		if unicode.Is(unicode.Han, r) {
			runes = append(runes, r)
		}
	}
	if len(runes) < 2 {
		return nil
	}
	bigrams := make([]string, 0, len(runes)-1)
	for i := 0; i < len(runes)-1; i++ {
		bigrams = append(bigrams, string(runes[i:i+2]))
	}
	return bigrams
}

// OllamaEmbedder 通过 Ollama REST API 获取真实语义向量。
//
// 前置条件：
//  1. 安装 Ollama: https://ollama.com/download/windows
//  2. 拉取模型: ollama pull nomic-embed-text
//
// 启动时设置环境变量: EMBEDDER=ollama
type OllamaEmbedder struct {
	baseURL string
	model   string
	dim     int
}

func NewOllamaEmbedder(baseURL, model string) *OllamaEmbedder {
	return &OllamaEmbedder{
		baseURL: baseURL,
		model:   model,
		dim:     768, // nomic-embed-text 输出 768 维
	}
}

func (e *OllamaEmbedder) Dim() int { return e.dim }

func (e *OllamaEmbedder) Embed(ctx context.Context, text string) ([]float64, error) {
	body := map[string]interface{}{
		"model":  e.model,
		"input":  []string{text},
		"truncate": true,
	}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(ctx, "POST", e.baseURL+"/api/embed", bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("build ollama request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ollama embed failed: %w\nhint: is ollama running? try 'ollama serve'", err)
	}
	defer resp.Body.Close()

	var result struct {
		Embeddings [][]float64 `json:"embeddings"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode ollama response: %w", err)
	}
	if len(result.Embeddings) == 0 {
		return nil, fmt.Errorf("ollama returned empty embeddings")
	}
	return result.Embeddings[0], nil
}

// ============================================================================
// 第三部分：向量存储与相似度检索
// ============================================================================

// VectorStore 内存向量库，存储文档及其对应的向量。
type VectorStore struct {
	docs    []KnowledgeDoc
	vectors [][]float64
}

// Add 向向量库中添加一篇文档及其向量。
func (vs *VectorStore) Add(doc KnowledgeDoc, vec []float64) {
	vs.docs = append(vs.docs, doc)
	vs.vectors = append(vs.vectors, vec)
}

// Search 基于余弦相似度检索与 queryVec 最相似的 topK 篇文档。
// 返回按相似度降序排列的文档切片。
func (vs *VectorStore) Search(queryVec []float64, topK int) []KnowledgeDoc {
	if topK <= 0 || topK > len(vs.docs) {
		topK = len(vs.docs)
	}

	type scoredDoc struct {
		doc   KnowledgeDoc
		score float64
	}
	results := make([]scoredDoc, len(vs.docs))

	for i, vec := range vs.vectors {
		results[i] = scoredDoc{
			doc:   vs.docs[i],
			score: cosineSimilarity(queryVec, vec),
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})

	docs := make([]KnowledgeDoc, topK)
	for i := 0; i < topK; i++ {
		docs[i] = results[i].doc
	}
	return docs
}

// cosineSimilarity 计算两个 L2 归一化向量的余弦相似度。
// 如果向量已经 L2 归一化，就是直接的点积。
func cosineSimilarity(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0
	}
	var dot, normA, normB float64
	for i := range a {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dot / (math.Sqrt(normA) * math.Sqrt(normB))
}

// ============================================================================
// 第四部分：示例知识库
// ============================================================================

// sampleKnowledgeBase 返回示例知识库文档。
// 实际使用时替换为你们企业内部的真实文档即可。
func sampleKnowledgeBase() []KnowledgeDoc {
	return []KnowledgeDoc{
		{
			ID:      1,
			Title:   "员工考勤管理制度",
			Content: "公司实行弹性工作制，核心工作时间为上午十点到下午四点。员工每日需在考勤系统完成签到和签退，每月累计迟到超过五次将影响当月全勤奖。远程办公需提前一天在OA系统中提交申请，经直属上级审批后方可执行。",
			Source:   "员工手册 v3.2",
		},
		{
			ID:      2,
			Title:   "差旅报销流程说明",
			Content: "差旅报销需在出差结束后的五个工作日内提交报销申请，需附上交通票据、住宿发票及出差审批单。机票需通过公司指定的差旅平台预订，私自购买不予报销。住宿标准为一线城市每晚不超过500元，其他城市每晚不超过350元。餐补按出差天数计算，每人每天100元。",
			Source:   "财务制度手册",
		},
		{
			ID:      3,
			Title:   "产品功能介绍：智能客服系统",
			Content: "智能客服系统支持多渠道接入，包括网页、微信小程序和APP。核心功能包括：自动回复常见问题、人工客服转接、会话记录查询、客户满意度评价。系统采用自然语言处理技术，能够理解用户意图并匹配知识库中的答案，平均响应时间小于2秒。",
			Source:   "产品白皮书 2025",
		},
		{
			ID:      4,
			Title:   "服务器部署架构说明",
			Content: "生产环境采用Kubernetes集群部署，共有三个节点组成高可用集群。数据库使用MySQL主从架构，主库负责写入，两个从库负责读取查询。Redis集群用于缓存用户会话和热点数据。所有服务均部署在阿里云华北二区，通过SLB负载均衡对外提供服务。日志采集使用ELK技术栈，监控告警使用Prometheus配合Grafana。",
			Source:   "运维手册 v2.1",
		},
		{
			ID:      5,
			Title:   "新员工入职指南",
			Content: "入职第一天请携带身份证和银行卡到行政部办理入职手续，领取办公设备和门禁卡。IT部门会协助安装必要的开发工具和配置公司邮箱。入职第一周需完成新员工培训课程，包括公司文化、安全规范和业务流程。试用期为三个月，试用期满后需提交转正申请并参加转正答辩。",
			Source:   "人力资源部",
		},
		{
			ID:      6,
			Title:   "代码评审规范",
			Content: "所有代码提交前必须通过至少一位高级工程师的代码评审。评审重点包括：代码逻辑正确性、命名规范、异常处理、单元测试覆盖率和安全性检查。每个PR的变更行数建议不超过500行，超过需拆分为多个PR提交。评审意见需在24小时内处理，紧急修复可以走快速通道但需事后补评审。",
			Source:   "技术规范 v4.0",
		},
		{
			ID:      7,
			Title:   "数据安全管理制度",
			Content: "所有生产数据禁止下载到本地设备，开发测试需使用脱敏数据。数据库密码需每季度更换一次，密码长度不少于16位且包含大小写字母、数字和特殊字符。发现安全漏洞请立即上报安全部门，不得自行对外披露。员工离职前需完成数据清理确认和权限回收。",
			Source:   "信息安全制度",
		},
	}
}

// ============================================================================
// 第五部分：工具入参定义
// ============================================================================

type addArgs struct {
	A float64 `json:"a" jsonschema:"第一个数字"`
	B float64 `json:"b" jsonschema:"第二个数字"`
}

type subArgs struct {
	A float64 `json:"a" jsonschema:"被减数"`
	B float64 `json:"b" jsonschema:"减数"`
}

type searchArgs struct {
	Query string `json:"query" jsonschema:"用户的提问内容"`
	TopK  int    `json:"top_k" jsonschema:"返回的相关文档数量，默认3"`
}

// ============================================================================
// 第六部分：主函数
// ============================================================================

func main() {
	// 初始化日志
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	slog.Info("initializing MCP RAG knowledge base server",
		"name", "rag-knowledge-server",
		"version", "1.0.0",
		"transport", "streamable-http",
	)

	// 加载知识库
	docs := sampleKnowledgeBase()
	slog.Info("knowledge base loaded", "doc_count", len(docs))

	// 选择向量化引擎
	var embedder Embedder
	embedderType := os.Getenv("EMBEDDER")
	switch embedderType {
	case "ollama":
		embedder = NewOllamaEmbedder("http://localhost:11434", "nomic-embed-text")
		slog.Info("using Ollama embedder", "url", "http://localhost:11434", "model", "nomic-embed-text")
	default:
		embedder = NewSimpleEmbedder(docs)
		slog.Info("using simple bigram embedder (set EMBEDDER=ollama for semantic search)")
	}

	// 将知识库文档全部向量化并存入向量库
	store := &VectorStore{}
	slog.Info("indexing documents...")
	for _, doc := range docs {
		vec, err := embedder.Embed(context.Background(), doc.Title+" "+doc.Content)
		if err != nil {
			slog.Error("failed to embed document", "doc_id", doc.ID, "title", doc.Title, "error", err)
			os.Exit(1)
		}
		store.Add(doc, vec)
	}
	slog.Info("indexing complete", "doc_count", len(store.docs), "vector_dim", embedder.Dim())

	// 创建 MCP 服务
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "rag-knowledge-server",
		Version: "1.0.0",
	}, &mcp.ServerOptions{
		Logger: logger,
	})

	// 注册加法工具
	mcp.AddTool(server, &mcp.Tool{
		Name:        "calculator_add",
		Description: "计算两个数字的加法，返回 a + b 的结果",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args addArgs) (*mcp.CallToolResult, any, error) {
		slog.Debug("calculator_add invoked", "a", args.A, "b", args.B)
		result := args.A + args.B
		slog.Info("calculator_add completed", "a", args.A, "b", args.B, "result", result)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("加法计算结果：%.2f", result)},
			},
		}, nil, nil
	})

	// 注册减法工具
	mcp.AddTool(server, &mcp.Tool{
		Name:        "calculator_sub",
		Description: "计算两个数字的减法，返回 a - b 的结果",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args subArgs) (*mcp.CallToolResult, any, error) {
		slog.Debug("calculator_sub invoked", "a", args.A, "b", args.B)
		result := args.A - args.B
		slog.Info("calculator_sub completed", "a", args.A, "b", args.B, "result", result)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("减法计算结果：%.2f", result)},
			},
		}, nil, nil
	})

	// 注册知识库检索工具 —— RAG 核心
	mcp.AddTool(server, &mcp.Tool{
		Name: "knowledge_base_search",
		Description: "检索企业内部知识库，返回与用户问题最相关的文档片段。" +
			"适用于需要查找公司制度、产品功能、技术架构等内部信息的场景。",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args searchArgs) (*mcp.CallToolResult, any, error) {
		// 参数默认值处理
		query := strings.TrimSpace(args.Query)
		if query == "" {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: "请提供检索关键词或问题描述。"},
				},
			}, nil, nil
		}
		topK := args.TopK
		if topK <= 0 {
			topK = 3
		}

		slog.Info("knowledge_base_search invoked", "query", query, "top_k", topK)

		// Step 1: 将用户问题转为向量
		queryVec, err := embedder.Embed(ctx, query)
		if err != nil {
			slog.Error("embedding failed", "query", query, "error", err)
			return nil, nil, fmt.Errorf("向量化失败：%w", err)
		}

		// Step 2: 从向量库中检索 Top K 相关文档
		results := store.Search(queryVec, topK)

		// Step 3: 格式化检索结果
		var sb strings.Builder
		sb.WriteString("知识库检索到的相关内容：\n\n")
		for i, doc := range results {
			sb.WriteString(fmt.Sprintf("【%d】%s\n", i+1, doc.Title))
			sb.WriteString(fmt.Sprintf("内容：%s\n", doc.Content))
			sb.WriteString(fmt.Sprintf("来源：%s\n\n", doc.Source))
		}
		sb.WriteString("---\n请基于以上检索到的文档内容回答用户的问题。如果文档内容不足以回答问题，请如实告知用户。")

		resultText := sb.String()
		slog.Info("knowledge_base_search completed", "query", query, "results_count", len(results))

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: resultText},
			},
		}, nil, nil
	})

	// 创建 HTTP 处理器并启动服务
	handler := mcp.NewStreamableHTTPHandler(func(r *http.Request) *mcp.Server {
		return server
	}, nil)

	http.HandleFunc("/mcp", handler.ServeHTTP)

	slog.Info("HTTP server starting, listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("HTTP server failed", "error", err)
		os.Exit(1)
	}
}
