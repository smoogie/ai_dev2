package google

/*
Rozwiąż zadanie API o nazwie ‘google’. Do jego wykonania będziesz potrzebować darmowego konta w
usłudze SerpAPI. Celem zadania jest samodzielne zaimplementowanie rozwiązania podobnego do tego,
znanego z ChatGPT Plus, gdzie po wpisaniu zapytania na temat, o którym model nie ma pojęcia,
uruchamiana jest wyszukiwarka BING. My posłużymy się wyszukiwarką Google, a Twój skrypt będzie wyszukiwał
odpowiedzi na pytania automatu sprawdzającego i będzie zwracał je w czytelnej dla człowieka formie.
Więcej informacji znajdziesz w treści zadania /task/, a podpowiedzi dostępne
są pod https://zadania.aidevs.pl/hint/google.
*/
import (
	"ai_dev/private_server_help"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type respponseC05L03_google struct {
	Code int
	Msg  string
}

func process(body []byte) (string, error) {
	responseJson := respponseC05L03_google{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	url, err := private_server_help.InitConversation("google")
	if err != nil {
		return "", err
	}
	return "\"" + url + "\"", nil
}
