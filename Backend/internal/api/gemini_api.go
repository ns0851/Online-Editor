package api

import (
	"context"
	"os"
	"google.golang.org/genai"
)

func GetGeminiResponse(user_prompt string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "", err
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(user_prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}
