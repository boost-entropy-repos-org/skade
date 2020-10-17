package server

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

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
