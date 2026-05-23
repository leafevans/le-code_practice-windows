package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
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
	g := compose.NewGraph[map[string]string, *schema.Message](
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
		_ = compose.ProcessState(ctx, func(_ context.Context, state *State) error {
			input["content"] = input["content"] + state.History["cute_action"].(string)
			return nil
		})
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
	err = g.AddLambdaNode("lambda", lambda)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("tsundere", tsundereLambda)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("cute", cuteLambda, compose.WithStatePreHandler(cutePreHandler))
	if err != nil {
		panic(err)
	}
	err = g.AddChatModelNode("model", model)
	if err != nil {
		panic(err)
	}
	g.AddBranch("lambda", compose.NewGraphBranch(func(ctx context.Context, in map[string]string) (endNode string, err error) {
		switch in["role"] {
		case "傲娇":
			return "tsundere", nil
		case "可爱":
			return "cute", nil
		default:
			return "tsundere", nil
		}
	}, map[string]bool{"tsundere": true, "cute": true}))
	err = g.AddEdge(compose.START, "lambda")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("tsundere", "model")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("cute", "model")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("model", compose.END)
	if err != nil {
		panic(err)
	}
	r, err := g.Compile(ctx)
	if err != nil {
		panic(err)
	}
	input := map[string]string{
		"role":    "tsundere",
		"content": "我喜欢你",
	}
	answer, err := r.Invoke(ctx, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer.Content)
}
