package services

import (
	"context"
	"customer_service_gpt/config"

	"github.com/sashabaranov/go-openai"
)

type GPTService struct {
	client *openai.Client
}

func NewGPTService(config *config.Config) *GPTService {
	client := openai.NewClient(config.GPTAPIKey)
	return &GPTService{client: client}
}

func (s *GPTService) GetResponse(message string) (string, error) {
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
