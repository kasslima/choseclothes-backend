package main

import (
	"choseclothes/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userHandler := user.UserHandler{}

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)

	r.Run()
}
