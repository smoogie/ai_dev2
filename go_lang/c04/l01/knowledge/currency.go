package knowledge

import (
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"io/ioutil"
	"net/http"
)

const currencyFunctionType = "currency"

type CurrencyQuery struct {
	Currency string
}

func getCurrencyFunctionDef() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name:        currencyFunctionType,
		Description: "Run this question when question is about currency rate.",
		Parameters: jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"currency": {
					Type:        jsonschema.String,
					Description: "Currency from user question. Currency must be returned as three letters in ISO 4217 standard. For example USD, PLN, EUR",
				},
			},
			Required: []string{"currency"},
		},
	}
}

func processCurrencyQuery(functionCall *openai.FunctionCall) (string, error) {
	args := functionCall.Arguments
	currencyQuery := CurrencyQuery{}
	err := json.Unmarshal([]byte(args), &currencyQuery)
	if err != nil {
		return "", err
	}
	return getCurrencyExchange(currencyQuery.Currency)
}

type CurrencyRate struct {
	No            string
	EffectiveDate string
	Mid           float32
}

type CurrencyResponse struct {
	Table    string
	Currency string
	Code     string
	Rates    []CurrencyRate
}

func getCurrencyExchange(currency string) (string, error) {
	fmt.Println("GET CURRENCY VIA API")
	fmt.Println("currency:", currency)
	url := "https://api.nbp.pl/api/exchangerates/rates/A/" + currency + "/?format=json"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
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
	responseJson := CurrencyResponse{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	midValue := responseJson.Rates[0].Mid
	stringValue := fmt.Sprintf("%v", midValue)
	return stringValue, nil
}
