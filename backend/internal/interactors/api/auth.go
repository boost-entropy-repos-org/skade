package api

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


var superSecretTest = []byte("superSecret")

func (s *APIServer) authMiddleware(ctx *gin.Context) {
	s.logger.Debug("authorizing request")

	if ctx.Request.Header["Token"] == nil {
		s.logger.Debug("JWT missing")
		fmt.Fprintf(ctx.Writer, "Not Authorized")
		ctx.Abort()
	}

	token, err := jwt.Parse(ctx.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Failed to parse JWT")
		}
		return superSecretTest, nil
	})
	if err != nil {
		s.logger.Debug("JWT Parsing failed, not a valid JWT token")
		fmt.Fprintf(ctx.Writer, "Not Authorized %w", err)
		ctx.Abort()
		return
	}

	if !token.Valid {
		s.logger.Debug("JWT Validation failed")
		fmt.Fprintf(ctx.Writer, "Not Authorized")
		ctx.Abort()
		return
	}

	ctx.Next()
}

func (s* APIServer) getToken(ctx *gin.Context) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "test"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(superSecretTest)
	if err != nil {
		fmt.Errorf("Something went wrong: %w", err)
		return
	}

	//fmt.Fprintf(ctx.Writer, tokenString)
	ctx.JSON(http.StatusOK, tokenString)
}