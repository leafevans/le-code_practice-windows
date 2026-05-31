package record_evaluation

import (
	"ai-eino-interview-agent/chatApp/chat"
	"ai-eino-interview-agent/chatApp/tool"
	"context"
	"fmt"

	"github.com/cloudwego/eino/adk"
	componenttool "github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
)

func NewAnswerRecordAgent(userID uint) (adk.Agent, error) {
	ctx := context.Background()

	// 创建聊天模型
	model, err := chat.CreateOpenAIChatModel(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat model: %w", err)
	}

	// 创建智能体
	agent, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "AnswerRecordAgent",
		Description: "一个专业的面试记录评估智能体，负责对面试记录进行评估，并为候选人提供专业的反馈。",
		Instruction: AnswerRecordAgentInstruction,
		Model:       model,
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []componenttool.BaseTool{
					tool.GetMianshiInfoTool(),
				},
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create AnswerRecordAgent: %w", err)
	}

	return agent, nil
}
