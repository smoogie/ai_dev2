package open_ai_help

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func SendFunctionCallingRequest(system string, user string, model string, functions []openai.FunctionDefinition, enableLogin bool) (string, *openai.FunctionCall, error) {
	//prepare prompt
	if enableLogin {
		fmt.Print("System prompt:", contextSeparator, system, contextSeparator)
		fmt.Print("User prompt:", contextSeparator, user, contextSeparator)
	}
	//send request to open ai
	client := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: system,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: user,
				},
			},
			Functions: functions,
		},
	)

	if err != nil {
		fCall := openai.FunctionCall{}
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", &fCall, err
	}
	response := resp.Choices[0].Message.Content
	functionCall := resp.Choices[0].Message.FunctionCall
	if enableLogin {
		fmt.Print("Open AI response:", contextSeparator, response, contextSeparator)
		fmt.Print(functionCall)
	}
	return response, functionCall, nil
}
