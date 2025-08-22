package api

import (
	"context"
	"log"
	"os"

	"github.com/kosa3/pexels-go"
)

func GetPexelsResponse(user_request string) (string, error) {
	apiKey := os.Getenv("PEXELS_API_KEY") // safer than hardcoding
	if apiKey == "" {
		log.Fatal("Set PEXELS_API_KEY environment variable")
	}

	client := pexels.NewClient(apiKey)

	// Example: search photos
	ctx := context.Background()
	params := &pexels.PhotoParams{
		Query:   user_request,
		Page:    1,
		PerPage: 1,
	}

	photos, err := client.PhotoService.Search(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	return photos.Photos[0].Src.Original, nil
}
