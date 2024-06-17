package openai_api

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIAPI struct {
	APIKey string
}

func NewOpenAIAPI(APIKey string) *OpenAIAPI {
	return &OpenAIAPI{
		APIKey: APIKey,
	}
}

func (o *OpenAIAPI) GetChatCompletion(prompt []string, userPrompt string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(o.APIKey)

	var chatMessages []openai.ChatCompletionMessage
	chatMessages = append(chatMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "Anda adalah seorang customer service AI dari sebuah aplikasi bernama KeluhProv Banten, yang merupakan aplikasi pengaduan masyarakat di Provinsi Banten",
	})

	for _, p := range prompt {
		chatMessages = append(chatMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: p,
		})
	}

	if userPrompt != "" {
		chatMessages = append(chatMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userPrompt,
		})
	}

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: chatMessages,
	}

	// Get response
	resp, err := client.CreateChatCompletion(ctx, req)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
