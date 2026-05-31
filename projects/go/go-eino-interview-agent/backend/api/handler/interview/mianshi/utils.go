package mianshi

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// WaitForAnswerWithHeartbeat 等待用户答案，并定期发送心跳保活
func WaitForAnswerWithHeartbeat(ctx context.Context, manager *SessionManager, sessionID string, heartbeatInterval time.Duration, timeout time.Duration, writer io.Writer) (string, bool) {
	log.Printf("Waiting for answer with heartbeat: sessionID=%s, heartbeatInterval=%s, timeout=%s", sessionID, heartbeatInterval, timeout)

	// 创建心跳计时器，每隔 heartbeatInterval 发送一次心跳
	heartbeatTicker := time.NewTicker(heartbeatInterval)
	defer heartbeatTicker.Stop()

	// 创建一个超时计时器，超过 timeout 后停止等待
	timeoutTimer := time.NewTimer(timeout)
	defer timeoutTimer.Stop()

	session := manager.GetSession(sessionID)
	if session == nil {
		log.Printf("Session not found: sessionID=%s", sessionID)
		return "", false
	}

	heartbeatCount := 0

	// 无限循环等待用户输入或心跳事件
	for {
		select {
		// 事件 1: 上下文本被取消，客户端断开连接，主动结束等待
		case <-ctx.Done():
			log.Printf("Context cancelled, stopping wait: sessionID=%s", sessionID)
			return "", false
		// 事件 2: 心跳计时器触发，发送心跳保活
		case <-heartbeatTicker.C:
			heartbeatCount++
			log.Printf("Sending heartbeat #%d: sessionID=%s", heartbeatCount, sessionID)
			err := SendHeartbeatEvent(writer)
			if err != nil {
				log.Printf("Failed to send heartbeat: sessionID=%s, error=%v", sessionID, err)
				return "", false
			}
		// 事件 3：用户输入答案，结束等待
		case answer := <-session.AnswerChan:
			log.Printf("Received answer: sessionID=%s, answer=%s", sessionID, answer)
			return answer, true
		// 事件 4: 超时计时器触发，结束等待
		case <-timeoutTimer.C:
			log.Printf("Timeout reached, stopping wait: sessionID=%s", sessionID)
			return "", false
		}
	}
}

// SetupSSEResonse 设置 SSE 响应头
// 参数 c 是一个接口，包含设置 HTTP 状态码和响应头的方法
func SetupSSEResonse(c interface {
	SetStatusCode(int)     // 设置 HTTP 状态码
	Header(string, string) // 设置响应头（key, value）
}) {
	c.SetStatusCode(consts.StatusOK)                                                       // 设置 HTTP 状态码为 200 OK
	c.Header("Content-Type", "text/event-stream; charset=utf-8")                           // 设置响应头，指定内容类型为 SSE
	c.Header("Cache-Control", "no-cache")                                                  // 设置响应头，禁止缓存
	c.Header("Connection", "keep-alive")                                                   // 设置响应头，保持连接
	c.Header("Access-Control-Allow-Origin", "*")                                           // 设置响应头，允许跨域请求
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                         // 设置响应头，允许的 HTTP 方法
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Cache-Control") // 设置响应头，允许的 HTTP 头
	c.Header("Transfer-Encoding", "chunked")                                               // 设置响应头，启用分块传输编码
	c.Header("X-Accel-Buffering", "no")                                                    // 设置响应头，禁用 Nginx 的响应缓冲
	c.Header("X-Content-Type-Options", "nosniff")                                          // 设置响应头，防止 MIME 类型嗅探
}
