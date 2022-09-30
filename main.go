package main

import (
	"jwtAuthentication/controller"
	"jwtAuthentication/initiializers"
	"jwtAuthentication/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initiializers.Loadenv()
	initiializers.ConnectDb()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.GET("/validate", middleware.AuthRequire, controller.Validate)

	r.Run()
}
