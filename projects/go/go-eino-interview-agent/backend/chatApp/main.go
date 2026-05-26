package main

import (
	"ai-eino-interview-agent/agent/comprehensive"
	"ai-eino-interview-agent/agent/resume"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

func main() {
	schoolAgent, err := comprehensive.NewSchoolAgent()
	if err != nil {
		log.Fatalf("学校智能体创建失败：%v", err)
	}
	resumeAgent, err := resume.NewResumeAgent()
	if err != nil {
		log.Fatalf("简历智能体创建失败：%v", err)
	}
	// 按顺序执行
	agent := &adk.SequentialAgentConfig{
		Name:        "SequentialAgent",
		Description: "一个顺序执行的智能体，先执行简历解析智能体，再执行校招综合面试智能体",
		SubAgents: []adk.Agent{
			resumeAgent,
			schoolAgent,
		},
	}
	sequentialAgent, err := adk.NewSequentialAgent(context.Background(), agent)
	if err != nil {
		log.Fatalf("顺序智能体创建失败：%v", err)
	}
	AgentTest(sequentialAgent)
}

func AgentTest(agent adk.Agent) {
	ctx := context.Background()

	runner := adk.NewRunner(ctx, adk.RunnerConfig{
		Agent: agent,
	})

	fmt.Println("=== 智能体执行开始 ===")
	fmt.Println("输入问题进行测试：输入 'exit' 或 'quit' 结束测试")

	var fullHistory []*schema.Message

	processInput := func(input string) {
		fullHistory = append(fullHistory, schema.UserMessage(input))

		iter := runner.Run(ctx, fullHistory)

		fmt.Println("智能体回复：")

		for {
			event, ok := iter.Next()
			if !ok {
				break
			}
			if event.Err != nil {
				log.Printf("智能体执行错误：%v", event.Err)
				continue
			}
			if event.Output != nil && event.Output.MessageOutput != nil {
				msg := event.Output.MessageOutput.Message
				content := msg.Content
				if content != "" {
					fmt.Printf("%s", content)
				}
				fullHistory = append(fullHistory, msg)
			}
		}
		fmt.Println()
	}

	autoInput := "C:\\Users\\31866\\Documents\\Code\\le-code_practice-windows\\projects\\go\\go-eino-interview-agent\\backend\\chatApp\\GoTest.pdf"

	if autoInput != "" {
		fmt.Printf("自动输入：%s\n", autoInput)
		processInput(autoInput)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n 用户输入：")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("用户输入失败：%v", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		if input == "exit" || input == "quit" {
			fmt.Println("退出测试")
			break
		}
		processInput(input)
	}
}
