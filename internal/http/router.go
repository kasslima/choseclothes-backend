package http

import (
	"choseclothes/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *user.UserHandler) *gin.Engine {
	r := gin.Default()

	user.RegisterRoutes(r, userHandler)

	return r
}
