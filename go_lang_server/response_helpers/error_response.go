package response_helpers

import (
	"io"
	"net/http"
)

func SendErrorResponse(response http.ResponseWriter){
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusInternalServerError)
	io.WriteString(response, `{"error":"Sorry Something went wrong, please try again later"}`)
}
