package handlers

import (
	"io"
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

	// Calling services //
	response, err := services.FullProcessService(data.Text)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
}
