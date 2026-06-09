package mq

import (
	"context"
	"fmt"
	"sync"

	"github.com/bytedance/sonic"
)

// 消息类型定义
type MessageType string

// 定义消息类型常量
const (
	MessageTypeEvaluationReport MessageType = "evaluation_report"
	MessageTypeTopicEvaluation  MessageType = "topic_evaluation"
)

// Message 定义消息结构
type Message struct {
	Type    MessageType    `json:"type"`
	Payload map[string]any `json:"payload"`
}

// EvaluationReportPayload 定义评测结果消息的负载结构
type EvaluationReportPayload struct {
	UserID   uint   `json:"user_id"`
	ReportID uint64 `json:"report_id"`
}

// TopicEvaluationPayload 定义题目评测消息的负载结构
type TopicEvaluationPayload struct {
	UserID   uint   `json:"user_id"`
	ReportID uint64 `json:"report_id"`
}

// MessageQueue 定义消息队列接口
type MessageQueue interface {
	Publish(ctx context.Context, payload *Message) error
	Subscribe(ctx context.Context, handler MessageHandler) error
	Close() error
}

// MessageHandler 定义消息处理函数类型
type MessageHandler func(ctx context.Context, msg *Message) error

// InMemoryQueue 实现了 MessageQueue 接口的内存消息队列
type InMemoryQueue struct {
	mu        sync.Mutex
	messages  chan *Message
	handlers  []MessageHandler
	done      chan struct{}
	closeOnce sync.Once
	startOnce sync.Once
}

// NewInMemoryQueue 创建一个新的 InMemoryQueue 实例
func NewInMemoryQueue(bufferSize int) *InMemoryQueue {
	// 如果缓冲区大小小于等于0，使用默认值
	if bufferSize <= 0 {
		bufferSize = 100 // 默认缓冲区大小
	}

	// 创建并返回一个新的 InMemoryQueue 实例
	return &InMemoryQueue{
		messages: make(chan *Message, bufferSize),
		done:     make(chan struct{}),
	}
}

// Publish 发布消息到队列
func (q *InMemoryQueue) Publish(ctx context.Context, msg *Message) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-q.done:
		return fmt.Errorf("message queue is closed")
	case q.messages <- msg:
		return nil
	}
}

// Subscribe 注册消息处理器
func (q *InMemoryQueue) Subscribe(ctx context.Context, handler MessageHandler) error {
	q.mu.Lock()
	q.handlers = append(q.handlers, handler)
	q.mu.Unlock()
	return nil
}

// StartProcessing 启动消息消费循环，应在所有 handler 注册完成后调用
func (q *InMemoryQueue) StartProcessing(ctx context.Context) {
	q.startOnce.Do(func() {
		go q.processMessages(ctx)
	})
}

// processMessages 处理消息队列中的消息
func (q *InMemoryQueue) processMessages(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-q.done:
			return
		case msg := <-q.messages:
			if msg == nil {
				continue
			}
			q.mu.Lock()
			handlers := q.handlers
			q.mu.Unlock()
			for _, handler := range handlers {
				go func(h MessageHandler, m *Message) {
					if err := h(ctx, m); err != nil {
						fmt.Printf("Error handling message: %v\n", err)
					}
				}(handler, msg)
			}
		}
	}
}

// Close 关闭消息队列
func (q *InMemoryQueue) Close() error {
	q.closeOnce.Do(func() {
		close(q.done)
	})
	return nil
}

// 全局消息队列实例和互斥锁
var (
	globalMQ MessageQueue
	mqMutex  sync.RWMutex
)

// InitMessageQueue 初始化全局消息队列实例
func InitMessageQueue(mq MessageQueue) {
	mqMutex.Lock()
	defer mqMutex.Unlock()
	globalMQ = mq
}

// GetMessageQueue 获取全局消息队列实例
func GetMessageQueue() MessageQueue {
	mqMutex.RLock()
	defer mqMutex.RUnlock()
	if globalMQ == nil {
		return NewInMemoryQueue(100) // 默认使用内存消息队列
	}
	return globalMQ
}

// structToMap 将结构体转换为 map[string]any（基于 json tag）
func structToMap(v any) (map[string]any, error) {
	data, err := sonic.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("structToMap: failed to marshal: %w", err)
	}

	var result map[string]any
	if err := sonic.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("structToMap: failed to unmarshal: %w", err)
	}

	return result, nil
}

// PublishEvaluationReport 发布评测结果消息
func PublishEvaluationReport(ctx context.Context, userID uint, reportID uint64) error {
	// 获取全局消息队列实例
	mq := GetMessageQueue()

	// 将评测结果负载结构体转换为 map[string]any
	payloadMap, err := structToMap(&EvaluationReportPayload{
		UserID:   userID,
		ReportID: reportID,
	})
	if err != nil {
		return fmt.Errorf("failed to convert payload: %w", err)
	}

	// 发布消息到队列
	return mq.Publish(ctx, &Message{
		Type:    MessageTypeEvaluationReport,
		Payload: payloadMap,
	})
}

// PublishTopicEvaluation 发布题目评测消息
func PublishTopicEvaluation(ctx context.Context, userID uint, reportID uint64) error {
	mq := GetMessageQueue()

	payloadMap, err := structToMap(&TopicEvaluationPayload{
		UserID:   userID,
		ReportID: reportID,
	})
	if err != nil {
		return fmt.Errorf("failed to convert payload: %w", err)
	}

	return mq.Publish(ctx, &Message{
		Type:    MessageTypeTopicEvaluation,
		Payload: payloadMap,
	})
}
