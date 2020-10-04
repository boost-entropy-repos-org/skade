package main

import (
	"github.com/Mindslave/skade/cmd"
	"net/http"
)

var (
	wd  string
	err error
)

type IndexPage struct {
	Title string
}

func main() {
	srv := cmd.NewServer()
	http.ListenAndServe(":8080", srv)
}
