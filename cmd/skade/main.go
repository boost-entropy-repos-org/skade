package main

import (
	"github.com/Mindslave/skade/internal/http"
	"net/http"
)

func main() {
	srv := server.NewServer()
	http.ListenAndServe(":8080", srv)
}
