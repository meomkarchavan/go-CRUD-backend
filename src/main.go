package main

import (
	"blog_rest_api_gin/src/controller"
	"os"
)

func main() {
	os.Setenv("ACCESS_SECRET", "omkar")
	r := controller.RegisterRoutes()
	r.Run("127.0.0.1:8080")

}
