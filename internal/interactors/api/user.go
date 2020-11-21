package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
	models "github.com/Mindslave/skade/internal/entities"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type createUserParams struct {
	Username string
	Email	 string

}

func (s *APIServer) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := models.DbCreateUserParams {
		Username: req.Username,
		Email: req.Email,
	}

	user, err := s.repo.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
	
}


type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}


func (s *APIServer) getUser(ctx gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
}