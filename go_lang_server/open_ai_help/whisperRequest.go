package open_ai_help

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"os"
)

func SendWhisperRequestA(filePath string) (string, error) {
	client := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: filePath,
	}

	resp, err := client.CreateTranscription(context.Background(), req)
	if err != nil {
		return "", err
	}
	return resp.Text, nil
}
