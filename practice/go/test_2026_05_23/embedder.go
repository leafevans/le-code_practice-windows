package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/openai"
)

func NewZhipuEmbedder(ctx context.Context) *openai.Embedder {
	embedder, err := openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
		APIKey:  os.Getenv("ZHIPU_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		BaseURL: "https://open.bigmodel.cn/api/paas/v4",
	})
	if err != nil {
		panic(err)
	}
	return embedder
}
