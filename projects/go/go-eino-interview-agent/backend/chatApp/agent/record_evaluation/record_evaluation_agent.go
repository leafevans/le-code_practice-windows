package record_evaluation

import (
	"ai-eino-interview-agent/chatApp/chat"
	"ai-eino-interview-agent/chatApp/tool"
	"fmt"

	"context"

	"github.com/cloudwego/eino/adk"
	componenttool "github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
)

func NewRecordEvaluationAgent(userID uint) (adk.Agent, error) {
	ctx := context.Background()

	// 创建聊天模型
	model, err := chat.CreateOpenAIChatModel(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAI chat model: %w", err)
	}

	// 创建智能体
	baseAgent, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "RecordEvaluationAgent",
		Description: "一个专业评估面试记录并生成专业报告的智能体",
		Model:       model,
		Instruction: RecordEvaluationInstruction,
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []componenttool.BaseTool{
					tool.GetMianshiInfoTool(),
				},
			},
		},
		MaxIterations: 15,
	})

	// 处理智能体创建错误
	if err != nil {
		return nil, fmt.Errorf("failed to create evaluation agent: %w", err)
	}
	return baseAgent, nil
}
