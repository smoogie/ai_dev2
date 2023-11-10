package scraper

/*Rozwiąż zadanie z API o nazwie "scraper". Otrzymasz z API link do artykułu (format TXT),
który zawiera pewną wiedzę, oraz pytanie dotyczące otrzymanego tekstu.
Twoim zadaniem jest udzielenie odpowiedzi na podstawie artykułu.
Trudność polega tutaj na tym, że serwer z artykułami działa naprawdę kiepsko —
w losowych momentach zwraca błędy typu "error 500", czasami odpowiada bardzo wolno na
Twoje zapytania, a do tego serwer odcina dostęp nieznanym przeglądarkom internetowym.
Twoja aplikacja musi obsłużyć każdy z napotkanych błędów. Pamiętaj, że pytania,
jak i teksty źródłowe, są losowe, więc nie zakładaj, że uruchamiając aplikację kilka razy,
za każdym razem zapytamy Cię o to samo i będziemy pracować na tym samym artykule.


{
	"code": 0,
	"msg": "Return answer for the question in POLISH language, based on provided article. Maximum length for the answer is 200 characters",
	"input": "https:\/\/zadania.aidevs.pl\/text_pizza_history.txt",
	"question": "z którego roku pochodzi łaciński dokument, który pierwszy raz wspomina o pizzy? "
}
*/
import (
	"ai_dev/base_flow"
	"ai_dev/open_ai_help"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type respponseC03L02_functions struct {
	Code     int
	Msg      string
	Input    string
	Question string
}

const MaxTries = 5
const SavedFilePath = "receivedFile"

const questionSystemPromptStart = `
You respond to the question based on the context provided below.
Your response is in Polish language.

context ###
`
const questionSystemPromptEnd = `
###`

func process() error {
	answer, token, err := processLoop()
	if err != nil {
		return err
	}
	answer = strings.Replace(answer, "\"", "", -1)
	answer = strings.Replace(answer, "\n", "", -1)
	return base_flow.SendAnswer("\""+answer+"\"", token)
}

func processLoop() (string, string, error) {
	var err error
	err = nil
	token := ""
	answer := ""
	tries := 0

	for tries < MaxTries {
		start := geTimestamp()
		token, err = base_flow.GetToken("scraper")
		if err != nil {
			break
		}
		var data []byte
		data, err = base_flow.GetData(token)
		if err != nil {
			break
		}
		responseJson := respponseC03L02_functions{}
		err = json.Unmarshal(data, &responseJson)
		if err != nil {
			break
		}
		filePath := responseJson.Input
		question := responseJson.Question
		fmt.Println("response code:", responseJson.Code)
		fmt.Println("response msg:", responseJson.Msg)
		fmt.Println("response input:", filePath)
		fmt.Println("response question:", question)

		context := ""
		context, err = getContextFromFile(filePath, start)
		if err != nil {
			continue
		}

		answer, err = getAnswer(context, question)
		if err != nil {
			break
		}
		end := geTimestamp()
		if end-start < 110 {
			break
		}
		tries++
	}

	return answer, token, nil
}

func geTimestamp() int64 {
	now := time.Now()
	return now.Unix()
}

func getContextFromFile(fileUrl string, start int64) (string, error) {
	var err error
	err = nil

	//TODO: fix "bot detected" issue
	fmt.Println("GET FILE")
	for {
		err = getFile(fileUrl)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			end := geTimestamp()
			if end-start < 80 {
				fmt.Println("TRY AGAIN")
				continue
			}
		}
		break
	}
	if err != nil {
		return "", err
	}
	end := geTimestamp()
	if end-start > 95 {
		return "", errors.New("timeout")
	}
	content, err := os.ReadFile(SavedFilePath)
	if err != nil {
		return "", err
	}
	fmt.Println(string(content))
	return string(content), nil
}

func getFile(fileUrl string) error {
	req, err := http.NewRequest(http.MethodGet, fileUrl, nil)
	if err != nil {
		return err
	}
	//do request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	// Create the file
	out, err := os.Create(SavedFilePath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, response.Body)
	return err
}

func getAnswer(context string, question string) (string, error) {
	fmt.Println("OPEN AI ANSWER QUESTION")
	systemPrompt := questionSystemPromptStart + context + questionSystemPromptEnd
	return open_ai_help.SendBasePromptRequest(systemPrompt, question, openai.GPT4, true)
}
