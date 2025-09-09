package api

import (
	"context"
	"fmt"
	"os"

	"github.com/kosa3/pexels-go"
)

func GetPexelsResponse(user_request []SceneQuery) ([]SceneQuery, error) {
	apiKey := os.Getenv("PEXELS_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("PEXELS_API_KEY not set")
	}

	client := pexels.NewClient(apiKey)
	ctx := context.Background()

	for i, scene := range user_request {
		switch scene.Type {
		case "image":
			url, err := searchPhoto(ctx, client, scene.Query)
			if err != nil {
				return nil, err
			}
			user_request[i].URL = url
		case "video":
			url, err := searchVideo(ctx, client, scene.Query)
			if err != nil {
				return nil, err
			}
			user_request[i].URL = url
		}
		
	}

	return user_request, nil
}

func searchPhoto(ctx context.Context, client *pexels.Client, query string) (string, error) {
	params := &pexels.PhotoParams{Query: query, Page: 1, PerPage: 1}
	res, err := client.PhotoService.Search(ctx, params)
	if err != nil {
		return "", err
	}
	if len(res.Photos) == 0 {
		return "", fmt.Errorf("no photo found for %s", query)
	}
	return res.Photos[0].Src.Original, nil
}

func searchVideo(ctx context.Context, client *pexels.Client, query string) (string, error) {
	params := &pexels.VideoParams{Query: query, Page: 1, PerPage: 1}
	res, err := client.VideoService.Search(ctx, params)
	if err != nil {
		return "", err
	}
	if len(res.Videos) == 0 || len(res.Videos[0].VideoFiles) == 0 {
		return "", fmt.Errorf("no video found for %s", query)
	}
	return res.Videos[0].VideoFiles[0].Link, nil
}
