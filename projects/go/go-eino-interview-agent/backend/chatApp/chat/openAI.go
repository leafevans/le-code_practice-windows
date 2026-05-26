package chat

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

func CreateOpenAIChatModel(ctx context.Context) (model.ToolCallingChatModel, error) {
	APIKey := "sk-9f624ecf95bd44ff80b50275250ab9c0"
	modelName := "deepseek-v4-flash"
	baseURL := "https://api.deepseek.com"

	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  APIKey,
		Model:   modelName,
		BaseURL: baseURL,
	})
	if err != nil {
		log.Fatalf("创建大模型失败： %v", err)
	}
	return chatModel, nil
}
