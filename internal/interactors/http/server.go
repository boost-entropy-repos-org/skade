package server

import (
	"github.com/Mindslave/skade/internal/engine"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func (s *Server) indexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := Page{
			Title: "Index",
		}
		t, err := template.ParseFiles(filepath.Join(wd + "/web/templates/upload.html"))
		if err != nil {
			panic(err)
		}
		t.Execute(w, p)
	}
}

func (s *Server) homePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := Page{
			Title: "Index",
		}

		files := []string{
			filepath.Join(wd + "/web/templates/test.html"),
			filepath.Join(wd + "/web/templates/base.html"),
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			panic(err)
		}
		t.Execute(w, p)
	}
}

func (s *Server) bulmaCss() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/css")
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := Page{
			Title: "Index",
		}
		files := []string{
			filepath.Join(wd + "/web/static/css/bulma.min.css"),
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			panic(err)
		}
		t.Execute(w, p)
	}
}

func (s *Server) dropzoneCss() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/css")
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := Page{
			Title: "Index",
		}
		files := []string{
			filepath.Join(wd + "/web/static/css/dropzone.min.css"),
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			panic(err)
		}
		t.Execute(w, p)
	}
}

func (s *Server) basicCss() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/css")
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := Page{
			Title: "Index",
		}
		files := []string{
			filepath.Join(wd + "/web/static/css/basic.min.css"),
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			panic(err)
		}
		t.Execute(w, p)
	}
}

func (s *Server) dropzoneJs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/javascript")
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := Page{
			Title: "Index",
		}
		files := []string{
			filepath.Join(wd + "/web/static/js/dropzone.min.js"),
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			panic(err)
		}
		t.Execute(w, p)
	}
}
