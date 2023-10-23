package lo1

import (
	"encoding/json"
	"fmt"
)

type respponseC01L01 struct {
	Code   int
	Msg    string
	Cookie string
}

func process(body []byte) (string, error) {
	responseJson := respponseC01L01{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", string(responseJson.Code))
	fmt.Println("response msg:", string(responseJson.Msg))
	fmt.Println("response cookie:", string(responseJson.Cookie))
	return string(responseJson.Cookie), nil
}
