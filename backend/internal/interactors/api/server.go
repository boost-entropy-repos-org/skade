package api

import (
	"github.com/Mindslave/skade/internal/engine"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	router 	*gin.Engine
	engine 	engine.AnalysisService
	logger 	engine.Logger
	repo	engine.Repository
}


// NewAPIServer returns a new APIServer object
func NewAPIServer(engine engine.AnalysisService, logger engine.Logger, repo engine.Repository) APIServer {
	s := APIServer{
		engine: engine,
		logger: logger,
		repo: repo,
	}
	router := gin.Default()

	router.GET("/token", s.getToken)
	router.POST("/token", s.getToken)
	router.Use(s.authMiddleware)
	router.POST("/createUser", s.createUser)
	router.GET("/users/:id", s.getUser)
	router.POST("/upload", s.upload)
	s.router = router
	return s
} 

//Start starts the API Server
func (api *APIServer) Start(address string) error {
	return api.router.Run()
}


func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}