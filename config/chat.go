package config

import (
	openai "github.com/sashabaranov/go-openai"
)

var chat *openai.Client

func GetChat() *openai.Client {
	if chat != nil {
		return chat
	}

	cfg := LoadConfig(".")
	chat = openai.NewClient(cfg.ApiKey)

	return chat
}
