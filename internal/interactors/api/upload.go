package api

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (api *APIServer) upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("Failed to get the file: %w", err))
		return
	}
	filename := filepath.Base(file.Filename)
	err = ctx.SaveUploadedFile(file, filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("Failed to get the file: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("file has been uploaded succesfully"))
}