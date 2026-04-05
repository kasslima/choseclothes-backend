package user

import (
	"choseclothes/internal/shared/errors"
	"choseclothes/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

var users []User

func (h *UserHandler) GetUsers(c *gin.Context) {
	response.OK(c, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {

		if err.Error() == "EOF" {
			errors.Respond(c, errors.NewBadRequest("body is required"))
			return
		}

		validationErrors := errors.ParseValidationErrors(err)

		if len(validationErrors) > 0 {
			errors.Respond(c, errors.NewValidationError(validationErrors))
			return
		}

		errors.Respond(c, errors.NewBadRequest(err.Error()))
		return
	}

	users = append(users, user)

	response.Created(c, user)
}
