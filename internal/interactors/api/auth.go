package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *APIServer) authMiddleware(ctx *gin.Context) {
	s.logger.Debug("authorizing request")
	if ctx.Request.Header["Token"] == nil {
		s.logger.Debug("authorization failed")
		fmt.Fprintf(ctx.Writer, "Not Authorized")
		ctx.Abort()
	}
}