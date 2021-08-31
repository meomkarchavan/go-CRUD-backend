package middleware

import (
	"go_visitors_maintain_backend/src/models"
	"go_visitors_maintain_backend/src/services"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckRoleMiddleware(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, "Token Required")
		c.Abort()
		return
	}
	token_arr := strings.Split(token, "Bearer ")
	if len(token_arr) != 2 {
		c.JSON(http.StatusUnauthorized, "Token Required")
		c.Abort()
		return
	}
	token = token_arr[1]
	claims := jwt.MapClaims{}

	err := services.CheckToken(token, claims, c)
	if err {
		c.Abort()
		return
	}
	data := claims["user_id"]

	loginToken, ok := models.LoginTokens[data.(string)]

	if !ok {
		c.JSON(http.StatusUnauthorized, "Token Not found on server")
		c.Abort()
		return
	}
	if loginToken.Access_token != token {
		c.JSON(http.StatusUnauthorized, "Server Token not matched")
		c.Abort()
		return
	}
	role := claims["role"]

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, "Admin Access Required")
		c.Abort()
		return

	}
	c.Next()
}
