package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ooy30/tutorial-go/controller"
	"github.com/ooy30/tutorial-go/entity"
)

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	// User Routes
	r.GET("/users", controller.ListUsers)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
	// Run the server
	r.Run()
}
