package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Mindslave/skade/internal/engine"
	"github.com/Mindslave/skade/internal/interactors/api"

	//"github.com/Mindslave/skade/internal/interactors/http"
	"github.com/Mindslave/skade/internal/log/zap"
	"github.com/Mindslave/skade/internal/repositories/postgresql"
	//"net/http"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://skadeuser:test@127.0.0.1:5432/skadedb?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
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

	conn, err := sqlx.Open(dbDriver, dbSource)
	if err != nil {
		logger.Error("Cannot connect to Database")
	}
	logger.Debug("successfully connected to Database")
	repoType := "postgres"
	switch repoType {
	case "postgres":
		repo = postgresql.NewRepo(conn)
	default:
		panic("no repository")
	}

	logger.Debug("starting server...")
	engine := engine.NewAnalysisService(logger, repo)
	//httpInteractor := server.NewHttpInteractor(engine, logger, repo)
	apiServer := api.NewAPIServer(engine, logger, repo)
	apiServer.Start(serverAddress)
	fmt.Println(apiServer)
}
