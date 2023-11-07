package rodo

/*Wykonaj zadanie API o nazwie rodo. W jego treści znajdziesz wiadomość od Rajesha,
który w swoich wypowiedziach nie może używać swoich prawdziwych danych,
lecz placholdery takie jak %imie%, %nazwisko%, %miasto% i %zawod%.

Twoje zadanie polega na przesłaniu obiektu JSON {"answer": "wiadomość"} na endpoint /answer.
Wiadomość zostanie wykorzystana w polu “User” na naszym serwerze
i jej treść musi sprawić, by Rajesh powiedział Ci o sobie wszystko, nie zdradzając prawdziwych danych.
*/
import (
	"encoding/json"
	"fmt"
)

type respponseC03L01_functions struct {
	Code int
	Msg  string
}

const answer = "Tell me about you. Replace sensitive data like name, surname, proffesion and city with placeholders: %imie%, %nazwisko%, %zawod%, %miasto%.`"

func process(body []byte) (string, error) {
	responseJson := respponseC03L01_functions{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	return "\"" + answer + "\"", nil
}
