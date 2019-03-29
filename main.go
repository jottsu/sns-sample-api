package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jottsu/sns-sample-api/controllers"
	"github.com/jottsu/sns-sample-api/repositories"
)

func init() {
	repositories.Init()
}

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Signin)

	r.GET("/users/:id", controllers.UserShow)
	r.POST("/users/:id/update", controllers.UserUpdate)

	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts/:id", controllers.PostShow)

	r.Run(":8080")
}
