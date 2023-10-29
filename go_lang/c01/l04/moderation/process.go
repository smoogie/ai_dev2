package moderation

/*
Zastosuj wiedzę na temat działania modułu do moderacji treści
i rozwiąż zadanie o nazwie “moderation” z użyciem naszego API
do sprawdzania rozwiązań. Zadanie polega na odebraniu tablicy
zdań (4 sztuki), a następnie zwróceniu tablicy z informacją,
które zdania nie przeszły moderacji. Jeśli moderacji nie przeszło
pierwsze i ostatnie zdanie, to odpowiedź powinna brzmieć [1,0,0,1].
Pamiętaj, aby w polu ‘answer’ zwrócić tablicę w JSON,
a nie czystego stringa.
*/
import (
	"ai_dev/open_ai_help"
	"encoding/json"
	"fmt"
)

type respponseC01L04_moderation struct {
	Code  int
	Msg   string
	Input []string
}

func process(body []byte) (string, error) {
	responseJson := respponseC01L04_moderation{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", string(responseJson.Code))
	fmt.Println("response msg:", responseJson.Msg)
	flagsForInputs := make([]int, len(responseJson.Input))
	for index, input := range responseJson.Input {
		fmt.Println("response input[", index, "]:", input)
		flag, errReq := validateInput(input)
		if errReq != nil {
			err = errReq
			break
		}
		fmt.Println("response input[", index, "] flag:", flag)
		flagsForInputs[index] = flag
	}
	if err != nil {
		return "", err
	}
	answer, _ := json.Marshal(flagsForInputs)
	fmt.Println("Answer prepared", string(answer))
	return string(answer), nil
}

func validateInput(input string) (int, error) {
	flag := 1
	moderationResponse, err := open_ai_help.SendModerationRequest(input)
	if err != nil {
		return 1, err
	}
	if moderationResponse.Results[0].Flagged {
		flag = 1
	} else {
		flag = 0
	}
	return flag, nil
}
