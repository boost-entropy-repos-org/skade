package server

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

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

func (s *Server) customCss() http.HandlerFunc {
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
			filepath.Join(wd + "/web/static/css/custom.css"),
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
