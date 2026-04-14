package user

import (
	"choseclothes/internal/shared/errors"
	"choseclothes/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service Service
}

func NewUserHandler(service Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.service.GetUsers()
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

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		if appErr, ok := err.(errors.AppError); ok {
			errors.Respond(c, appErr)
			return
		}

		errors.Respond(c, errors.NewInternal(err.Error()))
		return
	}

	response.Created(c, createdUser)
}
