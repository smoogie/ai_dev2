package base_flow

import "fmt"

func RunProcessWithPost(task string, processor func([]byte) (string, error), getPostData func() map[string]string) error {
	token, err := getToken(task)
	if err != nil {
		return err
	}
	data, err := getDataWithPost(token, getPostData)
	if err != nil {
		return err
	}

	fmt.Println("---PROCESS DATA--")
	answer, err := processor(data)
	if err != nil {
		return err
	}

	return sendAnswer(answer, token)
}
