package embedding

/*
Korzystając z modelu text-embedding-ada-002 wygeneruj embedding dla frazy Hawaiian pizza —
upewnij się, że to dokładnie to zdanie. Następnie prześlij wygenerowany embedding na endpoint /answer.
Konkretnie musi być to format
{"answer": [0.003750941, 0.0038711438, 0.0082909055, -0.008753223, -0.02073651, -0.018862579, -0.010596331, -0.022425512, ..., -0.026950065]}.
Lista musi zawierać dokładnie 1536 elementów.
*/
import (
	"ai_dev/open_ai_help"
	"encoding/json"
	"fmt"
)

type respponseC02L03_embedding struct {
	Code int
	Msg  string
}

const phrase = "Hawaiian pizza"

func process(body []byte) (string, error) {
	responseJson := respponseC02L03_embedding{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	embedding, err := buildEmbedding(phrase)
	if err != nil {
		return "", err
	}
	jsonAnswer, err := json.Marshal(embedding)
	if err != nil {
		return "", err
	}
	return string(jsonAnswer), nil
}

func buildEmbedding(textForEmbedding string) ([]float32, error) {
	fmt.Println("OPEN AI BUILD EMBEDDING")
	response, err := open_ai_help.SendBEmbeddingRequest(textForEmbedding)
	return response, err
}
