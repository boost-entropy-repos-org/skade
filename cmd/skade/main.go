package main

import (
	"github.com/Mindslave/skade/internal/engine"
	"github.com/Mindslave/skade/internal/interactors/http"
	"github.com/Mindslave/skade/internal/log/zap"
	"github.com/Mindslave/skade/internal/repositories/postgres"
	//"net/http"
)

func main() {
	var logger engine.Logger 
	var repo engine.Repository 
	loggerType := "zap"
	switch loggerType {
	case "zap":
		//engine.logger is now a zaplogger
		logger = zaplogger.NewEngineZapLogger()
	default:
		panic("no logger")
	}

	repoType := "postgres"
	switch repoType {
	case "postgres":
		repo = db.Repo
	default:
		panic("no repository")
	}

	logger.Debug("starting server...")
	engine := engine.NewAnalysisService(logger)
	httpInteractor := server.NewHttpInteractor(engine)
	httpInteractor.StartServer()
	//engine.Scan("testfiles/garbage.exe")
}
