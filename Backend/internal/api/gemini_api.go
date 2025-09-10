package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"google.golang.org/genai"
)

type SceneQuery struct {
	Query string `json:"query"`
	Type  string `json:"type"`
	URL   string `json:"url,omitempty"` // store Pexels URL
}

func GetGeminiResponse(user_prompt string) ([]SceneQuery, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	before_prompt := "Analyze this script and return a JSON array of objects. Return as loads of scenes as required... the edit should be fast paced with variety of visuals or atleast have variations so the video remains fresh for edit. Each object should represent a scene and contain two keys: - 'query': a concise search query (3-5 keywords) for a stock site like Pexels, keywords should make sure some consistency remains not just describe the exact mood. - 'type': either 'image' or 'video', depending on whether the scene would be best represented with a still image or a video which can be available in website like pexels. Make sure that characters stay consistent (e.g., if a white boy appears, he stays white throughout). Do not include any explanations, text, or formatting other than the JSON array. DO NOT INCLUDE ANY BACKTICKS (`) OR EXTRA TEXT. ONLY RETURN THE JSON ARRAY."

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(before_prompt+user_prompt),
		nil,
	)
	if err != nil {
		return nil, err
	}

	fmt.Printf("  response from Gemini: %s\n", result.Text())

	cleanString := strings.TrimPrefix(result.Text(), "```json\n")
	cleanString = strings.TrimSuffix(cleanString, "\n```")
	// cleanString = strings.TrimSpace(cleanString)

	fmt.Println(cleanString)

	var arr []SceneQuery
	erre := json.Unmarshal([]byte(cleanString), &arr)
	if erre != nil {
		panic(erre)
	}

	fmt.Println("this is yayaya", arr)

	fmt.Println("actual lenght brooooooo", len(arr))

	return arr, nil
}
