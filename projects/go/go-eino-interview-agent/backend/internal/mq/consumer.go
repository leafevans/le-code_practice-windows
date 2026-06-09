package mq

import (
	"ai-eino-interview-agent/chatApp/agent_service/evaluation"
	"context"
	"fmt"
	"log"
)

// ConsumerHandler 定义消费者处理器
type ConsumerHandler struct{}

// NewConsumerHandler 创建一个新的 ConsumerHandler 实例
func NewConsumerHandler() *ConsumerHandler {
	return &ConsumerHandler{}
}

// HandleMessage 处理接收到的消息
func (h *ConsumerHandler) HandleMessage(ctx context.Context, message *Message) error {
	switch message.Type {
	case MessageTypeEvaluationReport:
		return h.handleEvaluationReport(ctx, message)
	case MessageTypeTopicEvaluation:
		return h.handleTopicEvaluation(ctx, message)
	default:
		return fmt.Errorf("unknown message type: %s", message.Type)
	}
}

// handleEvaluationReport 处理评测结果消息
func (h *ConsumerHandler) handleEvaluationReport(ctx context.Context, message *Message) error {
	log.Printf("[Consumer] Processing evaluation report message")

	// 从消息负载中提取 user_id 和 report_id
	userID, ok := message.Payload["user_id"].(uint)
	if !ok {
		return fmt.Errorf("invalid user_id in payload")
	}

	// 评测报告ID通常是一个较大的数字，使用 uint64 类型
	reportID, ok := message.Payload["report_id"].(uint64)
	if !ok {
		return fmt.Errorf("invalid report_id in payload")
	}

	log.Printf("[Consumer] Generating evaluation report: userID=%d, reportID=%d", userID, reportID)

	// 调用评测服务生成评测报告
	_, err := evaluation.GenerateRecordEvaluation(ctx, userID, reportID)
	if err != nil {
		log.Printf("[Consumer] Failed to generate evaluation report: %v", err)
		return err
	}

	log.Printf("[Consumer] Evaluation report generated successfully: userID=%d, reportID=%d", userID, reportID)
	return nil
}

// handleTopicEvaluation 处理题目评测消息
func (h *ConsumerHandler) handleTopicEvaluation(ctx context.Context, message *Message) error {
	log.Printf("[Consumer] Processing topic evaluation message")

	// 从消息负载中提取 user_id 和 report_id
	userID, ok := message.Payload["user_id"].(uint)
	if !ok {
		return fmt.Errorf("invalid user_id in payload")
	}

	// 题目评测报告ID通常也是一个较大的数字，使用 uint64 类型
	reportID, ok := message.Payload["report_id"].(uint64)
	if !ok {
		return fmt.Errorf("invalid report_id in payload")
	}

	log.Printf("[Consumer] Generating topic evaluation: userID=%d, reportID=%d", userID, reportID)

	// 调用评测服务生成题目评测
	_, err := evaluation.GenerateAnswerRecordEvaluation(ctx, userID, reportID)
	if err != nil {
		log.Printf("[Consumer] Failed to generate topic evaluation: %v", err)
		return err
	}

	log.Printf("[Consumer] Topic evaluation generated successfully: userID=%d, reportID=%d", userID, reportID)
	return nil
}

// StartConsumer 启动消息消费者
func StartConsumer(ctx context.Context) error {
	// 获取全局消息队列实例
	mq := GetMessageQueue()
	handler := NewConsumerHandler()

	log.Printf("[Consumer] Starting message consumer")
	log.Printf("[Consumer] MQ type: %T", mq)

	if err := mq.Subscribe(ctx, handler.HandleMessage); err != nil {
		return fmt.Errorf("failed to subscribe: %w", err)
	}

	// 注册完成后启动消费循环
	if imq, ok := mq.(*InMemoryQueue); ok {
		imq.StartProcessing(ctx)
	}

	log.Printf("[Consumer] Consumer started, waiting for messages...")
	<-ctx.Done()
	log.Printf("[Consumer] Consumer stopped: %v", ctx.Err())
	return ctx.Err()
}
