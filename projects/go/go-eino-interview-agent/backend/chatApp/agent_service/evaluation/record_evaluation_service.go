package evaluation

import (
	"ai-eino-interview-agent/api/model/mianshi"
	"ai-eino-interview-agent/chatApp/agent/record_evaluation"
	"ai-eino-interview-agent/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

// GenerateRecordEvaluation 生成面试评估结果
func GenerateRecordEvaluation(ctx context.Context, userID uint, reportID uint64) (*mianshi.GetMianshiEvaluationResponse, error) {
	// 设置评测超时时间，防止长时间占用资源
	timeout, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	// 创建评测智能体
	agent, err := record_evaluation.NewRecordEvaluationAgent(userID)
	if err != nil {
		log.Printf("[GenerateRecordEvaluation] Error creating record evaluation agent: %v", err)
		return nil, err
	}

	// 创建评测智能体运行器
	runner := adk.NewRunner(timeout, adk.RunnerConfig{
		Agent: agent,
	})

	// 构建查询消息
	query := fmt.Sprintf(`请根据用户 ID %d 和面试报告 ID %d 生成面试评估结果，要求内容详实
请按照以下步骤进行：
1. 首先调用 get_mianshi_info 工具获取面试的完整问题和对话记录
2. 仔细分析候选人的回答内容
3. 根据回答质量进行综合评估
4. 生成详细的评估反馈

评估应包含：
- 评分（0-100分）
- 关键知识点的掌握情况
- 问题难度评估
- 回答的优势
- 回答的不足
- 改进建议
- 相关知识点总结
- 思考过程分析
- 参考答案或最佳实践	
`, userID, reportID)

	// 构建用户消息
	userMsg := &schema.Message{
		Role:    schema.User,
		Content: query,
	}

	// 运行评测智能体
	messages := []adk.Message{
		userMsg,
	}

	// 运行评测智能体
	iter := runner.Run(timeout, messages)

	// 处理评测结果
	var lastMessage string

	// 迭代获取评测智能体的输出，直到完成或超时
	for {
		// 检查是否超时
		select {
		case <-timeout.Done():
			log.Printf("[GenerateRecordEvaluation] Timeout reached while generating evaluation for user %d, report %d", userID, reportID)
			return nil, fmt.Errorf("timeout reached while generating evaluation")
		default:
		}

		// 获取评测智能体的输出
		event, ok := iter.Next()
		if !ok {
			break
		}

		// 处理评测智能体的输出
		if event.Err != nil {
			log.Printf("[GenerateRecordEvaluation] Error during evaluation generation for user %d, report %d: %v", userID, reportID, event.Err)
			return nil, fmt.Errorf("error during evaluation generation: %v", event.Err)
		}

		// 更新最后一条消息内容
		if event.Output != nil && event.Output.MessageOutput != nil {
			lastMessage = event.Output.MessageOutput.Message.Content
		}
	}

	// 构建评测结果响应
	records := buildEvaluationResponse(lastMessage)
	// 检查 records 是否为空
	if records == nil {
		log.Printf("[GenerateRecordEvaluation] No valid evaluation generated for user %d, report %d", userID, reportID)
		return nil, fmt.Errorf("no valid evaluation generated")
	}

	// 保存评测结果到数据库
	if err := saveEvaluationToDatabase(ctx, userID, reportID, records); err != nil {
		log.Printf("[GenerateRecordEvaluation] Error saving evaluation to database for user %d, report %d: %v", userID, reportID, err)
		return nil, fmt.Errorf("error saving evaluation to database: %v", err)
	}

	// 返回评测结果
	return records, nil
}

// saveEvaluationToDatabase 保存评测结果到数据库
func saveEvaluationToDatabase(_ context.Context, userID uint, reportID uint64, response *mianshi.GetMianshiEvaluationResponse) error {
	if response == nil {
		log.Printf("[saveEvaluationToDatabase] Invalid evaluation response for user %d, report %d", userID, reportID)
		return fmt.Errorf("invalid evaluation response")
	}
	// 构建评测结果数据结构
	var dimensionsList []*model.EvaluationDimension
	for _, dim := range response.Dimensions {
		dimensionsList = append(dimensionsList, &model.EvaluationDimension{
			DimensionName: dim.DimensionName,
			Evaluation:    dim.Evaluation,
			Score:         float64(dim.Score),
		})
	}
	// 计算总体评分（各个维度的平均分）
	var totalScore float64
	if len(response.Dimensions) > 0 {
		for _, dim := range response.Dimensions {
			totalScore += float64(dim.Score)
		}
		totalScore /= float64(len(response.Dimensions))
	}
	// 构建评测结果对象
	evaluation := &model.InterviewEvaluation{
		UserID:     userID,
		ReportID:   reportID,
		Comment:    response.Comment,
		Dimensions: dimensionsList,
		Score:      totalScore,
		Deleted:    0,
	}
	err := model.InterviewEvaluationDao.CreateEvaluation(evaluation)
	if err != nil {
		log.Printf("[saveEvaluationToDatabase] Error creating evaluation record for user %d, report %d: %v", userID, reportID, err)
		return fmt.Errorf("error creating evaluation record: %v", err)
	}
	// 实现保存逻辑
	return nil
}

// buildEvaluationResponse 构建评测结果响应
func buildEvaluationResponse(agentResponse string) *mianshi.GetMianshiEvaluationResponse {
	response := &mianshi.GetMianshiEvaluationResponse{
		Comment:    "",
		Dimensions: make([]*mianshi.MianshiEvaluationDimension, 0),
	}

	// 尝试直接解析 JSON
	if err := json.Unmarshal([]byte(agentResponse), response); err != nil {
		// 尝试从文本中提取 JSON
		jsonStr := ExtractJSONFromResponse(agentResponse)
		if jsonStr == "" {
			log.Printf("[buildEvaluationResponse] 无法提取 JSON，使用默认响应")
			return response // 返回默认值
		}

		// 尝试解析提取的 JSON
		if err := json.Unmarshal([]byte(jsonStr), response); err != nil {
			log.Printf("[buildEvaluationResponse] 解析提取的 JSON 失败: %v", err)
			return response // 返回默认值
		}
	}

	return response
}
