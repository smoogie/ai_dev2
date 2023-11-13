package knowledge

/*
Wykonaj zadanie API o nazwie ‘knowledge’. Automat zada Ci losowe pytanie na temat kursu walut,
populacji wybranego kraju lub wiedzy ogólnej. Twoim zadaniem jest wybór odpowiedniego narzędzia
do udzielenia odpowiedzi (API z wiedzą lub skorzystanie z wiedzy modelu). W treści zadania uzyskanego przez API,
zawarte są dwa API, które mogą być dla Ciebie użyteczne.
*/
import (
	"ai_dev/open_ai_help"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type respponseC04L01_functions struct {
	Code     int
	Msg      string
	Question string
}

const generalFunctionType = "general"

type GeneralQuery struct {
	Response string
}

func process(body []byte) (string, error) {
	responseJson := respponseC04L01_functions{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	question := responseJson.Question
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	fmt.Println("question msg:", question)
	return answerQuestion(question)
}

func answerQuestion(question string) (string, error) {
	functionCall, response, err := identifyQuestion(question)
	if err != nil {
		return "", err
	}
	answer := ""
	switch functionCall.Name {
	case populationFunctionType:
		answer, err = processPopulationQuery(functionCall)
		break
	case currencyFunctionType:
		answer, err = processCurrencyQuery(functionCall)
		break
	case generalFunctionType:
		answer, err = processGeneralQuery(functionCall)
		break
	default:
		answer = response
		break

	}
	return "\"" + answer + "\"", err
}

func identifyQuestion(question string) (*openai.FunctionCall, string, error) {
	fmt.Println("OPEN AI IDENTIFY QUESTION")
	generalFunction := getGeneralFunctionDef()
	populationFunction := getPopulationFunctionDef()
	currencyFunction := getCurrencyFunctionDef()
	functions := []openai.FunctionDefinition{populationFunction, currencyFunction, generalFunction}
	openAiResponse, functionCall, err := open_ai_help.SendFunctionCallingRequest("", question, openai.GPT4, functions, false)
	if err != nil {
		return &openai.FunctionCall{}, "", err
	}
	fmt.Println("response: ", openAiResponse)
	fmt.Println(functionCall)
	return functionCall, openAiResponse, err

}

func getGeneralFunctionDef() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name:        generalFunctionType,
		Description: "Run this function if this is general question. If it is not question about currency or population",
		Parameters: jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"response": {
					Type:        jsonschema.String,
					Description: "Return standard assistant response for User question.",
				},
			},
			Required: []string{"response"},
		},
	}
}

func processGeneralQuery(functionCall *openai.FunctionCall) (string, error) {
	args := functionCall.Arguments
	generalQuery := GeneralQuery{}
	err := json.Unmarshal([]byte(args), &generalQuery)
	if err != nil {
		return "", err
	}
	return generalQuery.Response, nil
}
