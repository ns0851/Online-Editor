package handlers

import (
	"io"
	"fmt"
	"net/http"
	"encoding/json"
	"Backend/internal/services"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	// Getting user request from BODY //
	body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading body", http.StatusBadRequest)
        return
    }
	defer r.Body.Close()

	// Creating a struct to make body.text work by storing it in Text of the struct //
	var data struct {
    Text string `json:"text"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Getting response from Gemini //
	respo, err := services.GetGeminiResponseService(data.Text)
	if err != nil {
		http.Error(w, "Error getting response from Gemini: "+err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, respo)

	// Getting response from Pexels by passing response from Gemini to it //
	res, err := services.GetPexelsResponseService(respo)
	if err != nil {
		http.Error(w, "Error getting response from Pexels: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// io.WriteString(w, res)
	// fmt.Fprintf(w, `<h1 >Image</h1><img style="width:35vw; height:70vh;" src="%s" alt="photo">`, response)
	
}
