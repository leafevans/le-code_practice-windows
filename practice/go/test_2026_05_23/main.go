package main

import (
	"context"
	"fmt"
	"os"

	ccb "github.com/cloudwego/eino-ext/callbacks/cozeloop"
	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"github.com/coze-dev/cozeloop-go"
)

type State struct {
	History map[string]any
}

func getFunc(ctx context.Context) *State {
	return &State{
		History: make(map[string]any),
	}
}

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic("Error loading .env file")
	}
	ctx := context.Background()
	client, err := cozeloop.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close(ctx)
	handler := ccb.NewLoopHandler(client)
	callbacks.AppendGlobalHandlers(handler)
	insideGraph := compose.NewGraph[map[string]string, *schema.Message](
		compose.WithGenLocalState(getFunc),
	)
	lambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output map[string]string, err error) {
		_ = compose.ProcessState(ctx, func(_ context.Context, state *State) error {
			state.History["tsundere_action"] = "我喜欢你"
			state.History["cute_action"] = "摸摸头"
			return nil
		})
		switch input["role"] {
		case "tsundere":
			return map[string]string{"role": "傲娇", "content": input["content"]}, nil
		case "cute":
			return map[string]string{"role": "可爱", "content": input["content"]}, nil
		default:
			return map[string]string{"role": "user", "content": input["content"]}, nil
		}
	})
	tsundereLambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output []*schema.Message, err error) {
		_ = compose.ProcessState(ctx, func(_ context.Context, state *State) error {
			input["content"] = input["content"] + state.History["tsundere_action"].(string)
			return nil
		})
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: "你是一个傲娇的女孩子，喜欢被人夸奖，但又不愿意承认。",
			},
			{
				Role:    schema.User,
				Content: input["content"],
			},
		}, nil
	})
	cuteLambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output []*schema.Message, err error) {
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: "你是一个可爱的女孩子，喜欢被人夸奖，并且会开心地接受。",
			},
			{
				Role:    schema.User,
				Content: input["content"],
			},
		}, nil
	})
	cutePreHandler := func(ctx context.Context, input map[string]string, state *State) (map[string]string, error) {
		input["content"] = input["content"] + state.History["cute_action"].(string)
		return input, nil
	}
	model, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		APIKey: os.Getenv("DEEPSEEK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	err = insideGraph.AddLambdaNode("lambda", lambda)
	if err != nil {
		panic(err)
	}
	err = insideGraph.AddLambdaNode("tsundere", tsundereLambda)
	if err != nil {
		panic(err)
	}
	err = insideGraph.AddLambdaNode("cute", cuteLambda, compose.WithStatePreHandler(cutePreHandler))
	if err != nil {
		panic(err)
	}
	err = insideGraph.AddChatModelNode("model", model)
	if err != nil {
		panic(err)
	}
	insideGraph.AddBranch("lambda", compose.NewGraphBranch(func(ctx context.Context, in map[string]string) (endNode string, err error) {
		switch in["role"] {
		case "傲娇":
			return "tsundere", nil
		case "可爱":
			return "cute", nil
		default:
			return "tsundere", nil
		}
	}, map[string]bool{"tsundere": true, "cute": true}))
	err = insideGraph.AddEdge(compose.START, "lambda")
	if err != nil {
		panic(err)
	}
	err = insideGraph.AddEdge("tsundere", "model")
	if err != nil {
		panic(err)
	}
	err = insideGraph.AddEdge("cute", "model")
	if err != nil {
		panic(err)
	}
	err = insideGraph.AddEdge("model", compose.END)
	if err != nil {
		panic(err)
	}
	outsideGraph := compose.NewGraph[map[string]string, string]()
	outLambda1 := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output map[string]string, err error) {
		return input, nil
	})
	writeLambda := compose.InvokableLambda(func(ctx context.Context, input *schema.Message) (output string, err error) {
		f, err := os.OpenFile("orc_graph_withgraph.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", err
		}
		defer f.Close()
		if _, err := f.WriteString(input.Content + "\n---\n"); err != nil {
			return "", err
		}
		return "已经写入文件，请前往文件内查看内容", nil
	})
	err = outsideGraph.AddLambdaNode("outLambda1", outLambda1)
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddGraphNode("insideGraph", insideGraph)
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddLambdaNode("writeLambda", writeLambda)
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddEdge(compose.START, "outLambda1")
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddEdge("outLambda1", "insideGraph")
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddEdge("insideGraph", "writeLambda")
	if err != nil {
		panic(err)
	}
	err = outsideGraph.AddEdge("writeLambda", compose.END)
	if err != nil {
		panic(err)
	}
	r, err := outsideGraph.Compile(ctx)
	if err != nil {
		panic(err)
	}
	input := map[string]string{
		"role":    "tsundere",
		"content": "你好鸭WWW",
	}
	answer, err := r.Invoke(ctx, input, compose.WithCallbacks(genCallback()))
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}

func genCallback() callbacks.Handler {
	handler := callbacks.NewHandlerBuilder().OnStartFn(func(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
		fmt.Printf("当前%s节点输入：%s\n", info.Component, input)
		return ctx
	}).OnEndFn(func(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
		fmt.Printf("当前%s节点输出：%s\n", info.Component, output)
		return ctx
	}).Build()
	return handler
}
