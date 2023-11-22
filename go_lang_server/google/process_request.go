package google

import (
	"ai_dev_private_api/conversation_help"
	"ai_dev_private_api/open_ai_help"
	"ai_dev_private_api/private_api_tasks_help"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func ProcessRequest(response http.ResponseWriter, req *http.Request) {
	private_api_tasks_help.BaseRequestProcess(response, req, processMessage)
}

func processMessage(conversation conversation_help.PrivateApiConversation, message string) (string, error) {
	userMassage := conversation_help.Message{Sender: conversation_help.SenderUser, Body: message}
	conversation.History = append(conversation.History, userMassage)
	err := conversation_help.UpdateConversation(conversation)
	if err != nil {
		return "", err
	}
	response, err := getResponse(message)
	if err != nil {
		return "", err
	}
	systemMassage := conversation_help.Message{Sender: conversation_help.SenderAssistant, Body: response}
	conversation.History = append(conversation.History, systemMassage)
	err = conversation_help.UpdateConversation(conversation)
	if err != nil {
		return "", err
	}
	return response, nil
}

func getResponse(question string) (string, error) {
	fmt.Println("OPEN AI ANSWER QUESTION")
	systemPrompt := ""
	//TODO: make search and use promt to prepare response based on searched data....
	return open_ai_help.SendBasePromptRequest(systemPrompt, question, openai.GPT4, true)
}
