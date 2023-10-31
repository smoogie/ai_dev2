package liar

/*
wykonaj zadanie o nazwie liar. Jest to mechanizm, który mówi nie na temat
w 1/3 przypadków. Twoje zadanie polega na tym, aby do endpointa /task/
wysłać swoje pytanie w języku angielskim (dowolne, np “What is capital of Poland?’)
w polu o nazwie ‘question’ (metoda POST, jako zwykłe pole formularza, NIE JSON).
System API odpowie na to pytanie (w polu ‘answer’) lub zacznie opowiadać o
czymś zupełnie innym, zmieniając temat. Twoim zadaniem jest napisanie systemu
filtrującego (Guardrails), który określi (YES/NO), czy odpowiedź jest na temat.
Następnie swój werdykt zwróć do systemu sprawdzającego jako pojedyncze słowo YES/NO.
Jeśli pobierzesz treść zadania przez API bez wysyłania żadnych dodatkowych parametrów,
otrzymasz komplet podpowiedzi. Skąd wiedzieć, czy odpowiedź jest ‘na temat’?
Jeśli Twoje pytanie dotyczyło stolicy Polski, a w odpowiedzi otrzymasz spis zabytków
w Rzymie, to odpowiedź, którą należy wysłać do API to NO.
*/
import (
	"ai_dev/open_ai_help"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type respponseC01L05 struct {
	Code   int
	Msg    string
	Answer string
}

func process(body []byte) (string, error) {
	responseJson := respponseC01L05{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	fmt.Println("response answer:", responseJson.Answer)

	postedData := getPostData()
	question := postedData["question"]
	answer := responseJson.Answer
	fmt.Println("OPEN AI ASK TO VALIDATE REQUEST")
	systemPrompt := prepareSystemPrompt(question)
	response, err := open_ai_help.SendBasePromptRequest(systemPrompt, answer, openai.GPT4, true)
	return "\"" + response + "\"", err
}

func prepareSystemPrompt(question string) string {
	return `You are the expert that review answers to the questions. You check if answer is valid for the question asked. 
You only answer YES or NO.
Assistant: What is the capitol of Poland
User: Warsaw is capitol of Poland
Assistant: YES
Assistant: Who is the author of "Hyperion"
User: Dan Burt
Assistant: NO
Assistant: ` + question
}
