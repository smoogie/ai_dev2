package tools

/*
Rozwiąż zadanie API o nazwie ‘tools’. Celem zadania jest zdecydowanie, czy podane przez API
zadanie powinno zostać dodane do listy zadań (ToDo), czy do kalendarza (jeśli ma ustaloną datę).
Oba narzędzia mają lekko definicje struktury JSON-a (różnią się jednym polem).
Spraw, aby Twoja aplikacja działała poprawnie na każdym zestawie danych testowych.
*/
import (
	"ai_dev/open_ai_help"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"time"
)

type respponseC04L02_tools struct {
	Code     int
	Msg      string
	Question string
}

func process(body []byte) (string, error) {
	responseJson := respponseC04L02_tools{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	question := responseJson.Question
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	fmt.Println("response que:", question)
	return getJsonForTool(question)
}

func getJsonForTool(question string) (string, error) {
	functionCall, err := identifyTool(question)
	if err != nil {
		return "", err
	}
	toolResponse := ""
	switch functionCall.Name {
	case CalendarType:
		toolResponse, err = processCalendarQuery(functionCall)
		break
	case ToDoType:
		toolResponse, err = processToDoQuery(functionCall)
		break
	}

	return toolResponse, nil
}

func identifyTool(question string) (*openai.FunctionCall, error) {
	fmt.Println("OPEN AI IDENTIFY QUESTION")
	toDoFunction := getToDoFunctionDef()
	calendarFunction := getCalendarFunctionDef()
	functions := []openai.FunctionDefinition{toDoFunction, calendarFunction}
	currentTime := time.Now()
	systemPrompt := `You are automatization assistant. You receive question about the task and select correct function to run. Today is ` + currentTime.Format("2006-01-02 Monday")
	fmt.Println(systemPrompt)
	openAiResponse, functionCall, err := open_ai_help.SendFunctionCallingRequest(systemPrompt, question, openai.GPT4, functions, false)
	if err != nil {
		return &openai.FunctionCall{}, err
	}
	fmt.Println("response: ", openAiResponse)
	fmt.Println(functionCall)
	return functionCall, err
}
