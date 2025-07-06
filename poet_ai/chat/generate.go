package chat

import (
	"context"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"log"
)

func Generate(ctx context.Context, llm model.ToolCallingChatModel, in []*schema.Message) *schema.Message {
	result, err := llm.Generate(ctx, in)
	if err != nil {
		log.Fatal("failed to generate message:", err)
	}
	return result
}

func Stream(ctx context.Context, llm model.ToolCallingChatModel, in []*schema.Message) *schema.StreamReader[*schema.Message] {
	result, err := llm.Stream(ctx, in)
	if err != nil {
		log.Fatalf("llm generate failed: %v", err)
	}
	return result
}
