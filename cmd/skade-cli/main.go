package main

import (
	"flag"
	"github.com/Mindslave/skade/internal/engine"
	"github.com/Mindslave/skade/internal/log/zap"
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
	fileName := flag.String("File", "", "Name of the suspicious file")
	flag.Parse()
	engine.Scan(*fileName)
}
