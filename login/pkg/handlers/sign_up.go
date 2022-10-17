package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/login/pkg/builder"
	"github.com/verasthiago/quiz/login/pkg/validator"
	error_handler "github.com/verasthiago/quiz/shared/errors"
)

type CreateUserAPI interface {
	Handler(context *gin.Context)
}

type CreateUserHandler struct {
	builder.Builder
}

func (c *CreateUserHandler) InitFromBuilder(builder builder.Builder) *CreateUserHandler {
	c.Builder = builder
	return c
}

func (c *CreateUserHandler) Handler(context *gin.Context) {
	var request validator.SignUpRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	if errors := request.Validate(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if err := request.HashPassword(request.Password); err != nil {
		error_handler.HandleInternalServerError(context, err, c.GetLog())
		return
	}

	if err := c.GetRepository().CreateUser(request.User); err != nil {
		error_handler.HandleInternalServerError(context, err, c.GetLog())
		return
	}

	context.JSON(http.StatusCreated, gin.H{"email": request.Email, "username": request.Username})
}
