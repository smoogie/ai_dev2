package base_flow

import "fmt"

func RunProcess(task string, processor func([]byte) (string, error)) error {
	token, err := GetToken(task)
	if err != nil {
		return err
	}
	data, err := GetData(token)
	if err != nil {
		return err
	}

	fmt.Println("---PROCESS DATA--")
	answer, err := processor(data)
	if err != nil {
		return err
	}

	return SendAnswer(answer, token)
}
