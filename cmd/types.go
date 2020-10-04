package cmd

import (
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

type IndexPage struct {
	Title string
}
