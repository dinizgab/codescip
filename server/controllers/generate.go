package controllers

import (
	"codescip/config"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
)

type Response struct {
	TextResponse string
}

func GenerateDocumentation(c echo.Context) error {
	client := config.GetChat()

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Give me a random golang fact!", // TODO - Trocar o content da mensagem pra fazer as documentações
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	return c.JSON(http.StatusOK, resp.Choices[0].Message.Content)
}
