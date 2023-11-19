package private_api_tasks_help

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Question string
}

func GetMessage(req *http.Request) (Message, error) {
	responseJson := Message{}
	err := json.NewDecoder(req.Body).Decode(&responseJson)
	if err != nil {
		return Message{}, err
	}
	return responseJson, nil
}
