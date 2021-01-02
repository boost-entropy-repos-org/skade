package api

import (
	"github.com/Mindslave/skade/backend/internal/engine"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type APIServer struct {
	router 	*gin.Engine
	engine 	engine.AnalysisService
	logger 	engine.Logger
	repo	engine.Repository
}


// NewAPIServer returns a new APIServer object
func NewAPIServer(engine engine.AnalysisService, logger engine.Logger, repo engine.Repository) APIServer {
	router := gin.Default()

	api := APIServer{
		router: router,
		engine: engine,
		logger: logger,
		repo: repo,
	}


	api.router.Use(cors.Default())

	v1 := api.router.Group("/api/v1") 
	{
		v1.GET("/auth", api.authenticate)
		v1.POST("/auth", api.authenticate)
		v1.POST("/createUser", api.createUser)
		v1.Use(api.authMiddleware)
		v1.GET("/users/:id", api.getUser)
		v1.POST("/upload", api.upload)
	}

	return api
} 

//Start starts the API Server
func (api *APIServer) Start(address string) error {
	api.router.Use(cors.Default())
	return api.router.Run()
}


func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}