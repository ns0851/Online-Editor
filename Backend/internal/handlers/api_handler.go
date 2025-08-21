package handlers

import (
	"Backend/internal/api"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GeminiHandler(w http.ResponseWriter, r *http.Request) {

	response, err := api.GetGeminiResponse("who came first egg or chicken?")
	if err != nil {
		log.Fatal("Error getting response from Gemini: ", err)
	}
	io.WriteString(w, response)
	fmt.Println("got /gemini request")
}

func PexelsHandler(w http.ResponseWriter, r *http.Request) {

	response, err := api.GetPexelsResponse("nature")
	if err != nil {
		log.Fatal("Error getting response from Pexels: ", err)
	}
	// io.WriteString(w, `<h1>%s</h1><img src="%s" alt="photo">`, response)
	fmt.Fprintf(w, `<h1 >Image</h1><img style="width:35vw; height:50vh;" src="%s" alt="photo">`, response)
	// `<h1>%s</h1><img src="%s" alt="photo">`, response
	fmt.Println("got /pexels request")
}
