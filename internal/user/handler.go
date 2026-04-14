package user

import (
	"choseclothes/internal/shared/apperrors"
	"choseclothes/internal/shared/response"
	"errors"

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
			apperrors.Respond(
				c,
				apperrors.NewValidationError(
					apperrors.ParseValidationErrors(validationErrors),
				),
			)
			return
		}

		appErr := apperrors.ParseBindError(err)
		apperrors.Respond(c, appErr)
		return
	}

	createdUser, err := h.service.CreateUser(c.Request.Context(), input)
	if err != nil {

		var appErr apperrors.AppError
		if errors.As(err, &appErr) {
			apperrors.Respond(c, appErr)
			return
		}

		apperrors.Respond(c, apperrors.NewInternal())
		return
	}

	response.Created(c, createdUser)
}
