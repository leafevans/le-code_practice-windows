package main

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/components/document"
)

func NewTransformer(ctx context.Context) document.Transformer {
	splitter, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":      "h1",
			"##":     "h2",
			"###":    "h3",
			"####":   "h4",
			"#####":  "h5",
			"######": "h6",
		},
		TrimHeaders: true,
	})
	if err != nil {
		panic(err)
	}
	return splitter
}
