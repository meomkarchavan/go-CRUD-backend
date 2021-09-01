package controller

import (
	"go_visitors_maintain_backend/src/middleware"
	"go_visitors_maintain_backend/src/routes"
	"go_visitors_maintain_backend/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	// r.Use(middleware.LoginMiddleware)
	// r.Use(middleware.CORSMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome")
	})
	pass := r.Group("/pass")
	pass.GET("/", middleware.LoginMiddleware, routes.GetAllPass)

	pass.POST("/add", middleware.LoginMiddleware, routes.AddPass)
	pass.POST("/update", middleware.LoginMiddleware, middleware.CheckRoleMiddleware, routes.UpdatePass)
	pass.POST("/delete", middleware.LoginMiddleware, routes.DeletePass)
	pass.GET("/:userId", middleware.LoginMiddleware, routes.GetUserPass)

	user := r.Group("/users")
	user.GET("/", middleware.LoginMiddleware, middleware.CheckRoleMiddleware, routes.GetAllUsers)
	user.GET("/id/:userId", middleware.LoginMiddleware, middleware.CheckRoleMiddleware, routes.GetUser)
	user.GET("/u/:username", middleware.LoginMiddleware, middleware.CheckRoleMiddleware, routes.GetUser)
	user.POST("/add", middleware.LoginMiddleware, middleware.CheckRoleMiddleware, routes.AddUser)
	user.POST("/update", middleware.LoginMiddleware, middleware.CheckRoleMiddleware, routes.UpdateUser)
	user.POST("/delete", middleware.LoginMiddleware, middleware.CheckRoleMiddleware, routes.DeleteUser)

	r.GET("/purpose", routes.GetPurpose)

	r.POST("/login", services.Login)
	r.POST("/signup", services.SignUp)

	return r
}
