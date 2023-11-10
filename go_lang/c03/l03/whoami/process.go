package whoami

import (
	"ai_dev/base_flow"
	"ai_dev/open_ai_help"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"strings"
	"time"
)

type respponseC03L03_whoami struct {
	Code int
	Msg  string
	Hint string
}

const MaxTries = 5
const SystemPrompt = `
Your task is to identify famous person based on hints provided by user.
!Important! You answer 'More hints' if you guess. Do not guess. If you are sure you response with "NAME:" and give me a name when you will be sure who it is.
User: Was born in England.
Assistant: More hints
User: Was prime minister of Great Britain.
Assistant: More hints
User: Resigned as prime minister in 1955.
Assistant: NAME: Sir Winston Leonard Spencer Churchill`

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
		hints := ""
		hint := ""
		token, err = base_flow.GetToken("whoami")
		if err != nil {
			break
		}

		for {
			//TODO: instead of stacking hints pass the whole conversation
			hint, err = getHint(token, start)
			if err != nil {
				break
			}
			hints += "." + hint
			response := ""
			response, err = getAnswer(hints)
			if err != nil {
				break
			}
			needMoreHints := checkAnswer(response)
			if needMoreHints {
				continue
			}
			answer = response
			break
		}

		if err != nil {
			break
		}
		if answer != "" {
			break
		}
		tries++
	}

	return answer, token, err
}

func geTimestamp() int64 {
	now := time.Now()
	return now.Unix()
}

func getHint(token string, start int64) (string, error) {
	end := geTimestamp()
	if end-start > 100 {
		return "", errors.New("timeout")
	}
	data, err := base_flow.GetData(token)
	if err != nil {
		return "", err
	}
	responseJson := respponseC03L03_whoami{}
	err = json.Unmarshal(data, &responseJson)
	if err != nil {
		return "", err
	}
	hint := responseJson.Hint
	fmt.Println("response hint:", hint)
	return hint, nil
}
func getAnswer(hints string) (string, error) {
	fmt.Println("OPEN AI ANSWER QUESTION")
	return open_ai_help.SendBasePromptRequest(SystemPrompt, hints, openai.GPT4, true)
}

func checkAnswer(response string) bool {
	needMoreHints := true
	//TODO: improve validation
	needMoreHints = !strings.Contains(response, "NAME:") || strings.Contains(response, "More hints")
	return needMoreHints
}
