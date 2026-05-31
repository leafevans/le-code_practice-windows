package mianshi

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bytedance/sonic"
)

// SendSSEEvent 发送 SSE 事件
func SendSSEEvent(writer io.Writer, event map[string]any) error {
	// 将事件数据序列化为 JSON 格式
	eventJSON, err := sonic.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v", err)
		return err
	}

	// 获取事件类型
	eventType := "message" // 默认事件类型
	if t, ok := event["type"].(string); ok {
		eventType = t
	}

	// 构建 SSE 事件字符串
	message := fmt.Sprintf("event: %s\ndata: %s\n\n", eventType, string(eventJSON))

	// 将事件写入响应流
	n, err := fmt.Fprint(writer, message)
	if err != nil {
		log.Printf("Failed to write event: %v, bytes written: %d", err, n)
		return err
	}

	// 刷新缓冲区，确保事件立即发送到客户端
	if flusher, ok := writer.(http.Flusher); ok {
		flusher.Flush() // 刷新缓冲区
	}
	return nil
}

// SendErrorEvent 发送错误事件
func SendErrorEvent(writer io.Writer, message string) {
	event := map[string]any{
		"type":    "error",
		"message": message,
	}
	err := SendSSEEvent(writer, event)
	if err != nil {
		return
	}
}

// SendCompleteEvent 发送完成事件
func SendCompleteEvent(writer io.Writer) {
	event := map[string]any{
		"type":    "complete",
		"message": "面试结束",
	}
	err := SendSSEEvent(writer, event)
	if err != nil {
		return
	}
}

// SendReadyEventWithSession 发送就绪事件
func SendReadyEventWithSession(writer io.Writer, questionIndex int, sessionID string) {
	event := map[string]any{
		"type":          "ready_for_answer", // 事件类型
		"message":       "请回答上述问题",          // 事件消息
		"questionIndex": questionIndex,      // 当前问题索引
		"sessionID":     sessionID,          // 会话 ID
	}
	err := SendSSEEvent(writer, event)
	if err != nil {
		return
	}
}

// SendHeartbeatEvent 发送心跳事件
func SendHeartbeatEvent(writer io.Writer) error {
	event := map[string]any{
		"type":    "heartbeat",
		"message": "连接保活",
	}
	return SendSSEEvent(writer, event)
}
