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
	s := APIServer{
		engine: engine,
		logger: logger,
		repo: repo,
	}
	router := gin.Default()
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))

	v1 := router.Group("/api/v1/") 
	{
		v1.GET("/auth", s.authenticate)
		v1.POST("/auth", s.authenticate)
		v1.Use(s.authMiddleware)
		v1.POST("/createUser", s.createUser)
		v1.GET("/users/:id", s.getUser)
		v1.POST("/upload", s.upload)
	}
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