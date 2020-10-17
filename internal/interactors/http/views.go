package server

import (
	"fmt"
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

//here we get the uploaded 'suspicious' file
func (s *Server) uploadEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(0 << 200)
		file, handler, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Printf("Upload File: %+v\n", handler.Filename)
		fmt.Printf("File size: %+v\n", handler.Size)
		fmt.Printf("File Header: %+v\n", handler.Header)
	}
}
