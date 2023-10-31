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
)

type respponseC02L02_inprompt struct {
	Code     int
	Msg      string
	Input    []string
	Question string
}

const questionSystemPromptStart = `
You respond to the question based on the context provided below.

context ###
`
const questionSystemPromptEnd = `
###`
const classifySystemPrompt = `
You are the clasificator of the sentence - that identify name of the person about whom the text is.
You always response and return only the name of the person. Follow below examples:
User: "Many people liked Fred. As he was realy nice person"
Assistant: Fred
User: "Kahut eat many fruits. He liked them so much that he had own garden."
Assistant: Kahut
User: "Well you know the family. Bob realy loved his kids. HE had son Jon and daugther Kaya."
Assistant: Bob
`

func process(body []byte) (string, error) {
	responseJson := respponseC02L02_inprompt{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	sentences := responseJson.Input
	question := responseJson.Question
	people, personInQuestion, err := identifyPeople(sentences, question)
	if err != nil {
		return "", err
	}
	knowledgeText := filterKnowledge(sentences, people, personInQuestion)
	response, err := askQuestion(knowledgeText, question)
	if err != nil {
		return "", err
	}
	return "\"" + response + "\"", nil
}

func identifyPeople(sentences []string, question string) ([]string, string, error) {
	fmt.Println("OPEN AI CLASSIFICATION")
	people := make([]string, len(sentences))
	for key, val := range sentences {
		name, err := open_ai_help.SendBasePromptRequest(classifySystemPrompt, val, openai.GPT4, false)
		if err != nil {
			return make([]string, 0), "", err
		}
		people[key] = name
		fmt.Println(name + ":" + val)
	}
	personInQuestion, err := open_ai_help.SendBasePromptRequest(classifySystemPrompt, question, openai.GPT4, false)
	if err != nil {
		return people, "", err
	}
	fmt.Println(personInQuestion + ":" + question)
	return people, personInQuestion, nil
}

func filterKnowledge(knowledge []string, people []string, personInQuestion string) string {
	filteredKnowledge := ""
	for key, val := range people {
		if val == personInQuestion {
			filteredKnowledge += knowledge[key] + "."
		}
	}
	return filteredKnowledge
}

func askQuestion(knowledgeText string, question string) (string, error) {
	fmt.Println("OPEN AI ANSWER QUESTION")
	systemPrompt := questionSystemPromptStart + knowledgeText + questionSystemPromptEnd
	response, err := open_ai_help.SendBasePromptRequest(systemPrompt, question, openai.GPT4, true)
	return response, err
}
