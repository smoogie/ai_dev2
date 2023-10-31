package inprompt

/*
Skorzystaj z API zadania.aidevs.pl, aby pobrać dane zadania inprompt.
Znajdziesz w niej dwie właściwości — input, czyli tablicę / listę zdań na
temat różnych osób (każde z nich zawiera imię jakiejś osoby) oraz question
będące pytaniem na temat jednej z tych osób. Lista jest zbyt duża, a
by móc ją wykorzystać w jednym zapytaniu, więc dowolną techniką odfiltruj te zdania,
które zawierają wzmiankę na temat osoby wspomnianej w pytaniu.
Ostatnim krokiem jest wykorzystanie odfiltrowanych danych jako kontekst na podstawie
którego model ma udzielić odpowiedzi na pytanie. Zatem: pobierz listę zdań oraz pytanie,
skorzystaj z LLM, aby odnaleźć w pytaniu imię, programistycznie lub z pomocą
no-code odfiltruj zdania zawierające to imię. Ostatecznie spraw by model
odpowiedział na pytanie, a jego odpowiedź prześlij do naszego API
w obiekcie JSON zawierającym jedną właściwość “answer”.
*/
import (
	"ai_dev/open_ai_help"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"strings"
)

func processHardCoded(body []byte) (string, error) {
	responseJson := respponseC02L02_inprompt{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	sentences := responseJson.Input
	question := responseJson.Question
	personInQuestion, err := identifyPersonInQuestion(question)
	if err != nil {
		return "", err
	}
	knodlwedgeText := filterKnowledgeHardcoded(sentences, personInQuestion)
	response, err := askQuestion(knodlwedgeText, question)
	if err != nil {
		return "", err
	}
	return "\"" + response + "\"", nil
}

func identifyPersonInQuestion(question string) (string, error) {
	fmt.Println("OPEN AI CLASSIFICATION")
	personInQuestion, err := open_ai_help.SendBasePromptRequest(classifySystemPrompt, question, openai.GPT4, false)
	if err != nil {
		return "", err
	}
	fmt.Println(personInQuestion + ":" + question)
	return personInQuestion, nil
}

func filterKnowledgeHardcoded(knowledge []string, personInQuestion string) string {
	filteredKnowledge := ""
	for key, val := range knowledge {
		if strings.Contains(val, personInQuestion) {
			filteredKnowledge += knowledge[key] + "."
		}
	}
	return filteredKnowledge
}
