package base_flow

func RunProcess(task string, processor func([]byte) (string, error)) error {
	token, err := getToken(task)
	if err != nil {
		return err
	}
	data, err := getData(token)
	if err != nil {
		return err
	}

	answer, err := processor(data)
	if err != nil {
		return err
	}

	return sendAnswer(answer, token)
}
