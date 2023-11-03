package usecase

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

type FurnitureUsecase interface {
	SubmitFurniture(userInput, name, email, openAIKey string) (string, error)
}

type furnitureUsecase struct{}

func NewFurnitureUsecase() FurnitureUsecase {
	return &furnitureUsecase{}
}

func (uc *furnitureUsecase) SubmitFurniture(userInput, name, email, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Halo, saya adalah sistem pertanyan seputar furniture",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: name,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: email,
		},
	}

	resp, err := uc.getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func (uc *furnitureUsecase) getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}