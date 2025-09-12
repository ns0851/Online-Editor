package api

import (
	"context"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
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
			url, err := searchPhoto(ctx, apiKey, scene.Query)
			if err != nil {
				fmt.Printf("photo search failed for %q: %v\n", scene.Query, err)
				continue // skip this one, keep others
			}
			user_request[i].URL = url

		case "video":
			url, err := searchVideo(ctx, client, scene.Query)
			if err != nil {
				fmt.Printf("video search failed for %q: %v\n", scene.Query, err)
				continue
			}
			user_request[i].URL = url
		}
	}
	return user_request, nil

}

type PexelsPhotoSearchResponse struct {
    Photos []struct {
        Src struct {
            Original string `json:"original"`
            Medium   string `json:"medium"`
        } `json:"src"`
    } `json:"photos"`
}

func searchPhoto(ctx context.Context, apiKey string, query string) (string, error) {
    client := &http.Client{}

    baseURL := "https://api.pexels.com/v1/search"
    reqURL := fmt.Sprintf("%s?query=%s&per_page=1", baseURL, url.QueryEscape(query))

    req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
    if err != nil {
        return "", fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Authorization", apiKey)

    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("request execution failed: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("pexels API returned non-200 status: %s", resp.Status)
    }

    var pexelsResponse PexelsPhotoSearchResponse
    if err := json.NewDecoder(resp.Body).Decode(&pexelsResponse); err != nil {
        return "", fmt.Errorf("failed to decode json response: %w", err)
    }

    if len(pexelsResponse.Photos) == 0 {
        return "", fmt.Errorf("no photo found for query: '%s'", query)
    }

    if pexelsResponse.Photos[0].Src.Original == "" {
        return "", fmt.Errorf("photo found but original link is empty for query: '%s'", query)
    }

    return pexelsResponse.Photos[0].Src.Original, nil
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
