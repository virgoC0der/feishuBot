package services

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"go.uber.org/zap"

	"feishuBot/internal/conf"
	"feishuBot/utils/logger"
)

var (
	client openai.Client
)

func InitOpenAI() {
	client = openai.NewClient(
		option.WithAPIKey(conf.GConfig.LLM.ApiKey),
		option.WithBaseURL(conf.GConfig.LLM.BaseUrl),
	)
}

// CallDeepSeekAPI sends a request to the DeepSeek API and returns the response
func CallDeepSeekAPI(message string) (string, error) {
	chatCompletion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		},
		Model:       openai.ChatModel(conf.GConfig.LLM.Model),
		Temperature: openai.Float(0.7),
	})

	if err != nil {
		logger.Error("chat with deepseek failed", zap.Error(err))
		return "", err
	}

	return chatCompletion.Choices[0].Message.Content, nil
}
