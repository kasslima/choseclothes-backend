package app

import (
	"choseclothes/internal/http"
	"choseclothes/internal/user"

	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	// user
	userRepo := user.NewRepository()
	userService := user.NewService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// router
	r := http.SetupRouter(userHandler)

	return r
}