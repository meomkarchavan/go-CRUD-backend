package controller

import (
	"go_visitors_maintain_backend/src/routes"
	"go_visitors_maintain_backend/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	// r.Use(loginMiddleware)
	// r.Use(middleware.CORSMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome")
	})
	visit := r.Group("/visit")
	visit.GET("/", routes.GetAllVisit)

	visit.POST("/add", routes.AddVisit)
	visit.POST("/update", routes.UpdateVisit)
	visit.POST("/delete", routes.DeleteVisit)
	visit.GET("/:userId", routes.GetUserVisit)

	user := r.Group("/users")
	user.GET("/", routes.GetAllUsers)
	user.GET("/id/:userId", routes.GetUser)
	user.GET("/u/:username", routes.GetUser)
	user.POST("/add", routes.AddUser)
	user.POST("/update", routes.UpdateUser)
	user.POST("/delete", routes.DeleteUser)

	r.POST("/login", services.Login)
	r.POST("/signup", services.SignUp)

	return r
}
