package controller

import (
	"blog_rest_api_gin/src/middleware"
	"blog_rest_api_gin/src/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome")
	})
	post := r.Group("/post")
	post.GET("/", routes.GetAllPost)

	post.POST("/add", routes.AddPost)
	post.POST("/update", routes.UpdatePost)
	post.POST("/delete", routes.DeletePost)
	post.GET("/:userId", routes.GetUserPost)

	user := r.Group("/users")
	user.GET("/", routes.GetAllUsers)
	user.GET("/id/:userId", routes.GetUser)
	// user.GET("/u/:username", routes.GetUser)
	user.POST("/add", routes.AddUser)
	user.POST("/update", routes.UpdateUser)
	user.POST("/delete", routes.DeleteUser)

	return r
}
