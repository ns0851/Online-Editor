package main

import (
	"Backend/internal/router"
	"log"
	"net/http"
)

func main() {
	mux := router.SetupRouter() // get the router
	log.Println("Server running on http://localhost:9090")
	err := http.ListenAndServe(":9090", mux) // start HTTP server
	if err != nil {
		log.Fatal(err)
	}
}

//pexels - PiFlGwhKIh4z30H33wXUcXSUbYJKqml3Em2jrOYGLPF6LWts51VgzVVv
