package main

import (
	"github.com/Mindslave/skade/internal/engine"
	"github.com/Mindslave/skade/internal/interactors/http"
	"github.com/Mindslave/skade/internal/log/zap"
	"github.com/Mindslave/skade/internal/repositories/postgres"
	"net/http"
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
	var repository engine.Repository
	repositoryType := "postgres"
	switch repositoryType {
	case "postgres":
		repository = postgres.NewPostgresRepository()
	default:
		panic("no storage")
	}

	// the interactor determines how our application will be used
	// for exmaple with an http interactor our application will be a web based one
	// with a cli interactor our application will be a cli based one
	var interactor engine.Interactor
	interactorType := "http"
	switch interactorType {
	case "http":
		interactor = http.NewHttpInteractor()
	default:
		panic("no interactor")
	}
	engine := engine.NewAnalysisService(logger, repository, interactor)
	engine.Scan("testfiles/garbage.exe")
	srv := server.NewServer()
	http.ListenAndServe(":8080", srv)
}
