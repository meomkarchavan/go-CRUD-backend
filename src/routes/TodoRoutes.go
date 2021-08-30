package routes

import (
	"blog_rest_api_gin/src/database"
	helper "blog_rest_api_gin/src/helpers"
	"blog_rest_api_gin/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func AddTodo(c *gin.Context) {
	var todo models.Todo
	err := c.Bind(&todo)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if (todo == models.Todo{}) {
		c.JSON(
			http.StatusBadRequest,
			"Please Provide Data",
		)
		return
	}
	todo.TodoId = helper.Uuid(1)
	todo.Done = false
	validate := validator.New()
	err = validate.Struct(todo)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	result, err := database.CreateTodo(todo)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}
	c.JSON(
		http.StatusOK,
		result,
	)
}

func GetUserTodo(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.FindUserTodo(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, "No Todo Found")
		return
	}
	c.JSON(http.StatusOK, result)
}
func GetAllTodo(c *gin.Context) {

	result, err := database.FindAllTodo()
	if err != nil {
		c.JSON(http.StatusNotFound, "No Todo Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.DeleteTodo(todo)
	if err != nil {
		c.JSON(http.StatusNotFound, "No User Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	err := c.Bind(&todo)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := database.UpdateTodo(todo)
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
