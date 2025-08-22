package services

import (
	"Backend/internal/api"
	"fmt"
	"log"
)

func GetGeminiResponseService(user_prompt string) (string, error) {
	response, err := api.GetGeminiResponse(user_prompt)
	if err != nil {
		log.Fatal("Error getting response from Gemini: ", err)
	}

	fmt.Println("got /gemini request")
	return response, nil
}

func GetPexelsResponseService(gemini_prompt string) (string, error) {
	response, err := api.GetPexelsResponse(gemini_prompt)
	if err != nil {
		log.Fatal("Error getting response from Pexels: ", err)
	}

	fmt.Println("got /pexels request")
	return response, nil
}
