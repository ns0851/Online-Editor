package router

import (
	"Backend/internal/handlers"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.GetRoot)
	mux.HandleFunc("/gemini", handlers.GeminiHandler)
	mux.HandleFunc("/pexels", handlers.PexelsHandler)
	return mux
}
