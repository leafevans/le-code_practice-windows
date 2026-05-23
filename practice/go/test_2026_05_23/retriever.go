package main

import (
	"context"

	"github.com/cloudwego/eino-ext/components/embedding/openai"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
)

func NewZhipuRetriever(ctx context.Context, embedder *openai.Embedder) *milvus.Retriever {

	retriever, err := milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:       MilvusCli,
		Collection:   "test",
		VectorField:  "vector",
		OutputFields: []string{"id", "content", "metadata"},
		TopK:         3,
		Embedding:    embedder,
	})
	if err != nil {
		panic(err)
	}
	return retriever
}
