package test

import (
	"context"
	"log"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func createChatTemplate() prompt.ChatTemplate {
	return prompt.FromMessages(
		schema.FString,
		schema.SystemMessage("你是{role}，用{style}的语气回答，帮助用户回答面试上的问题"),
		schema.UserMessage("问题: {question}"),
	)
}

func MessageTemplate() []*schema.Message {
	template := createChatTemplate()
	message, err := template.Format(context.Background(), map[string]any{
		"role":     "面试官",
		"style":    "正式",
		"question": "请介绍一下你自己",
	},
	)

	if err != nil {
		log.Fatalf("格式化消息模板失败：%v", err)
	}

	return message
}
