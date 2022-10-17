package quiz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	validator "github.com/verasthiago/quiz/quiz/pkg/validator/quiz"
	"github.com/verasthiago/quiz/shared/auth"
	error_handler "github.com/verasthiago/quiz/shared/errors"
	"github.com/verasthiago/quiz/shared/models"
)

type DeleteQuizAPI interface {
	Handler(context *gin.Context)
}

type DeleteQuizHandler struct {
	builder.Builder
}

func (d *DeleteQuizHandler) InitFromBuilder(builder builder.Builder) *DeleteQuizHandler {
	d.Builder = builder
	return d
}

func (d *DeleteQuizHandler) Handler(context *gin.Context) {
	var err error
	var quiz *models.Quiz
	var request *validator.DeleteRequest

	if err = context.ShouldBindJSON(&request); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), d.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, d.GetLog())
		return
	}

	if quiz, err = d.GetRepository().GetQuizByID(request.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, d.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if err = d.GetRepository().DeleteQuizByID(request.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, d.GetLog())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success"})
}
