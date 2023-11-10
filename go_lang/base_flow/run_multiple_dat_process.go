package base_flow

import "fmt"

func RunMultipleDataProcess(task string, processor func(string) (string, error)) error {
	token, err := GetToken(task)
	if err != nil {
		return err
	}
	fmt.Println("---PROCESS MULTIPLE DATA--")
	answer, err := processor(token)
	if err != nil {
		return err
	}

	return SendAnswer(answer, token)
}
