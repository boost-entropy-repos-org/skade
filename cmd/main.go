package main

import (
	"github.com/Mindslave/skade/internal"
	"net/http"
)

func main() {
	srv := cmd.NewServer()
	http.ListenAndServe(":8080", srv)
}
