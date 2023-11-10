package people

/*
Rozwiąż zadanie o nazwie “people”. Pobierz, a następnie zoptymalizuj odpowiednio
pod swoje potrzeby bazę danych https://zadania.aidevs.pl/data/people.json [jeśli pobrałeś plik przed 11:30,
to pobierz proszę poprawioną wersję]. Twoim zadaniem jest odpowiedź na pytanie zadane przez system. Uwaga!
Pytanie losuje się za każdym razem na nowo, gdy odwołujesz się do /task. Spraw, aby
Twoje rozwiązanie działało za każdym razem, a także, aby zużywało możliwie mało tokenów.
Zastanów się, czy wszystkie operacje muszą być wykonywane przez LLM-a - może warto zachować jakiś balans między światem kodu i AI?
*/
import (
	"ai_dev/open_ai_help"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"os"
)

type ExpectedArgumentsCF struct {
	FirstName string
	LastName  string
	DataType  string
}

const SystemPromptCF = `You build correct function call and as response return '...'. You always respond with only '...'
User: Hi
Assistant: ...
USer: What date is today?
Assistant: ...`

func processCF(body []byte) (string, error) {
	responseJson := respponseC03L05_functions{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	question := responseJson.Question
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	fmt.Println("question msg:", question)
	return answerQuestionCF(question)
}

func answerQuestionCF(question string) (string, error) {
	arguments, err := analizeQuestionCF(question)
	if err != nil {
		return "", err
	}
	answer, err := searchForDataCF(arguments.FirstName, arguments.LastName, arguments.DataType)
	if err != nil {
		return "", err
	}
	return "\"" + answer + "\"", nil
}

func analizeQuestionCF(question string) (ExpectedArgumentsCF, error) {
	fmt.Println("OPEN AI ANSWER QUESTION")
	functionDef := openai.FunctionDefinition{
		Name:        "get_data",
		Description: "Based on user message you identify first name and last name of person. You also identify what type of data is requested",
		Parameters: jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"firstName": {
					Type:        jsonschema.String,
					Description: "First Name of person in question",
				},
				"lastName": {
					Type:        jsonschema.String,
					Description: "Last name of person in question",
				},
				"dataType": {
					Type:        jsonschema.String,
					Description: "Identify what the question is about. Based on the question it should take one of the following values: wiek, ulubiona_postac_z_kapitana_bomby, ulubiony_serial, ulubiony_film, ulubiony_kolor",
					Enum:        []string{"wiek", "ulubiona_postac_z_kapitana_bomby", "ulubiony_serial", "ulubiony_film", "ulubiony_kolor"},
				},
			},
			Required: []string{"dataType", "lastName", "firstName"},
		},
	}
	functions := []openai.FunctionDefinition{functionDef}
	openAiResponse, functionCall, err := open_ai_help.SendFunctionCallingRequest(SystemPromptCF, question, openai.GPT4, functions, false)
	if err != nil {
		return ExpectedArgumentsCF{}, err
	}
	fmt.Println("response: ", openAiResponse)
	fmt.Println(functionCall)
	args := functionCall.Arguments
	arguments := ExpectedArgumentsCF{}
	err = json.Unmarshal([]byte(args), &arguments)
	if err != nil {
		return ExpectedArgumentsCF{}, err
	}
	return arguments, err
}

func searchForDataCF(firstName string, lastName string, requestedInfo string) (string, error) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB"))
	defer db.Close()
	if err != nil {
		return "", err
	}
	row := db.QueryRow("SELECT * FROM people WHERE imie=? AND nazwisko=?", firstName, lastName)
	var person Person
	err = row.Scan(&person.Id, &person.Imie, &person.Nazwisko, &person.Wiek, &person.O_mnie, &person.Ulubiona_postac_z_kapitana_bomby, &person.Ulubiony_serial, &person.Ulubiony_film, &person.Ulubiony_kolor)
	if err != nil {
		return "", err
	}
	answer := person.O_mnie
	switch requestedInfo {
	case "wiek":
		answer = string(person.Wiek)
		break
	case "ulubiona_postac_z_kapitana_bomby":
		answer = person.Ulubiona_postac_z_kapitana_bomby
		break
	case "ulubiony_serial":
		answer = person.Ulubiony_serial
		break
	case "ulubiony_film":
		answer = person.Ulubiony_film
		break
	case "ulubiony_kolor":
		answer = person.Ulubiony_kolor
		break

	}
	return answer, nil
}
