package server

import (
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

type Page struct {
	Title string
}
