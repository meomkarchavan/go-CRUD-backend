package services

import (
	"blog_rest_api_gin/src/database"
	"blog_rest_api_gin/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	result, err := database.FindUser(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No user found")
		return
	}

	if result.Password != user.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(result)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	lt := models.LoginToken{
		Access_token:  token["access_token"],
		Refresh_token: token["refresh_token"],
		UserId:        result.UserId,
	}
	models.LoginTokens[result.UserId] = &lt
	c.Header("Authorization", token["access_token"])
	c.Header("refresh_token", token["refresh_token"])
	c.JSON(http.StatusOK, token)
}

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	_, err := database.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Err")
		return
	}
	token, err := CreateToken(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	lt := models.LoginToken{
		Access_token:  token["access_token"],
		Refresh_token: token["refresh_token"],
		UserId:        user.UserId,
	}
	models.LoginTokens[user.UserId] = &lt
	c.Header("Authorization", token["access_token"])
	c.Header("refresh_token", token["refresh_token"])
	c.JSON(http.StatusOK, token)
}
