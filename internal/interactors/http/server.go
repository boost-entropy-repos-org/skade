package server

import (
	"github.com/Mindslave/skade/internal/engine"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

type httpInteractor struct {
	router *Server
}

func NewHttpInteractor(engine.AnalysisService) httpInteractor {
	interactor := NewServer()
	return httpInteractor{
		interactor,
	}
}

func (h *httpInteractor) StartServer() {
	log.Fatal(http.ListenAndServe(":8080", h.router))
}
