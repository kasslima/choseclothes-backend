package user

import (
	"choseclothes/internal/shared/errors"
	"choseclothes/internal/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	users, err := h.service.GetUsers(c.Request.Context())
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.OK(c, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errors.Respond(c, errors.NewValidationError(errors.ParseValidationErrors(validationErrors)))
			return
		}

		appErr := errors.ParseBindError(err)
		errors.Respond(c, appErr)
		return
	}

	createdUser, err := h.service.CreateUser(c.Request.Context(), input)
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
