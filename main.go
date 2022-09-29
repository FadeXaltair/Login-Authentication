package main

import (
	"jwtAuthentication/controller"
	"jwtAuthentication/initiializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initiializers.CheckJWTToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJUb2tlbiI6ImNyZWF0aW5nIHRva2VuIiwiZXhwIjoyNTI3MDAyODgzfQ.xxNiNwRAhlf-v5hfDd0rvA8RsCBJ02DdRbd6q7iVpxI")
	initiializers.ConnectDb()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	// r.GET("/validate", middleware.AuthRequire, controller.Validate)

	r.Run()
}
