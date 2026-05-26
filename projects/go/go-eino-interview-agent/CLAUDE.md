# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

An AI-powered interview agent system built with [CloudWeGo Eino](https://github.com/cloudwego/eino) (Go AI agent framework). It uses DeepSeek models via an OpenAI-compatible API to power chat-based interview assistance and resume parsing agents.

## Module and layout

The Go module root is `backend/` (module: `ai-eino-interview-agent`, Go 1.25). All Go commands run from there.

```
backend/
├── go.mod, go.sum
├── chatApp/
│   ├── main.go              # Entry point — wires chat model, template, and streaming demo
│   ├── chat/openAI.go       # Factory for OpenAI-compatible chat model (DeepSeek-backed)
│   └── test/
│       ├── chatTemplate.go  # Prompt template helpers (system/user message formatting)
│       └── streamOut.go     # Streaming output reader/reporter
└── agent/
    └── resume/
        └── resume_agent.go  # Eino ADK agent for parsing resumes (PDF → structured JSON)
```

## Build, run, test

```bash
# Build everything
cd backend && go build ./...

# Run the main chat demo
cd backend && go run ./chatApp/

# Run a specific package
cd backend && go run ./agent/resume/
```

There are no test files yet. When added, run them with:

```bash
cd backend && go test ./...
```

## Key dependencies

- **cloudwego/eino** — Core agent framework (components, ADK, schema, prompt templates)
- **cloudwego/eino-ext/components/model/openai** — OpenAI-compatible model component; configured to call DeepSeek API

## Architecture notes

- The model is created via `chat.CreateOpenAIChatModel()` in [chat/openAI.go](backend/chatApp/chat/openAI.go). It returns a `model.ToolCallingChatModel` configured for DeepSeek (`deepseek-v4-flash`). All packages that need a chat model import and call this factory.
- Eino's ADK (`github.com/cloudwego/eino/adk`) is used to define agents — see `NewResumeAgent()` in [resume_agent.go](backend/agent/resume/resume_agent.go) for the pattern. Agents are configured with a name, instructions (system prompt), a model, tool references, and a max iteration cap.
- Prompt templates use `prompt.FromMessages()` with `schema.FString` format, supporting `{key}` placeholders filled via `template.Format(ctx, vars)`.
- Streaming chat responses use `model.Stream()` which returns `*schema.StreamReader[*schema.Message]`; iterate with `Recv()` until `io.EOF`.
