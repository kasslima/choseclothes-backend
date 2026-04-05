package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *UserHandler) {
	users := r.Group("/users")

	users.GET("", handler.GetUsers)
	users.POST("", handler.CreateUser)
}
