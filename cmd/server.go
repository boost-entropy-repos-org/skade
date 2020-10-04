package cmd

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/", s.indexPage()).Methods("GET")
}

func (s *Server) indexPage() http.HandlerFunc {
	fmt.Println("this happens")
	return func(w http.ResponseWriter, r *http.Request) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := IndexPage{
			Title: "Index",
		}
		t, err := template.ParseFiles(filepath.Join(wd + "/cmd/templates/upload.html"))
		if err != nil {
			panic(err)
		}
		t.Execute(w, p)
	}
}
