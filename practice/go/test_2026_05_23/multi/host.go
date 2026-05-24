package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/flow/agent/multiagent/host"
)

func newHost(ctx context.Context) (*host.Host, error) {
	deepseekAPIKey := os.Getenv("DEEPSEEK_API_KEY")
	deepseekModelName := os.Getenv("MODEL")
	chatModel, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		APIKey: deepseekAPIKey,
		Model:  deepseekModelName,
	})
	if err != nil {
		fmt.Printf("failed to create chat model: %v", err)
		return nil, err
	}

	return &host.Host{
		ToolCallingModel: chatModel,
		SystemPrompt:     "你可以同时计算加法和减法。当用户提问时，你需要回答问题。",
	}, nil
}
