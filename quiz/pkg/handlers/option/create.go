package option

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	"github.com/verasthiago/quiz/quiz/pkg/validator/option"
	"github.com/verasthiago/quiz/shared/auth"
	error_handler "github.com/verasthiago/quiz/shared/errors"
	"github.com/verasthiago/quiz/shared/models"
)

type CreateOptionAPI interface {
	Handler(context *gin.Context)
}

type CreateOptionHandler struct {
	builder.Builder
}

func (c *CreateOptionHandler) InitFromBuilder(builder builder.Builder) *CreateOptionHandler {
	c.Builder = builder
	return c
}

func (c *CreateOptionHandler) Handler(context *gin.Context) {
	var err error
	var quiz *models.Quiz
	var request option.CreateRequest

	if err = context.ShouldBindJSON(&request); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), c.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, c.GetLog())
		return
	}

	if quiz, err = c.GetRepository().GetQuizByQuestionID(request.QuestionID); err != nil {
		error_handler.HandleInternalServerError(context, err, c.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if err = c.GetRepository().CreateOption(request.Option); err != nil {
		error_handler.HandleInternalServerError(context, err, c.GetLog())
		return
	}

	context.JSON(http.StatusCreated, gin.H{"id": request.ID})
}
