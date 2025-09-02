package api

import (
	"context"
	"log"
	"os"
	"github.com/kosa3/pexels-go"
)

func GetPexelsResponse(user_request []string) ([]string, error) {
	apiKey := os.Getenv("PEXELS_API_KEY") // safer than hardcoding
	if apiKey == "" {
		log.Fatal("Set PEXELS_API_KEY environment variable")
	}
	var arr []string

	client := pexels.NewClient(apiKey)

	// Example: search photos
	ctx := context.Background()

	for i := range user_request {
		
		params := &pexels.PhotoParams{
			Query:   user_request[i],
			Page:    1,
			PerPage: 1,
		}
	
		photos, err := client.PhotoService.Search(ctx, params)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, photos.Photos[0].Src.Original)
	}

	return arr, nil
}
