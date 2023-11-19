package response_helpers

import (
	"io"
	"net/http"
)

func SendWrongDataFormatResponse(response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)
	io.WriteString(response, `{"error":"Please use correct format of request"}`)
}
