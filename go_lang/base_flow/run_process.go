package base_flow

import "fmt"

func RunProcess(task string, processor func([]byte) (string, error)) error {
	token, err := getToken(task)
	if err != nil {
		return err
	}
	data, err := getData(token)
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
