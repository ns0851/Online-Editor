package router

import (
	"net/http"
	"Backend/internal/handlers"
)	

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.GetRoot)
	return mux
}
