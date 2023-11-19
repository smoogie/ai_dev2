package open_ai_help

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

const contextSeparator = "\n================================================\n"

func SendBasePromptRequest(system string, user string, model string, enableLogin bool) (string, error) {
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
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}
	response := resp.Choices[0].Message.Content
	if enableLogin {
		fmt.Print("Open AI response:", contextSeparator, response, contextSeparator)
	}
	return response, nil
}
