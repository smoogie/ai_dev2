package open_ai_help

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func SendEmbeddingRequest(textForEmbedding string) ([]float32, error) {
	client := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	queryReq := openai.EmbeddingRequest{
		Input: []string{textForEmbedding},
		Model: openai.AdaEmbeddingV2,
	}

	// Create an embedding for the user query
	queryResponse, err := client.CreateEmbeddings(context.Background(), queryReq)
	if err != nil {
		fmt.Printf("Embedding error: %v\n", err)
		return make([]float32, 1), err
	}
	queryEmbedding := queryResponse.Data[0].Embedding
	return queryEmbedding, nil
}
