package main

import (
	"Backend/internal/router"
	"fmt"
	"net/http"
)

func main() {
	mainMux := http.NewServeMux()

	// Root (non-API) routes
	mainMux.Handle("/", router.SetupRouter())

	// API routes mounted under /api
	mainMux.Handle("/api/", http.StripPrefix("/api", router.APIRouter()))

	fmt.Println("Server started at :9191")
	if err := http.ListenAndServe(":9191", mainMux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
