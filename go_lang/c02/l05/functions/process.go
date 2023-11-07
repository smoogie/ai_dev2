package functions

/*
Wykonaj zadanie o nazwie functions zgodnie ze standardem zgłaszania odpowiedzi
opisanym na zadania.aidevs.pl. Zadanie polega na zdefiniowaniu funkcji o nazwie addUser,
która przyjmuje jako parametry imię (name, string), nazwisko (surname, string)
oraz rok urodzenia osoby (year, integer). Jako odpowiedź musisz wysłać jedynie ciało funkcji w postaci JSON-a.
Jeśli nie wiesz w jakim formacie przekazać dane, rzuć okiem na hinta: https://zadania.aidevs.pl/hint/functions
*/
import (
	"encoding/json"
	"fmt"
)

type respponseC02L05_functions struct {
	Code int
	Msg  string
}

func process(body []byte) (string, error) {
	responseJson := respponseC02L05_functions{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	return getFunctionObject()
}

func getFunctionObject() (string, error) {
	functionResponse := `{
    "name": "addUser",
    "description": "Add user to the system based on passed parameter",
    "parameters": {
        "type": "object",
        "properties": {
            "name": {
                "type": "string",
                "description": "user name"
            },
            "surname": {
                "type": "string",
                "description": "User surname"
            },
            "year": {
                "type": "integer",
                "description": "The number that represents year of the birth date"
            }
        }
    }
}`
	return functionResponse, nil
}
