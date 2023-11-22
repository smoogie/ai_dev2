package optimaldb

/*
Rozwiąż zadanie API o nazwie ‘optimaldb’. Masz dostarczoną bazę danych o
rozmiarze ponad 30kb. https://zadania.aidevs.pl/data/3friends.json Musisz zoptymalizować ją w
taki sposób, aby automat korzystający z niej, a mający pojemność pamięci ustawioną na 9kb był
w stanie odpowiedzieć na 6 losowych pytań na temat trzech osób znajdujących się w bazie.
Zoptymalizowaną bazę wyślij do endpointa /answer/ jako zwykły string. Automat użyje jej jako fragment
swojego kontekstu i spróbuje odpowiedzieć na pytania testowe. Wyzwanie polega na tym, aby nie zgubić żadnej
informacji i nie zapomnieć kogo ona dotyczy oraz aby zmieścić się w wyznaczonym limicie rozmiarów bazy.
*/
import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type respponseC05L02_optimaldb struct {
	Code int
	Msg  string
}

func process(body []byte) (string, error) {
	responseJson := respponseC05L02_optimaldb{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	return prepareDatabse()
}

func prepareDatabse() (string, error) {
	fmt.Println("---GET DATA---")
	//process db file to make it smaller and return it content
	database := "{}"
	return "\"" + database + "\"", nil
}
