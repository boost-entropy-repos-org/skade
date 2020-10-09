package main

import (
	"github.com/Mindslave/skade/internal/engine"
	//"github.com/Mindslave/skade/internal/http"
	//"net/http"
)

func main() {
	engine := engine.NewAnalysisService()
	engine.Scan("testfile.txt")
	//srv := server.NewServer()
	//http.ListenAndServe(":8080", srv)
}
