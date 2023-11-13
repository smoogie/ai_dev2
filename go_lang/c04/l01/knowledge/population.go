package knowledge

import (
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"io/ioutil"
	"net/http"
)

const populationFunctionType = "population"

type PopulationQuery struct {
	Country string
}

func getPopulationFunctionDef() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name:        populationFunctionType,
		Description: "Run this function when question is about population.",
		Parameters: jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"country": {
					Type:        jsonschema.String,
					Description: "Country name from user question",
				},
			},
			Required: []string{"country"},
		},
	}
}

func processPopulationQuery(functionCall *openai.FunctionCall) (string, error) {
	args := functionCall.Arguments
	populationQuery := PopulationQuery{}
	err := json.Unmarshal([]byte(args), &populationQuery)
	if err != nil {
		return "", err
	}
	return getPopulation(populationQuery.Country)
}

type PopulationResponse struct {
	Population int
}

func getPopulation(country string) (string, error) {
	fmt.Println("GET POPULATION VIA API")
	fmt.Println("country:", country)
	url := "https://restcountries.com/v3.1/name/" + country + "?fields=population"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	fmt.Println("send request to:", url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	responseJson := []PopulationResponse{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	population := responseJson[0].Population
	stringValue := fmt.Sprintf("%v", population)
	return stringValue, nil
}
