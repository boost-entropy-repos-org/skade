package main

import (
	"github.com/Mindslave/skade/internal/database"
	"github.com/Mindslave/skade/internal/server"
	"net/http"
)

func main() {
	srv := server.NewServer()
	http.ListenAndServe(":8080", srv)
}
