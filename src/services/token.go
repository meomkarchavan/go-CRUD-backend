package services

import (
	"go_visitors_maintain_backend/src/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(user models.User) (map[string]string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.UserId
	atClaims["username"] = user.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	// refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  token,
		"refresh_token": rt,
	}, nil

}

func CheckToken(token string, claims jwt.MapClaims, c *gin.Context) bool {
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			c.JSON(http.StatusUnauthorized, "ErrSignatureInvalid")
			return true
		}
		c.JSON(http.StatusUnauthorized, "Token Expired")
		return true
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, "Token Invalid")
		return true
	}
	return false
}
