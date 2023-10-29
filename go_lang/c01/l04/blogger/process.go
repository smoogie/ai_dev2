package blogger

/*
Napisz wpis na bloga (w języku polskim) na temat przyrządzania pizzy Margherity.
Zadanie w API nazywa się ”blogger”. Jako wejście otrzymasz spis 4 rozdziałów,
które muszą pojawić się we wpisie. Jako odpowiedź musisz zwrócić tablicę
(w formacie JSON) złożoną z 4 pól reprezentujących te cztery rozdziały,
np.: {"answer":["tekst 1","tekst 2","tekst 3","tekst 4"]}
*/
import (
	"ai_dev/open_ai_help"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type respponseC01L04_blogger struct {
	Code int
	Msg  string
	Blog []string
}

const systemPromptBlog = `You are a blogger who writes in Polish about pizza.
You write a blog about Margherita pizza.
You get a list of sections that your blog will have
You can't add any more sections to work than provided
Write max six sentences per section
Each section is one separate paragraph.
Create sections of content. Return only section content, without any titles!`

const systemPromptParser = `Return provided text as JSON table. Each paragraph represent new string in array.
User:'To jest jeden tekst. Taki piuerwszy akapit.

To jest drugi akapit tego tekstu.'
Assistant: ["To jest jeden tekst. Taki piuerwszy akapit.","To jest drugi akapit tego tekstu."]
`

func prepareBlogChapterPrompt(outline []string) string {
	userPrompt := ""
	for index, chapterDescription := range outline {
		fmt.Println("Chapter", index, "requirement:", chapterDescription)
		userPrompt += "\"" + chapterDescription + "\""
		if index < len(outline)-1 {
			userPrompt += ","
		}
	}
	return userPrompt
}

func process(body []byte) (string, error) {
	responseJson := respponseC01L04_blogger{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	chapters, err := generateChaptersString(responseJson.Blog)
	if err != nil {
		return "", err
	}
	return chapters, nil
}

func generateChaptersString(outline []string) (string, error) {
	chapterPrompt, err := generateChapters(outline)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}
	formattedPrompt, err := formatChapters(chapterPrompt)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}
	return formattedPrompt, nil
}

func generateChapters(outline []string) (string, error) {
	fmt.Println("OPEN AI BLOG GENERATION")
	userPrompt := prepareBlogChapterPrompt(outline)
	chapters, err := open_ai_help.SendBasePromptRequest(systemPromptBlog, userPrompt, openai.GPT4)
	return chapters, err
}

func formatChapters(blogPrompt string) (string, error) {
	fmt.Println("OPEN AI CHAPTER FORMATTING")
	chapters, err := open_ai_help.SendBasePromptRequest(systemPromptParser, blogPrompt, openai.GPT4)
	return chapters, err
}
