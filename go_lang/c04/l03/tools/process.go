package gnome

import (
	"ai_dev/open_ai_help"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type respponseC04L03_gnome struct {
	Code int
	Msg  string
	Url  string
}

const systemPrompt1 = `Validate if on the image there is a gnome in the hat. If it is gnome in the hat return "yes", if not return "no".`
const systemPrompt2 = `On the image you'll find the gnome in the hat. Return color of the hat. Response must be in Polish language`

func process(body []byte) (string, error) {
	responseJson := respponseC04L03_gnome{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	url := responseJson.Url
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	fmt.Println("response url:", url)
	return processImage(url)
}

func processImage(url string) (string, error) {
	fmt.Println("OPEN AI ANSWER QUESTION")
	response, err := open_ai_help.SendBasePromptRequest(systemPrompt1, url, openai.GPT4VisionPreview, true)
	if err != nil {
		return "", err
	}
	if response == "no" {
		return `"error"`, nil
	}
	response, err = open_ai_help.SendBasePromptRequest(systemPrompt2, url, openai.GPT4VisionPreview, true)
	return "\"" + response + "\"", err
}
