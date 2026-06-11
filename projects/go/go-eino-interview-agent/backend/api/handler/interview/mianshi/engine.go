package mianshi

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"ai-eino-interview-agent/chatApp/agent_service/interview"
	"ai-eino-interview-agent/internal/model"
	"ai-eino-interview-agent/internal/mq"
	interviewservice "ai-eino-interview-agent/internal/service/interviews"

	"github.com/bytedance/sonic"
)

// InterviewDialogData 面试对话数据
type InterviewDialogData struct {
	Question string // 问题
	Answer   string // 回答
}

// InterviewEngine 面试引擎
type InterviewEngine struct {
	sessionManager *SessionManager                   // 依赖：会话管理器
	interviewSvc   interviewservice.InterviewManager // 依赖：面试服务
	writer         io.Writer                         // 依赖：输出流
}

// NewInterviewEngine 创建新的面试引擎实例
func NewInterviewEngine(sessionManager *SessionManager, interviewSvc interviewservice.InterviewManager, writer io.Writer) *InterviewEngine {
	return &InterviewEngine{
		sessionManager: sessionManager,
		interviewSvc:   interviewSvc,
		writer:         writer,
	}
}

// RunInterviewLoop 运行面试对话循环
func (e *InterviewEngine) RunInterviewLoop(ctx context.Context, session *InterviewSession) {
	// 初始化面试对话数据
	const answerTimeout = 30 * time.Minute     // 回答超时时间
	const heartbeatInterval = 15 * time.Second // 心跳间隔
	const maxQuestions = 20                    // 最多生成20道问题
	const historyContextSize = 5               // 保留前5道题作为历史上下文

	// 创建智能体服务
	agentService := interview.NewInterviewAgentService(session.UserID)

	// 确定智能体类型
	agentType := e.selectAgentType(session)

	// 用于存储所有问题和回答
	var allDialogues []*InterviewDialogData

	// 用于存储最近的历史记录（前 5 道题）
	type HistoryItem struct {
		Question string
		Answer   string
	}

	//
	var recentHistory []HistoryItem

	for questionIndex := 1; questionIndex <= maxQuestions; questionIndex++ {
		// 检查上下文是否被取消
		select {
		case <-ctx.Done():
			log.Printf("面试对话被取消: %v", ctx.Err())
			return
		default:
		}

		// 构建 AI 生成问题的提示词
		var prompt string
		if questionIndex == 1 {
			prompt = fmt.Sprintf(`请根据简历ID和难度等级生成一道面试问题。

简历ID: %d
难度等级: %s

要求：
1. 生成一道技术面试问题
2. 返回JSON格式

返回格式：
{
  "question_text": "问题内容"
}`, session.ResumeID, session.Difficulty)
		} else {
			// 后续问题：包含最近5道题的历史上下文
			historyText := ""
			for i, h := range recentHistory {
				historyText += fmt.Sprintf("问题%d：%s\n回答%d：%s\n\n", i+1, h.Question, i+1, h.Answer)
			}

			prompt = fmt.Sprintf(`根据简历ID、难度等级和最近的问答历史，生成下一道面试问题。

简历ID: %d
难度等级: %s

最近的问答历史（前%d道题）：
%s

要求：
1. 根据用户的回答情况，生成更有针对性的下一道问题
2. 避免重复之前问过的问题
3. 逐步深化问题难度
4. 返回JSON格式

返回格式：
{
  "question_text": "问题内容"
}`, session.ResumeID, session.Difficulty, len(recentHistory), historyText)
		}

		// 调用智能体服务生成问题
		var questionResult map[string]any

		// 运行智能体服务，传入回调函数处理 AI 返回的消息
		err := agentService.RunInterviewWithCallback(ctx, agentType, session.HasResume, prompt, func(message string) error {
			// 回调函数：解析 AI 返回 JSON 字符串
			var result map[string]any
			if err := sonic.Unmarshal([]byte(message), &result); err != nil {
				questionResult = result
			}
			return nil
		})

		// 错误处理
		if err != nil {
			log.Printf("运行AI面试服务时出错: %v", err)
			SendErrorEvent(e.writer, fmt.Sprint("面试服务出错: ", err))
			return
		}

		// 提取问题的文本
		questionText, ok := questionResult["question_text"].(string)
		if !ok || questionText == "" {
			log.Printf("AI返回的结果缺少有效的问题文本: %v", questionResult)
			SendErrorEvent(e.writer, "AI返回的结果缺少有效的问题文本")
			break
		}

		// 向客户端发送问题事件
		err = SendSSEEvent(e.writer, map[string]any{
			"type":  "question",
			"index": questionIndex,
			"total": maxQuestions,
			"data": map[string]any{
				"question_text": questionText,
			},
		})
		if err != nil {
			log.Printf("发送问题事件时出错: %v", err)
			return
		}

		// 向客户端发送就绪事件
		SendReadyEventWithSession(e.writer, questionIndex, session.SessionID)
		e.sessionManager.ClearAnswer(session.SessionID)

		// 等待用户回答
		log.Printf("等待用户回答问题%d...", questionIndex)
		answer, received := WaitForAnswerWithHeartbeat(ctx, e.sessionManager, session.SessionID, heartbeatInterval, answerTimeout, e.writer)
		if !received {
			log.Printf("未收到用户回答，结束面试对话: sessionID=%s", session.SessionID)
			SendErrorEvent(e.writer, "未收到用户回答，面试结束")
			break
		}

		// 保存问答数据，更新历史
		dialogue := &InterviewDialogData{
			Question: questionText,
			Answer:   answer,
		}
		allDialogues = append(allDialogues, dialogue)
		session.QuestionCount = int32(questionIndex)

		// 更新最近的历史记录
		recentHistory = append(recentHistory, HistoryItem{
			Question: questionText,
			Answer:   answer,
		})
		if len(recentHistory) > historyContextSize {
			recentHistory = recentHistory[len(recentHistory)-historyContextSize:]
		}

		// 向客户端发送收到答案的事件
		err = SendSSEEvent(e.writer, map[string]any{
			"type":     "answer_received",
			"index":    questionIndex,
			"total":    maxQuestions,
			"progress": float64(questionIndex) / float64(maxQuestions) * 100,
		})

		if err != nil {
			log.Printf("发送收到答案事件时出错: %v", err)
		}

		log.Printf("用户回答了问题%d: %s", questionIndex, answer)
	}

	// 保存所有对话数据到数据库
	err := e.saveAllDialogues(ctx, session, allDialogues)
	if err != nil {
		log.Printf("保存面试对话数据时出错: %v", err)
		SendErrorEvent(e.writer, "保存面试对话数据时出错")
		SendCompleteEvent(e.writer)
		return
	}

	// 发送完成事件
	SendCompleteEvent(e.writer)

	// 发布评估报告生成消息
	if err = mq.PublishEvaluationReport(ctx, session.UserID, session.RecordID); err != nil {
		log.Printf("[Interview Loop] Failed to publish evaluation report message: %v, sessionID: %s", err, session.SessionID)
	}

	// 发布主题评估消息
	if err = mq.PublishTopicEvaluation(ctx, session.UserID, session.RecordID); err != nil {
		log.Printf("[Interview Loop] Failed to publish topic evaluation message: %v, sessionID: %s", err, session.SessionID)
	}
}

