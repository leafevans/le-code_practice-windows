package evaluation

import (
	"ai-eino-interview-agent/chatApp/agent/record_evaluation"
	"ai-eino-interview-agent/internal/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

func GenerateAnswerRecordEvaluation(ctx context.Context, userID uint, reportID uint64) (*model.AnswerReport, error) {
	// 设置评测超时时间，防止长时间占用资源
	timeout, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	// 创建评测智能体
	agent, err := record_evaluation.NewAnswerRecordAgent(userID)
	if err != nil {
		log.Printf("Failed to create evaluation agent: %v", err)
		return nil, err
	}

	// 创建评测智能体运行器
	runner := adk.NewRunner(timeout, adk.RunnerConfig{
		Agent: agent,
	})

	// 构建查询消息
	query := fmt.Sprintf(`请对用户ID为 %d、报告ID为 %d 的答题记录进行详细评估。

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
- 参考答案或最佳实践`, userID, reportID)

	// 构建用户消息
	userMsg := &schema.Message{
		Role:    schema.User,
		Content: query,
	}
	messages := []adk.Message{userMsg}

	// 运行评测智能体
	iter := runner.Run(ctx, messages)

	// 迭代获取评测结果，直到完成或超时
	var lastMsg string
	for {
		// 检查上下文是否超时
		select {
		case <-timeout.Done():
			log.Printf("Evaluation timed out for userID %d, reportID %d", userID, reportID)
			return nil, fmt.Errorf("evaluation timed out")
		default:
		}

		// 获取下一条评测结果
		event, ok := iter.Next()
		if !ok {
			break
		}
		// 处理事件错误
		if event.Err != nil {
			log.Printf("Error during evaluation for userID %d, reportID %d: %v", userID, reportID, event.Err)
			return nil, fmt.Errorf("error during evaluation: %v", event.Err)
		}

		if event.Output != nil && event.Output.MessageOutput != nil {
			lastMsg = event.Output.MessageOutput.Message.Content
			log.Printf("Received evaluation output for userID %d, reportID %d: %s", userID, reportID, lastMsg)
		}
	}

	// 构建评测报告响应
	report := buildAnswerReportResponse(lastMsg, userID, reportID)

	// 保存评测报告到数据库
	if err := saveAnswerReportToDatabase(report); err != nil {
		log.Printf("Failed to save evaluation report for userID %d, reportID %d: %v", userID, reportID, err)
	}

	return report, nil
}

// saveAnswerReportToDatabase 将评测报告保存到数据库
func saveAnswerReportToDatabase(report *model.AnswerReport) error {
	err := model.AnswerReportDao.CreateAnswerReport(report)
	if err != nil {
		log.Printf("Error saving answer report to database for userID %d, reportID %d: %v", report.UserID, report.ReportID, err)
		return fmt.Errorf("failed to save answer report to database: %v", err)
	}
	log.Printf("Successfully saved answer report to database for userID %d, reportID %d", report.UserID, report.ReportID)
	return nil
}

// 构建评测报告响应
func buildAnswerReportResponse(agentResponse string, userID uint, reportID uint64) *model.AnswerReport {
	// 解析智能体响应，构建评测报告结构
	report := &model.AnswerReport{
		UserID:   userID,
		ReportID: reportID,
		Records:  make([]*model.AnswerRecordItem, 0),
		Deleted:  0,
	}

	// 定义临时结构体用于解析智能体响应
	type TempResponse struct {
		Records []map[string]any `json:"records"`
	}

	// 解析智能体响应
	var tempResp TempResponse

	if err := sonic.Unmarshal([]byte(agentResponse), &tempResp); err != nil {
		// 如果直接解析失败，尝试从响应中提取 JSON 字符串并重新解析
		jsonStr := ExtractJSONFromResponse(agentResponse)

		// 	如果提取失败或解析失败，记录日志并返回默认评测报告
		if jsonStr == "" {
			log.Printf("Failed to extract JSON from agent response for userID %d, reportID %d: %v", userID, reportID, err)
			return buildDefaultAnswerReport(userID, reportID)
		}

		//	尝试解析提取的 JSON 字符串
		if err := sonic.Unmarshal([]byte(jsonStr), &tempResp); err != nil {
			log.Printf("Failed to parse extracted JSON from agent response for userID %d, reportID %d: %v", userID, reportID, err)
			return buildDefaultAnswerReport(userID, reportID)
		}
	}

	for _, record := range tempResp.Records {
		record := convertMapToAnswerRecordItem(record)
		if record != nil {
			report.Records = append(report.Records, record)
		}
	}

	return report
}

// buildDefaultAnswerReport 构建默认的评测报告
func buildDefaultAnswerReport(userID uint, reportID uint64) *model.AnswerReport {
	return &model.AnswerReport{
		UserID:   userID,
		ReportID: reportID,
		Records:  make([]*model.AnswerRecordItem, 0),
		Deleted:  0,
	}
}

// convertMapToAnswerRecordItem 将 map 转换为 AnswerRecordItem 结构体
func convertMapToAnswerRecordItem(recordMap map[string]any) *model.AnswerRecordItem {
	// 创建 AnswerRecordItem 结构体实例
	record := &model.AnswerRecordItem{}

	// 解析 order
	if order, ok := recordMap["order"].(float64); ok {
		record.Order = int32(order)
	}

	// 解析 content
	if content, ok := recordMap["content"].(string); ok {
		record.Content = content
	}

	// 解析 comment
	if commentMap, ok := recordMap["comment"].(map[string]any); ok {
		record.Comment = convertMapToAnswerRecordComment(commentMap)
	}

	// 解析 message
	if messageList, ok := recordMap["message"].([]any); ok {
		// 初始化消息列表
		record.Message = make([]*model.AnswerRecordMsg, 0)
		for _, msg := range messageList {
			// 解析每条对话记录
			if msgMap, ok := msg.(map[string]any); ok {
				record.Message = append(record.Message, convertMapToAnswerRecordMsg(msgMap))
			}
		}
	}

	return record
}

// convertMapToAnswerRecordMsg 将 map 转换为 AnswerRecordMsg 结构体
func convertMapToAnswerRecordComment(commentMap map[string]any) *model.AnswerRecordComment {
	comment := &model.AnswerRecordComment{}

	// 解析 score
	if score, ok := commentMap["score"].(float64); ok {
		comment.Score = int32(score)
	}

	// 解析 key_points
	if keyPoints, ok := commentMap["key_points"].(string); ok {
		comment.KeyPoints = keyPoints
	}

	// 解析 difficulty
	if difficulty, ok := commentMap["difficulty"].(string); ok {
		comment.Difficulty = difficulty
	}

	// 解析 strengths
	if strengths, ok := commentMap["strengths"].(string); ok {
		comment.Strengths = strengths
	}

	// 解析 weaknesses
	if weaknesses, ok := commentMap["weaknesses"].(string); ok {
		comment.Weaknesses = weaknesses
	}

	// 解析 suggestion
	if suggestion, ok := commentMap["suggestion"].(string); ok {
		comment.Suggestion = suggestion
	}

	// 解析 know_points
	if knowPoints, ok := commentMap["know_points"].(string); ok {
		comment.KnowPoints = knowPoints
	}

	// 解析 thinking
	if thinking, ok := commentMap["thinking"].(string); ok {
		comment.Thinking = thinking
	}

	// 解析 reference
	if reference, ok := commentMap["reference"].(string); ok {
		comment.Reference = reference
	}

	return comment
}

// convertMapToAnswerRecordMsg 将 map 转换为 AnswerRecordMsg 结构体
func convertMapToAnswerRecordMsg(msgMap map[string]any) *model.AnswerRecordMsg {
	msg := &model.AnswerRecordMsg{}

	// 解析 order
	if order, ok := msgMap["order"].(float64); ok {
		msg.Order = int32(order)
	}

	// 解析 question
	if question, ok := msgMap["question"].(string); ok {
		msg.Question = question
	}

	// 解析 answer
	if answer, ok := msgMap["answer"].(string); ok {
		msg.Answer = answer
	}

	return msg
}
