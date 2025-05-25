package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
)

const (
	baseURL      = "https://api.deepseek.com/v1"
	defaultModel = "deepseek-chat"
)

func onGetTime() string {
	return "current time:" + time.Now().Format(time.RFC3339)
}

func main() {
	ctx := context.Background()

	config := openai.DefaultConfig(os.Getenv("DEEPSEEK_API_KEY"))
	config.BaseURL = baseURL
	client := openai.NewClientWithConfig(config)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Answer each question using slang. Always display time in the format: 13:05 (thirteen hours five minutes).",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "What time is it now?",
		},
	}

	toolGetTime := openai.Tool{
		Type: "function",
		Function: &openai.FunctionDefinition{
			Name:        "GetTime",
			Description: "Get the current date and time. The returned time string format is RFC3339.",
			Parameters: map[string]interface{}{
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
	}

	firstRequest := openai.ChatCompletionRequest{
		Model:      defaultModel,
		Messages:   messages,
		Tools:      []openai.Tool{toolGetTime},
		ToolChoice: "required",
	}

	firstResponse, err := client.CreateChatCompletion(ctx, firstRequest)
	if err != nil {
		log.Fatalf("error on first request: %v", err)
	}

	msg := firstResponse.Choices[0].Message

	log.Printf("tool calls: %v\n", msg.ToolCalls)

	messages = append(messages, openai.ChatCompletionMessage{
		Role:      openai.ChatMessageRoleAssistant,
		Content:   msg.Content,
		ToolCalls: msg.ToolCalls,
	})

	if len(msg.ToolCalls) > 0 {
		for _, toolCall := range msg.ToolCalls {
			var result string

			switch toolCall.Function.Name {
			case "GetTime":
				result = onGetTime()

				messages = append(messages, openai.ChatCompletionMessage{
					Role:       openai.ChatMessageRoleTool,
					Content:    result,
					ToolCallID: toolCall.ID,
				})
			default:
				result = "unknown tool"
				log.Print(result)
			}
		}
	}

	secondRequest := openai.ChatCompletionRequest{
		Model:    defaultModel,
		Messages: messages,
	}

	secondResponse, err := client.CreateChatCompletion(ctx, secondRequest)
	if err != nil {
		log.Fatalf("error on second request: %v", err)
	}

	log.Printf("response: %s", secondResponse.Choices[0].Message.Content)
}
