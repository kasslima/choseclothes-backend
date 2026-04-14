package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}