package controller

import (
	"blog_rest_api_gin/src/routes"
	"blog_rest_api_gin/src/services"
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
	todo := r.Group("/todo")
	todo.GET("/", routes.GetAllTodo)

	todo.POST("/add", routes.AddTodo)
	todo.POST("/update", routes.UpdateTodo)
	todo.POST("/delete", routes.DeleteTodo)
	todo.GET("/:userId", routes.GetUserTodo)

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
