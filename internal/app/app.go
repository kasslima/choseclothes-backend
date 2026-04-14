package app

import (
	"choseclothes/ent"
	"choseclothes/internal/http"
	"choseclothes/internal/user"

	"github.com/gin-gonic/gin"
)

func BuildRouter(client *ent.Client) *gin.Engine {
	// user
	userRepo := user.NewRepository(client)
	userService := user.NewService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// router
	r := http.SetupRouter(userHandler)

	return r
}
