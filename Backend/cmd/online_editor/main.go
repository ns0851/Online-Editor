package main

import (
	"fmt"
	"log"
	"net/http"
	"Backend/internal/router"
)

func main() {
	fmt.Println("Server Running on :9090")
	r := router.SetupRouter()
	log.Fatal(http.ListenAndServe(":9090", r))
}

