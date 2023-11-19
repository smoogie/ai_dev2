package response_helpers

import (
	"io"
	"net/http"
)

func SendNoConversationResponse(response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusNotFound)
	io.WriteString(response, `{"error":"Wrong Conversation UUID"}`)
}
