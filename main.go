package main

import (
	"fmt"
	"github.com/bednyak/go-jwt-auth-cookie-v2/controllers"
	"github.com/bednyak/go-jwt-auth-cookie-v2/initializers"
	"github.com/bednyak/go-jwt-auth-cookie-v2/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabse()
}

func main() {
	fmt.Println("Test2")
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
