package comprehensive

import (
	"ai-eino-interview-agent/chatApp/chat"
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino/adk"
)

func NewSchoolAgent() (adk.Agent, error) {
	ctx := context.Background()
	model, err := chat.CreateOpenAIChatModel(ctx)
	if err != nil {
		log.Fatalf("创建大模型失败：%v", err)
	}
	baseAgent, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:          "SchoolAgent",
		Description:   "校招综合面试官，负责对求职者进行综合面试，评估其综合能力和潜力。",
		Instruction:   SchoolComprehensiveAgentInstruction,
		Model:         model,
		MaxIterations: 5,
	})
	if err != nil {
		return nil, fmt.Errorf("创建智能体失败：%v", err)
	}
	return baseAgent, nil
}
