package api

import (
	"context"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"google.golang.org/genai"
)

func GetGeminiResponse(user_prompt string) ([]string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return []string{}, err
	}

	before_prompt := "Analyze this script and return a JSON array of appropriate number of search queries. The search queries should be different based on the script with different keywords for every scene or narration change. Make sure that characters stay consistent like white boy should stay white throughout etc. The search queries should be concise for a stock video site like Pexels. Each query in the array should be a string of 2-3 keywords. Only return the JSON array and nothing else. DO NOT INCLUDE ANY BACKTICKS (`) OR ANY OTHER TEXT LIKE JSON ETC ONLY JSON ARRAY \n\n"

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(before_prompt+user_prompt),
		nil,
	)
	if err != nil {
		return []string{}, err
	}

	fmt.Printf("Successfully got response from Gemini: %s\n", result.Text())

	cleanString := strings.TrimPrefix(result.Text(), "```json\n")
	cleanString = strings.TrimSuffix(cleanString, "\n```")
	// cleanString = strings.TrimSpace(cleanString)

	fmt.Println(cleanString)

	var arr []string
	erre := json.Unmarshal([]byte(cleanString), &arr)
	if erre != nil {
		panic(erre)
	}

	fmt.Println("this is yayaya", arr)

	fmt.Println("actual lenght brooooooo", len(arr))

	return arr, nil
}
