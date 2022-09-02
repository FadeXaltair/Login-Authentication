package main

import (
	"jwtAuthentication/controller"
	"jwtAuthentication/initiializers"
	"jwtAuthentication/middleware"
	"jwtAuthentication/models"

	"github.com/gin-gonic/gin"
)
func init() {
	initiializers.Loadenv()
	initiializers.ConnectDb()
}

func main() {
	initiializers.DB.AutoMigrate(&models.User{})
	r := gin.Default()

    r.POST( "/signup", controller.Signup)
	r.POST( "/login", controller.Login)
	r.GET(	"/validate", middleware.AuthRequire , controller.Validate)

	r.Run()
}
