package server

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func (s *Server) loginPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := Page{
			Title: "Login",
		}
		t, err := template.ParseFiles(filepath.Join(wd + "/web/templates/login.html"))
		if err != nil {
			panic(err)
		}
		t.Execute(w, p)
	}
}
