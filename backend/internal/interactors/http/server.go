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
	return s
}

type httpInteractor struct {
	router *Server
	engine engine.AnalysisService
}

func NewHttpInteractor(engine engine.AnalysisService) httpInteractor {
	interactor := NewServer()
	h := httpInteractor{
		interactor,
		engine,
	}
	h.routes()
	return h
}

func (h *httpInteractor) StartServer() {
	log.Fatal(http.ListenAndServe(":8080", h.router))
}
