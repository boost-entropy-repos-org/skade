package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Mindslave/skade/backend/internal/entities"
)

type createUserRequest struct {
	Username 		string `json:"username" binding:"required"`
	HashedPassword 	string `json:"hashed_password" binding:"required"`
	Email    		string `json:"email" binding:"required"`
}

func (s *APIServer) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := entities.DbCreateUserParams {
		Username: req.Username,
		HashedPassword: req.HashedPassword,
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


func (s *APIServer) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.repo.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}