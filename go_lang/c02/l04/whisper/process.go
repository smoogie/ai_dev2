package whisper

/*
Korzystając z modelu Whisper wykonaj zadanie API (zgodnie z opisem na zadania.aidevs.pl)
o nazwie whisper. W ramach zadania otrzymasz plik MP3 (15 sekund), który musisz wysłać do
transkrypcji, a otrzymany z niej tekst odeślij jako rozwiązanie zadania.
*/
import (
	"ai_dev/open_ai_help"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"os"
)

type respponseC02L04_whisper struct {
	Code int
	Msg  string
}

const filePath = "whisper.mp3"
const systemPrompt = `
You are support assistant that gather url for files passed in messages. you always return only catched url.
User: please return transcription of this file: https:\/\/zadania.aidevs.pl\/data\/mateusz.mp3
Assistant: https://zadania.aidevs.pl/data/mateusz.mp3
User: Can you get file http://localhost:90/tip/top
Assistant: http://localhost:90/tip/top
`

func process(body []byte) (string, error) {
	responseJson := respponseC02L04_whisper{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	fileUrl, err := getFileUrl(responseJson.Msg)
	if err != nil {
		return "", err
	}
	err = getFile(fileUrl)
	if err != nil {
		return "", err
	}
	transcript, err := getTranscript(filePath)
	if err != nil {
		return "", err
	}
	return "\"" + transcript + "\"", nil
}

func getFileUrl(textMsg string) (string, error) {
	fmt.Println("OPEN AI GET URL")
	response, err := open_ai_help.SendBasePromptRequest(systemPrompt, textMsg, openai.GPT4, true)
	return response, err
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
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, response.Body)
	return err
}

func getTranscript(fileName string) (string, error) {
	fmt.Println("OPEN AI TRANSCRIPTG")
	response, err := open_ai_help.SendWhisperRequestA(fileName)
	return response, err
}
