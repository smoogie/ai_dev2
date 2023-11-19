package ownapipro

import (
	"ai_dev/private_server_help"
	"encoding/json"
	"fmt"
)

type respponseC04L05_ownapipro struct {
	Code int
	Msg  string
}

func process(body []byte) (string, error) {
	responseJson := respponseC04L05_ownapipro{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	url, err := private_server_help.InitConversation("ownapipro")
	if err != nil {
		return "", err
	}
	return "\"" + url + "\"", nil
}
