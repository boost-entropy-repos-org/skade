package main

import (
	"github.com/Mindslave/skade/internal/engine"
	"github.com/Mindslave/skade/internal/interactors/http"
	"github.com/Mindslave/skade/internal/log/zap"
	//"github.com/Mindslave/skade/internal/repositories/postgres"
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

	// the repositoryType determines how we store data
	// if we use a PostgresRepository, then our app will use a Postgres Database
	// another option could be json or redis
	//var repository engine.Repository
	//repositoryType := "postgres"
	//switch repositoryType {
	//case "postgres":
	//	repository = postgres.NewPostgresRepository()
	//default:
	//	panic("no storage")
	//}
	logger.Debug("starting server...")
	engine := engine.NewAnalysisService(logger)
	httpInteractor := server.NewHttpInteractor(engine)
	httpInteractor.StartServer()
	//engine.Scan("testfiles/garbage.exe")
}
