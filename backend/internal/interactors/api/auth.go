package api

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


var superSecretTest = []byte("superSecret")

func (api *APIServer) authMiddleware(ctx *gin.Context) {
	api.logger.Debug("authorizing request")

	if ctx.Request.Header["Token"] == nil {
		api.logger.Debug("JWT missing")
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
		api.logger.Debug("JWT Parsing failed, not a valid JWT token")
		fmt.Fprintf(ctx.Writer, "Not Authorized %w", err)
		ctx.Abort()
		return
	}

	if !token.Valid {
		api.logger.Debug("JWT Validation failed")
		fmt.Fprintf(ctx.Writer, "Not Authorized")
		ctx.Abort()
		return
	}

	ctx.Next()
}


var authData struct {
	Username 	string
	Password 	string
}


func (api *APIServer) authenticate(ctx *gin.Context) {
	api.logger.Debug("received authentication request")

	//expect a json request body
	if ctx.Request.Header.Get("Content-Type") != "application/json" {
		err := fmt.Errorf("bad Content-Type")
		api.logger.Error("received auth request with invalid Content-Type header")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := ctx.ShouldBindJSON(&authData)

	if err != nil {
		err := fmt.Errorf("empty request body")
		api.logger.Debug("received auth request with empty body")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(authData.Password), bcrypt.DefaultCost)
	if err != nil {
		api.logger.Error("Failed to hash the Password")
		//what could we possibly do?
		panic(1)
	} else {
		ctx.JSON(http.StatusOK, hash)
		return
	}

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

	fmt.Fprintf(ctx.Writer, tokenString)
	ctx.JSON(http.StatusOK, tokenString)
}