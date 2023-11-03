package base_flow

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getFile(token string) (string, error) {
	fmt.Println("---GET FILE---")
	//configure request
	url := os.Getenv("API_URL") + "/task/" + token
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	//do request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	// Create the file
	out, err := os.Create(token)
	if err != nil {
		return "", err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, response.Body)
	return token, err
}
