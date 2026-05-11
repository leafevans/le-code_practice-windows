// MCP 计算器服务（HTTP/SSE 传输模式）
//
// MCP 协议支持三种传输方式：
//   - stdio: 客户端拉起子进程，通过标准输入/输出通信（适合本地桌面客户端）
//   - HTTP/SSE: 服务端作为独立 HTTP 进程运行，请求走 HTTP，流式响应走 SSE（适合远程调用、多客户端）
//   - Streamable HTTP: HTTP/SSE 的升级版（2025-03-26 规范），统一为单个 /mcp 端点
//
// 本服务使用 Streamable HTTP 模式，暴露两个工具：
//   - calculator_add: 加法运算
//   - calculator_sub: 减法运算
//
// 客户端配置示例（Claude Desktop claude_desktop_config.json）：
//
//	{
//	  "mcpServers": {
//	    "calculator": { "url": "http://localhost:8080/mcp" }
//	  }
//	}
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// addArgs 定义加法工具的输入参数。
// jsonschema 标签会自动生成 MCP 工具定义中的参数描述，
// 客户端据此向用户展示友好的输入界面。
type addArgs struct {
	A float64 `json:"a" jsonschema:"第一个数字"`
	B float64 `json:"b" jsonschema:"第二个数字"`
}

// subArgs 定义减法工具的输入参数。
type subArgs struct {
	A float64 `json:"a" jsonschema:"被减数"`
	B float64 `json:"b" jsonschema:"减数"`
}

func main() {
	// 初始化结构化日志，输出到 stderr。
	// HTTP 模式下 stdout 已不再被 MCP 协议占用，
	// 但日志打到 stderr 仍是惯例，方便与业务输出分离。
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo, // 默认 Info 级别，调试时可改为 LevelDebug
	}))
	slog.SetDefault(logger)

	slog.Info("initializing MCP calculator server",
		"name", "calculator-mcp-server",
		"version", "1.0.0",
		"transport", "streamable-http",
	)

	// 创建 MCP 服务实例，注入 logger 便于 go-sdk 内部日志统一输出
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "calculator-mcp-server",
		Version: "1.0.0",
	}, &mcp.ServerOptions{
		Logger: logger,
	})

	// 注册加法工具
	// AddTool 的泛型参数 addArgs 会自动推导参数类型，
	// go-sdk 内部会根据 jsonschema 标签生成 JSON Schema 并验证入参。
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

	// 创建 Streamable HTTP 处理器（2025-03-26 规范），绑定 /mcp 路由
	handler := mcp.NewStreamableHTTPHandler(func(r *http.Request) *mcp.Server {
		return server
	}, nil)

	http.HandleFunc("/mcp", handler.ServeHTTP)

	// 启动 HTTP 服务，阻塞直到进程被终止
	slog.Info("HTTP server starting, listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("HTTP server failed", "error", err)
		os.Exit(1)
	}
}
