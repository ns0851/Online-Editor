package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / requestttt\n")
	io.WriteString(w, "Main Page")
}
