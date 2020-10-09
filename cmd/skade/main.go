package main

import (
	"github.com/Mindslave/skade/internal/engine"
	//"github.com/Mindslave/skade/internal/http"
	//"net/http"
)

func main() {
	engine := engine.NewAnalysisService()
	engine.Scan("garbage.exe")
	//srv := server.NewServer()
	//http.ListenAndServe(":8080", srv)
}
