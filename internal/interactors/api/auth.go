package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *APIServer) authMiddleware(ctx *gin.Context) {
	s.logger.Debug("auth")
	fmt.Fprintf(ctx.Writer, "Not Authorized")
	ctx.Abort()
}