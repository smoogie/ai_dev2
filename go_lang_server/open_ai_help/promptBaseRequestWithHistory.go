package open_ai_help

import (
	"ai_dev_private_api/conversation_help"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func SendBasePromptRequestWithHistory(history conversation_help.History, system string, user string, model string, enableLogin bool) (string, error) {
	//prepare prompt
	if enableLogin {
		fmt.Print("System prompt:", contextSeparator, system, contextSeparator)
		fmt.Print("User prompt:", contextSeparator, user, contextSeparator)
	}
	messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleSystem, Content: system}}
	for _, msg := range history {
		switch msg.Sender {
		case conversation_help.SenderAssistant:
			messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: msg.Body})
			break
		case conversation_help.SenderUser:
			messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: msg.Body})
			break
		}
	}
	messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: user})
	//send request to open ai
	client := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
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
