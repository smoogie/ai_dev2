package base_flow

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func createForm(form map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()
	for key, val := range form {
		if strings.HasPrefix(val, "@") {
			val = val[1:]
			file, err := os.Open(val)
			if err != nil {
				return "", nil, err
			}
			defer file.Close()
			part, err := mp.CreateFormFile(key, val)
			if err != nil {
				return "", nil, err
			}
			io.Copy(part, file)
		} else {
			mp.WriteField(key, val)
		}
	}
	return mp.FormDataContentType(), body, nil
}

func getDataWithPost(token string, getPostData func() map[string]string) ([]byte, error) {
	//prepare data
	toPost := getPostData()
	ct, body, err := createForm(toPost)
	if err != nil {
		return nil, err
	}
	//configure request
	url := os.Getenv("API_URL") + "/task/" + token
	req, err := http.NewRequest(http.MethodPost, url, body)
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", ct)
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
	resBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(resBody))
	return resBody, nil
}
