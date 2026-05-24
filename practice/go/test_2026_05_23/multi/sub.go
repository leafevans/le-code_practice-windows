package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/flow/agent/multiagent/host"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

func newSubSpecialist(ctx context.Context) (*host.Specialist, error) {
	deepseekAPIKey := os.Getenv("DEEPSEEK_API_KEY")
	deepseekModelName := os.Getenv("MODEL")
	chatModel, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		APIKey: deepseekAPIKey,
		Model:  deepseekModelName,
	})
	if err != nil {
		return nil, err
	}
	subTool := GetSubTool()
	raAgent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: chatModel,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: []tool.BaseTool{subTool},
		},
	})
	return &host.Specialist{
		AgentMeta: host.AgentMeta{
			Name:        "sub_specialist",
			IntendedUse: "subtract two numbers and return the result",
		},
		Invokable: func(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.Message, error) {
			return raAgent.Generate(ctx, input, opts...)
		},
	}, nil

}