// selectAgentType 根据会话信息选择智能体类型
func (e *InterviewEngine) selectAgentType(session *InterviewSession) interview.InterviewAgentType {
	// 综合面试
	if session.Type == "综合面试" {
		switch session.Domain {
		case "校招简历面试":
			return interview.ComprehensiveSchool
		case "社招简历面试":
			return interview.ComprehensiveSocial
		default:
			// 社招简历面试为默认选项
			return interview.ComprehensiveSocial
		}
	}

	// 专项面试
	switch session.Domain {
	case "Java":
		return interview.SpecializedJava
	case "MQ":
		return interview.SpecializedMQ
	case "MySQL":
		return interview.SpecializedMySQL
	case "Redis":
		return interview.SpecializedRedis
	case "Go":
		fallthrough
	default:
		return interview.SpecializedGo
	}
}

// saveAllDialogues 保存所有对话数据到数据库
func (e *InterviewEngine) saveAllDialogues(ctx context.Context, session *InterviewSession, questions []*InterviewDialogData) error {
	log.Printf("正在保存面试对话数据: sessionID=%s, questionCount=%d", session.SessionID, len(questions))
	for i, q := range questions {
		dialogue := &model.InterviewDialogue{
			UserID:    session.UserID,
			ReportID:  session.RecordID,
			Question:  q.Question,
			Answer:    q.Answer,
			CreatedAt: time.Now(),
		}

		if err := model.InterviewDialogueDao.Create(dialogue); err != nil {
			log.Printf("保存第%d道对话数据时出错: %v", i+1, err)
			return fmt.Errorf("保存对话数据时出错: %v", err)
		}

		if (i+1)%10 == 0 {
			log.Printf("已保存%d道对话数据", i+1)
		}
	}
	log.Printf("完成保存面试对话数据: sessionID=%s, totalSaved=%d", session.SessionID, len(questions))
	return nil
}
