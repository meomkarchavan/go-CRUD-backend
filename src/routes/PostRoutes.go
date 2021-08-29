package routes

import (
	"blog_rest_api_gin/src/database"
	helper "blog_rest_api_gin/src/helpers"
	"blog_rest_api_gin/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func AddPost(c *gin.Context) {
	var post models.Post
	err := c.Bind(&post)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if (post == models.Post{}) {
		c.JSON(
			http.StatusBadRequest,
			"Please Provide Data",
		)
		return
	}
	post.PostId = helper.Uuid(1)
	validate := validator.New()
	err = validate.Struct(post)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	result, err := database.CreatePost(post)
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

func GetUserPost(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.FindUserPost(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, "No Post Found")
		return
	}
	c.JSON(http.StatusOK, result)
}
func GetAllPost(c *gin.Context) {

	result, err := database.FindAllPost()
	if err != nil {
		c.JSON(http.StatusNotFound, "No Post Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := c.Bind(&post); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.DeletePost(post)
	if err != nil {
		c.JSON(http.StatusNotFound, "No User Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	err := c.Bind(&post)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := database.UpdatePost(post)
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
