package base_flow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type tokenResponse struct {
	Code  int
	Msg   string
	Token string
}

func GetToken(task string) (string, error) {
	fmt.Println("---GET TOKEN---")
	//Configure request
	apiKey := os.Getenv("API_KEY")
	url := os.Getenv("API_URL") + "/token/" + task
	jsonData := []byte(`{
		"apikey": "` + apiKey + `"
	}`)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return "", err
	}
	//Send request
	client := &http.Client{}
	response, err := client.Do(request)
	fmt.Println("send request to:", url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	// process response
	fmt.Println("response Status:", response.Status)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	//parse response
	responseJson := tokenResponse{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)

	return responseJson.Token, nil
}
