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
	api := APIServer{
		engine: engine,
		logger: logger,
		repo: repo,
	}
	router := gin.Default()
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"*"}

	router.Use(cors.New(config))

	v1 := router.Group("/api/v1") 
	{
		v1.GET("/auth", api.authenticate)
		v1.POST("/auth", api.authenticate)
		v1.POST("/createUser", api.createUser)
		v1.Use(api.authMiddleware)
		v1.GET("/users/:id", api.getUser)
		v1.POST("/upload", api.upload)
	}
	api.router = router
	return api
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