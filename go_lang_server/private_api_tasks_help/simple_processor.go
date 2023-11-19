package private_api_tasks_help

import (
	"ai_dev_private_api/conversation_help"
	"ai_dev_private_api/response_helpers"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func BaseRequestProcess(response http.ResponseWriter, req *http.Request, processMessage func(conversation conversation_help.PrivateApiConversation, message string) (string, error)) {
	varsRequest := mux.Vars(req)
	conversationId := varsRequest["uuid"]
	conversation, err := conversation_help.FindConversationByUUID(conversationId)
	if err != nil {
		if err == sql.ErrNoRows {
			response_helpers.SendNoConversationResponse(response)
			fmt.Println(err)
		} else {
			response_helpers.SendErrorResponse(response)
			fmt.Println(err)
		}
		return
	}
	fmt.Println("Received question for conversation: ", conversation.Uuid)
	message, err := GetMessage(req)
	if err != nil || len(message.Question) < 1 {
		response_helpers.SendWrongDataFormatResponse(response)
		fmt.Println(err)
		return
	}
	fmt.Println("Message: ", message.Question)
	answer, err := processMessage(conversation, message.Question)
	if err != nil {
		response_helpers.SendErrorResponse(response)
		fmt.Println(err)
		return
	}
	reply := BuildReply(answer)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, reply)
}
