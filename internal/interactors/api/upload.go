package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *APIServer) upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("Failed to get the file: %w", err))
		return
	}
	


}