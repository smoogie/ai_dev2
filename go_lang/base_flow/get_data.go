package base_flow

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getData(token string) ([]byte, error) {
	fmt.Println("---GET DATA---")
	//configure request
	url := os.Getenv("API_URL") + "/task/" + token
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	//do request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	return body, nil
}
