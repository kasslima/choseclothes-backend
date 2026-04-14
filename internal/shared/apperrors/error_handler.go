package apperrors

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, err AppError) {
	c.JSON(err.Code, err)
}
