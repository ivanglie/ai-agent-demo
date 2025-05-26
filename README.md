# AI Agent Demo with DeepSeek

Go application demonstrating tool calling with DeepSeek's API using the OpenAI Go client library.

## Features

- DeepSeek API integration with tool calling
- Custom time retrieval function

## Prerequisites

- Go 1.24.0+
- DeepSeek API key

## Installation

```bash
git clone https://github.com/ivanglie/ai-agent-demo.git
cd ai-agent-demo
go mod download
```

Set up API key:
```bash
cp .env.example .env
# Edit .env and add your real API key
```

## Usage

```bash
make run
```

## How it works

1. Sends request to DeepSeek asking for current time
2. AI calls the `GetTime` tool function
3. Returns formatted time response with slang

## Dependencies

- [github.com/sashabaranov/go-openai v1.40.0](https://github.com/sashabaranov/go-openai)

## License

MIT