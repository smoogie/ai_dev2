package base_flow

import "fmt"

func RunProcessFile(task string, processor func(string) (string, error)) error {
	token, err := GetToken(task)
	if err != nil {
		return err
	}
	fileName, err := getFile(token)
	if err != nil {
		return err
	}

	fmt.Println("---PROCESS DATA--")
	answer, err := processor(fileName)
	if err != nil {
		return err
	}

	return SendAnswer(answer, token)
}
