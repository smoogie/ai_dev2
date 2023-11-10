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

type respponseC03L05_functions struct {
	Code     int
	Msg      string
	Question string
}

type Person struct {
	Id                               int
	Imie                             string
	Nazwisko                         string
	Wiek                             int
	O_mnie                           string
	Ulubiona_postac_z_kapitana_bomby string
	Ulubiony_serial                  string
	Ulubiony_film                    string
	Ulubiony_kolor                   string
}

type PersonIdentity struct {
	FirstName string
	LastName  string
}

const questionSystemPromptStart = `
You respond to the question based on the context provided below.
Your response is in Polish language.

context ###
`
const questionSystemPromptEnd = `
###`

func process(body []byte) (string, error) {
	responseJson := respponseC03L05_functions{}
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
	personIdentity, err := getPersonIdentity(question)
	if err != nil {
		return "", err
	}
	person, err := searchForData(personIdentity.FirstName, personIdentity.LastName)
	if err != nil {
		return "", err
	}
	answer, err := getResponse(question, person)
	if err != nil {
		return "", err
	}
	return "\"" + answer + "\"", nil
}

func getPersonIdentity(question string) (PersonIdentity, error) {
	fmt.Println("OPEN AI IDENTIFY PERSON")
	functionDef := openai.FunctionDefinition{
		Name:        "get_data",
		Description: "Based on user message you identify first name and last name of person",
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
			},
			Required: []string{"lastName", "firstName"},
		},
	}
	functions := []openai.FunctionDefinition{functionDef}
	openAiResponse, functionCall, err := open_ai_help.SendFunctionCallingRequest("", question, openai.GPT4, functions, false)
	if err != nil {
		return PersonIdentity{}, err
	}
	fmt.Println("response: ", openAiResponse)
	fmt.Println(functionCall)
	args := functionCall.Arguments
	personIdentity := PersonIdentity{}
	err = json.Unmarshal([]byte(args), &personIdentity)
	if err != nil {
		return PersonIdentity{}, err
	}
	return personIdentity, err

}

func searchForData(firstName string, lastName string) (Person, error) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB"))
	defer db.Close()
	if err != nil {
		return Person{}, err
	}
	row := db.QueryRow("SELECT * FROM people WHERE imie=? AND nazwisko=?", firstName, lastName)
	var person Person
	err = row.Scan(&person.Id, &person.Imie, &person.Nazwisko, &person.Wiek, &person.O_mnie, &person.Ulubiona_postac_z_kapitana_bomby, &person.Ulubiony_serial, &person.Ulubiony_film, &person.Ulubiony_kolor)
	if err != nil {
		return Person{}, err
	}
	return person, nil
}

func getResponse(question string, person Person) (string, error) {
	fmt.Println("OPEN AI ANSWER QUESTION")
	context := "Imie: " + person.Imie + "\n"
	context += "Nazwisko: " + person.Nazwisko + "\n"
	context += "Wiek: " + string(person.Wiek) + "\n"
	context += "Info: " + person.O_mnie + "\n"
	context += "Ulubiona postc z kapitana bomby: " + person.Ulubiona_postac_z_kapitana_bomby + "\n"
	context += "Ulubiony serial: " + person.Ulubiony_serial + "\n"
	context += "Ulubiony film: " + person.Ulubiony_film + "\n"
	context += "Ulubiony kolor: " + person.Ulubiony_kolor + "\n"
	systemPrompt := questionSystemPromptStart + context + questionSystemPromptEnd
	return open_ai_help.SendBasePromptRequest(systemPrompt, question, openai.GPT4, true)
}
