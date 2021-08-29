package routes

import (
	"blog_rest_api_gin/src/database"
	helper "blog_rest_api_gin/src/helpers"
	"blog_rest_api_gin/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func AddUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user.UserId = helper.Uuid(1)
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	result, err := database.CreateUser(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	c.JSON(
		http.StatusOK,
		result,
	)
}
func GetAllUsers(c *gin.Context) {

	result, err := database.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, "No Post Found")
		return
	}
	c.JSON(http.StatusOK, result)
}
func GetUserId(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	result, err := database.FindUserId(user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, "No User Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}

	result, err := database.FindUserFromID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, "No User Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.DeleteUser(user.UserId)
	if err != nil {
		c.JSON(http.StatusNotFound, "No User Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	result, err := database.UpdateUser(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	c.JSON(
		http.StatusOK,
		result,
	)
}
