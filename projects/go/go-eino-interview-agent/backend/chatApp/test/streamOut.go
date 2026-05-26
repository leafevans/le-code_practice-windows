package test

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

func ReportStream(reader *schema.StreamReader[*schema.Message]) {
	defer reader.Close()
	for {
		chunk, err := reader.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("failed to receive chunk: %v", err)
		}
		fmt.Print(chunk.Content)
	}
}

func Steam(ctx context.Context, model model.ToolCallingChatModel, message []*schema.Message) *schema.StreamReader[*schema.Message] {
	result, err := model.Stream(ctx, message)
	if err != nil {
		log.Fatalf("failed to stream: %v", err)
	}
	return result
}
