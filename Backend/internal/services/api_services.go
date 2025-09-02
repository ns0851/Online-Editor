package services

import (
	"Backend/internal/api"
	"fmt"
	"log"
)

func FullProcessService(user_prompt string) ([]string, error) {
	res, err := api.GetGeminiResponse(user_prompt)
	if err != nil {
		log.Fatal("Error getting response from Gemini: ", err)
	}

	fmt.Println("in service", len(res))

	fmt.Println("got /gemini request")

	resp, err := api.GetPexelsResponse(res)
	if err != nil {
		return nil, err
	}

	fmt.Println("got the thingggg", resp)
	fmt.Println("got /pexels request")
	return resp, nil
}
