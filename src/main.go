package main

import "blog_rest_api_gin/src/controller"

func main() {

	r := controller.RegisterRoutes()
	r.Run("127.0.0.1:8080")

}
