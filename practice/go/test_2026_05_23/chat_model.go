package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
)

func NewDeepseek(ctx context.Context) *deepseek.ChatModel {
	model, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		APIKey: os.Getenv("DEEPSEEK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	if err != nil {
		panic(err)
	}
	return model
}
