package routes

import (
	"go_visitors_maintain_backend/src/database"
	helper "go_visitors_maintain_backend/src/helpers"
	"go_visitors_maintain_backend/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddVisit(c *gin.Context) {
	var visit models.Visit
	err := c.Bind(&visit)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if (visit == models.Visit{}) {
		c.JSON(
			http.StatusBadRequest,
			"Please Provide Data",
		)
		return
	}
	visit.VisitId = helper.Uuid(1)
	visit.Date = primitive.Timestamp{T: uint32(time.Now().Unix())}
	visit.Approved = false
	visit.Rejected = false
	validate := validator.New()
	err = validate.Struct(visit)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	result, err := database.CreateVisit(visit)
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

func GetUserVisit(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.FindUserVisit(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, "No Visit Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetAllVisit(c *gin.Context) {

	result, err := database.FindAllVisit()
	if err != nil {
		c.JSON(http.StatusNotFound, "No Visit Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteVisit(c *gin.Context) {
	var visit models.Visit
	if err := c.Bind(&visit); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Data Provided")
		return
	}
	result, err := database.DeleteVisit(visit)
	if err != nil {
		c.JSON(http.StatusNotFound, "No User Found")
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateVisit(c *gin.Context) {
	var visit models.Visit
	err := c.Bind(&visit)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := database.UpdateVisit(visit)
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
