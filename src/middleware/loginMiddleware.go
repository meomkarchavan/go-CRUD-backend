package middleware

import (
	"go_visitors_maintain_backend/src/models"
	"go_visitors_maintain_backend/src/services"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func LoginMiddleware(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, "/login") ||
		strings.HasPrefix(c.Request.URL.Path, "/public") {
		return
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, "Token Required")
		return
	}
	token_arr := strings.Split(token, "Bearer ")
	if len(token_arr) != 2 {
		c.JSON(http.StatusUnauthorized, "Token Required")
		return
	}
	token = token_arr[1]
	claims := jwt.MapClaims{}

	err := services.CheckToken(token, claims, c)
	if err {

		return
	}
	data := claims["user_id"]

	loginToken, ok := models.LoginTokens[data.(string)]

	if !ok {
		c.JSON(http.StatusUnauthorized, "Token Not found on server")
		return
	}
	if loginToken.Access_token != token {
		c.JSON(http.StatusUnauthorized, "Server Token not matched")
		return
	}
	c.Next()
}
