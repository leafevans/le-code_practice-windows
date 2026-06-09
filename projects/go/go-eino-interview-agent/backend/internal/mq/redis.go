package mq

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
)

// RedisQueue 实现了 MessageQueue 接口的 Redis 消息队列
type RedisQueue struct {
	client   *redis.Client
	mu       sync.RWMutex
	handlers []MessageHandler
	done     chan struct{}
}

// NewRedisQueue 创建一个新的 RedisQueue 实例
func NewRedisQueue(client *redis.Client) *RedisQueue {
	// 确保 Redis 客户端不为 nil
	if client == nil {
		log.Fatal("Redis client cannot be nil")
	}

	// 创建并返回一个新的 RedisQueue 实例
	return &RedisQueue{
		client: client,
		done:   make(chan struct{}),
	}
}

// Publish 发布消息到 Redis
func (q *RedisQueue) Publish(ctx context.Context, msg *Message) error {
	// 将消息转换为 JSON 格式
	data, err := sonic.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// 构建 Redis 频道名称
	channel := fmt.Sprintf("message_queue:%s", msg.Type)
	log.Printf("Publishing message to channel %s: %s", channel, string(data))

	// 发布消息到 Redis 频道
	res := q.client.Publish(ctx, channel, data)
	if res.Err() != nil {
		return fmt.Errorf("failed to publish message: %w", res.Err())
	}

	// 获取订阅该频道的订阅者数量
	numSubscribers := res.Val()
	log.Printf("Message published to channel %s, number of subscribers: %d", channel, numSubscribers)

	return nil
}

// Subscribe 订阅消息并处理
func (q *RedisQueue) Subscribe(ctx context.Context, handler MessageHandler, msgTypes ...string) {
	// 构建订阅的频道列表
	q.mu.Lock()
	q.handlers = append(q.handlers, handler)
	q.mu.Unlock()

	// 如果没有指定消息类型，默认订阅所有类型
	channels := []string{
		fmt.Sprintf("message_queue:%s", MessageTypeEvaluationReport),
		fmt.Sprintf("message_queue:%s", MessageTypeTopicEvaluation),
	}
	log.Printf("Subscribing to channels: %v", channels)

	// 订阅 Redis 频道
	pubsub := q.client.Subscribe(ctx, channels...)
	defer pubsub.Close()

	// 处理订阅的消息
	ch := pubsub.Channel()
	msgCount := 0
	for {
		select {
		// 监听上下文取消和队列关闭事件
		case <-ctx.Done():
			log.Println("Context cancelled, stopping subscription")
			return
		// 监听队列关闭事件
		case <-q.done:
			log.Println("Message queue is closed, stopping subscription")
			return
		// 监听 Redis 频道的消息
		case msg := <-ch:
			// 处理接收到的消息
			if msg == nil {
				log.Println("Received nil message, skipping")
				continue
			}
			msgCount++
			log.Printf("Received message from channel %s: %s", msg.Channel, msg.Payload)
			var message Message
			if err := sonic.Unmarshal([]byte(msg.Payload), &message); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}
			q.mu.RLock()
			handlers := q.handlers
			q.mu.RUnlock()

			// 并发处理消息
			for _, handler := range handlers {
				go func(h MessageHandler, m *Message) {
					if err := h(ctx, &message); err != nil {
						log.Printf("Error handling message: %v", err)
					}
				}(handler, &message)
			}
		}
	}
}

// Close 关闭消息队列
func (q *RedisQueue) Close() error {
	close(q.done)
	return nil
}

// GetRedisClient 获取 Redis 客户端实例
func GetRedisClient(q *RedisQueue) *redis.Client {
	return q.client
}
