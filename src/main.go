package main

import (
	"go_visitors_maintain_backend/src/controller"
	"os"
)

func main() {
	os.Setenv("ACCESS_SECRET", "omkar")
	r := controller.RegisterRoutes()
	r.Run("127.0.0.1:8081")

}
