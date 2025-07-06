package chat

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/gemini"
	"github.com/cloudwego/eino/components/model"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"

	"google.golang.org/genai"
)

func CreateGeminiChatModel(ctx context.Context) model.ToolCallingChatModel {
	key := os.Getenv("GEMINI_API_KEY")
	modelName := os.Getenv("GEMINI_MODEL_NAME")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: key})
	if err != nil {
		log.Fatal("failed to create Gemini client:", err)
	}

	cm, err := gemini.NewChatModel(ctx, &gemini.Config{
		Client: client,
		Model:  modelName,
		ThinkingConfig: &genai.ThinkingConfig{
			IncludeThoughts: true,
			ThinkingBudget:  nil,
		},
	})
	if err != nil {
		log.Fatal("failed to create Gemini chat model:", err)
	}
	return cm
}
