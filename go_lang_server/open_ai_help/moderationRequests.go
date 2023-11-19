package open_ai_help

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const ModerationEndpoint = "https://api.openai.com/v1/moderations"

type moderationResponse struct {
	Id      string
	Model   string
	Results []moderationResult
}
type moderationResult struct {
	Flagged bool
}

func SendModerationRequest(input string) (moderationResponse, error) {
	//prepare request to moderation API
	jsonData := []byte(`{
		"input": "` + input + `"
	}`)
	request, err := http.NewRequest(http.MethodPost, ModerationEndpoint, bytes.NewBuffer(jsonData))
	authorizationHeader := "Bearer " + os.Getenv("OPEN_AI_KEY")
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", authorizationHeader)
	if err != nil {
		return moderationResponse{}, err
	}
	//Send request
	client := &http.Client{}
	response, err := client.Do(request)
	fmt.Println("send request to:", ModerationEndpoint)
	if err != nil {
		return moderationResponse{}, err
	}
	defer response.Body.Close()
	// process response
	fmt.Println("response Status:", response.Status)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	responseJson := moderationResponse{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return moderationResponse{}, err
	}
	return responseJson, nil
}
