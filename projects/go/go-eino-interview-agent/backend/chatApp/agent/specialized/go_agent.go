package specialized

import (
	"ai-eino-interview-agent/chatApp/chat"
	"ai-eino-interview-agent/chatApp/tool"
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino/adk"
	componenttool "github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
)

func NewGoSpecializedAgent(userID uint, needResumeTool bool) (adk.Agent, error) {
	ctx := context.Background()

	model, err := chat.CreateOpenAIChatModel(ctx, userID)
	if err != nil {
		log.Fatalf("创建大模型失败：%v", err)
	}

	var toolsConfig adk.ToolsConfig
	if needResumeTool {
		toolsConfig = adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []componenttool.BaseTool{
					tool.GetResumeInfoTool(),
				},
			},
		}
	}

	baseAgent, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:          "GoSpecializedAgent",
		Description:   "Go语言专项面试官，负责对求职者进行Go语言相关知识的深度考察。",
		Instruction:   GoSpecializedAgentInstruction,
		Model:         model,
		MaxIterations: 15,
		ToolsConfig:   toolsConfig,
	})
	if err != nil {
		return nil, fmt.Errorf("创建智能体失败：%v", err)
	}

	return baseAgent, nil
}
