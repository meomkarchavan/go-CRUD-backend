package routes

import (
	"go_visitors_maintain_backend/src/database"
	helper "go_visitors_maintain_backend/src/helpers"
	"go_visitors_maintain_backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func AddPass(c *gin.Context) {
	var pass models.Pass
	err := c.Bind(&pass)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if (pass == models.Pass{}) {
		c.JSON(
			http.StatusBadRequest,
			"Please Provide Data",
		)
		return
	}
	pass.PassId = helper.Uuid(1)
	pass.Approved = false
	pass.Rejected = false
	validate := validator.New()
	err = validate.Struct(pass)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	result, err := database.CreatePass(pass)
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

func GetUserPass(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.FindUserPass(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, "No Pass Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetAllPass(c *gin.Context) {

	result, err := database.FindAllPass()
	if err != nil {
		c.JSON(http.StatusNotFound, "No Pass Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeletePass(c *gin.Context) {
	var pass models.Pass
	if err := c.Bind(&pass); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.DeletePass(pass)
	if err != nil {
		c.JSON(http.StatusNotFound, "No User Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdatePass(c *gin.Context) {
	var pass models.Pass
	err := c.Bind(&pass)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := database.UpdatePass(pass)
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
