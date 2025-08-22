package router

import (
	"Backend/internal/handlers"
	"net/http"
)

func APIRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.ApiHandler)
	return mux
}
