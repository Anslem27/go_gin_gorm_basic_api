package main

import (
	"go_crud/controllers"
	"go_crud/initializers"
	"go_crud/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectDb()
}
func main() {

	r := gin.Default()
	/* users */
	r.POST("/api/v1/users/signup", controllers.SignUpUser)
	r.POST("/api/v1/users/login", controllers.LogIn)
	r.GET("/api/v1/users", controllers.GetAllUsers)
	r.GET("/api/v1/users/validate", middleware.RequireAuth, controllers.Validate)
	//
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	r.Run()
}
