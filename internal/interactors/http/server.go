package server

import (
	"github.com/Mindslave/skade/internal/engine"
	"github.com/gorilla/mux"
	"html/template"
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
	router *mux.Router
}

func NewHttpInteractor() engine.Interactor {
	interactor := mux.NewRouter()
	return httpInteractor{
		interactor,
	}
}

func (h *httpInteractor) GetFile(FileName string) ([]byte, error) {
	return nil, nil
}

func (h *httpInteractor) AuthenticateUser(Name string) {

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

func (s *Server) css() http.HandlerFunc {
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
