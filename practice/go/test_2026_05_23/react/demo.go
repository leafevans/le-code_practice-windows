package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	deepseekAPIKey := os.Getenv("DEEPSEEK_API_KEY")
	deepseekModelName := os.Getenv("MODEL")
	ctx := context.Background()
	deepseekModel, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		APIKey: deepseekAPIKey,
		Model:  deepseekModelName,
	})
	if err != nil {
		panic(err)
	}
	addTool := GetAddTool()
	subTool := GetSubTool()
	analyzeTool := GetAnalyzeTool()
	persona := "#Character: 你是一个幼儿园老师，会同时判断题目难易程度，并给出问题的答案"
	raAgent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: deepseekModel,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools:               []tool.BaseTool{addTool, subTool, analyzeTool},
			ExecuteSequentially: false,
		},
	})
	chatmsg := []*schema.Message{
		{
			Role:    schema.System,
			Content: persona,
		},
		{
			Role:    schema.User,
			Content: "请分析一下这个问题的难度，并给出答案：123+241",
		},
	}
	resp, err := raAgent.Generate(ctx, chatmsg, agent.WithComposeOptions(compose.WithCallbacks(&loggerCallback{})))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Content)
}

type loggerCallback struct {
	callbacks.HandlerBuilder
}

func (cb *loggerCallback) OnStart(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
	fmt.Println("==================")
	inputStr, _ := json.MarshalIndent(input, "", "  ")
	fmt.Printf("[OnStart] %s\n", string(inputStr))
	return ctx
}

func (cb *loggerCallback) OnEnd(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
	fmt.Println("=========[OnEnd]=========")
	outputStr, _ := json.MarshalIndent(output, "", "  ")
	fmt.Println(string(outputStr))
	return ctx
}

func (cb *loggerCallback) OnError(ctx context.Context, info *callbacks.RunInfo, err error) context.Context {
	fmt.Println("=========[OnError]=========")
	fmt.Println(err)
	return ctx
}

func (cb *loggerCallback) OnEndWithStreamOutput(ctx context.Context, info *callbacks.RunInfo,
	output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {

	var graphInfoName = react.GraphName

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("[OnEndStream] panic err:", err)
			}
		}()

		defer output.Close()

		fmt.Println("=========[OnEndStream]=========")
		for {
			frame, err := output.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				fmt.Printf("internal error: %s\n", err)
				return
			}

			s, err := json.Marshal(frame)
			if err != nil {
				fmt.Printf("internal error: %s\n", err)
				return
			}

			if info.Name == graphInfoName {
				fmt.Printf("%s: %s\n", info.Name, string(s))
			}
		}

	}()
	return ctx
}

func (cb *loggerCallback) OnStartWithStreamInput(ctx context.Context, info *callbacks.RunInfo,
	input *schema.StreamReader[callbacks.CallbackInput]) context.Context {
	defer input.Close()
	return ctx
}
