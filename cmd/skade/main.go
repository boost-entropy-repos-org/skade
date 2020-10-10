package main

import (
	"github.com/Mindslave/skade/internal/engine"
	"github.com/Mindslave/skade/internal/log/zap"
	//"github.com/Mindslave/skade/internal/http"
	//"net/http"
)

func main() {
	var logger engine.Logger
	//there is only zap at the moment, but there might be more in the future
	loggerType := "zap"
	switch loggerType {
	case "zap":
		//engine.logger is now a zaplogger
		logger = zaplogger.NewEngineZapLogger()
	default:
		panic("no logger")
	}
	engine := engine.NewAnalysisService(logger)
	engine.Scan("garbage.exe")
	//srv := server.NewServer()
	//http.ListenAndServe(":8080", srv)
}
