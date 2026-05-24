package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/flow/agent/multiagent/host"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	ctx := context.Background()
	h, err := newHost(ctx)
	if err != nil {
		panic(err)
	}
	adder, err := newAddSpecialist(ctx)
	if err != nil {
		panic(err)
	}
	subtractor, err := newSubSpecialist(ctx)
	if err != nil {
		panic(err)
	}
	hostMA, err := host.NewMultiAgent(ctx, &host.MultiAgentConfig{
		Host: *h,
		Specialists: []*host.Specialist{
			adder,
			subtractor,
		},
		Summarizer: &host.Summarizer{
			ChatModel:    h.ToolCallingModel,
			SystemPrompt: "你是一个总结者，负责总结加法专家和减法专家的回答，给用户一个最终的答案。",
		},
	})
	if err != nil {
		panic(err)
	}
	cb := &loggerCallback{}
	msg := &schema.Message{
		Role:    schema.User,
		Content: "请同时计算一下这个问题的加法和减法：123+456-78",
	}
	resp, err := hostMA.Generate(ctx, []*schema.Message{msg}, agent.WithComposeOptions(compose.WithCallbacks(cb)))
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
