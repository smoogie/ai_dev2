package base_flow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type answerResponse struct {
	Code int
	Msg  string
	Note string
}

func sendAnswer(answer string, token string) error {
	//Configure request
	url := os.Getenv("API_URL") + "/answer/" + token
	jsonData := []byte(`{
		"answer": ` + answer + `
	}`)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return err
	}
	//Send request
	client := &http.Client{}
	response, err := client.Do(request)
	fmt.Println("send request to:", url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	// process response
	fmt.Println("response Status:", response.Status)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	//parse response
	responseJson := answerResponse{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return err
	}
	fmt.Println("response code:", string(responseJson.Code))
	fmt.Println("response msg:", string(responseJson.Msg))
	fmt.Println("response note:", string(responseJson.Note))
	return nil
}
