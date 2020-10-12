package server

import (
	"fmt"
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
		if r.Method == "GET" {
			t, err := template.ParseFiles(filepath.Join(wd + "/web/templates/login.html"))
			if err != nil {
				panic(err)
			}
			t.Execute(w, p)
		} else {
			r.ParseForm()
			// here we validate the login
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
		}
	}
}
