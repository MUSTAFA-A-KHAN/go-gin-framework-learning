package main

import (
	"github.com/MUSTAFA-A-KHAN/go-gin-framework-learning/user-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handlers.Pong)
	r.GET("/users", handlers.GetAllUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.GET("/user/update/:id", handlers.UpdateUser)
	r.Run()
}
