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

type UpdateOptionAPI interface {
	Handler(context *gin.Context)
}

type UpdateOptionHandler struct {
	builder.Builder
}

func (u *UpdateOptionHandler) InitFromBuilder(builder builder.Builder) *UpdateOptionHandler {
	u.Builder = builder
	return u
}

func (u *UpdateOptionHandler) Handler(context *gin.Context) {
	var err error
	var quiz *models.Quiz
	var request option.UpdateRequest

	if err = context.ShouldBindJSON(&request); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), u.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, u.GetLog())
		return
	}

	if quiz, err = u.GetRepository().GetQuizByOptionID(request.Option.ID); err != nil {
		error_handler.HandleInternalServerError(context, err, u.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if err = u.GetRepository().UpdateOption(request.Option); err != nil {
		error_handler.HandleInternalServerError(context, err, u.GetLog())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success"})
}
