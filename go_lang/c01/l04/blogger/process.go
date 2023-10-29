package blogger

/*
Napisz wpis na bloga (w języku polskim) na temat przyrządzania pizzy Margherity.
Zadanie w API nazywa się ”blogger”. Jako wejście otrzymasz spis 4 rozdziałów,
które muszą pojawić się we wpisie. Jako odpowiedź musisz zwrócić tablicę
(w formacie JSON) złożoną z 4 pól reprezentujących te cztery rozdziały,
np.: {"answer":["tekst 1","tekst 2","tekst 3","tekst 4"]}
*/
import (
	"encoding/json"
	"fmt"
)

type respponseC01L04_blogger struct {
	Code int
	Msg  string
}

func process(body []byte) (string, error) {
	responseJson := respponseC01L04_blogger{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", string(responseJson.Code))
	fmt.Println("response msg:", string(responseJson.Msg))
	return "", nil
}
