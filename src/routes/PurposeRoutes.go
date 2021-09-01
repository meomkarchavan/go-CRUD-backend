package routes

import (
	"go_visitors_maintain_backend/src/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPurpose(c *gin.Context) {

	result, err := database.FindAllPurpose()
	if err != nil {
		c.JSON(http.StatusNotFound, "No Purpose Found")
		return
	}
	c.JSON(http.StatusOK, result)
}
