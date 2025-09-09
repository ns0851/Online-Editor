package services

import (
	"Backend/internal/api"
	"fmt"
	"log"
)

func FullProcessService(user_prompt string) ([]api.SceneQuery, error) {
	res, err := api.GetGeminiResponse(user_prompt)
	if err != nil {
		log.Printf("SERVICE ERROR from Gemini: %v", err)
		return nil, fmt.Errorf("failed to get a valid response from AI service: %w", err)
	}

	fmt.Println("in service", res)

	fmt.Println("got /gemini request")

	resp, err := api.GetPexelsResponse(res)
	if err != nil {
		return nil, err
	}

	fmt.Println("got the thingggg", resp)
	fmt.Println("got /pexels request")
	return res, nil
	// return []string{"haha"}, nil
}
