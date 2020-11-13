package api

import (
	"github.com/Mindslave/skade/internal/engine"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	router *gin.Engine
	engine engine.AnalysisService
}

// NewAPIServer returns a new APIServer object
func NewAPIServer(engine engine.AnalysisService) APIServer {
	s := APIServer{
		engine: engine,
	}
	router := gin.Default()
	s.router = router
	return s
}