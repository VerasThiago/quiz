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

type PublishQuizAPI interface {
	Handler(context *gin.Context)
}

type PublishQuizHandler struct {
	builder.Builder
}

func (p *PublishQuizHandler) InitFromBuilder(builder builder.Builder) *PublishQuizHandler {
	p.Builder = builder
	return p
}

func (p *PublishQuizHandler) Handler(context *gin.Context) {
	var err error
	var quiz *models.Quiz
	var request validator.PublishRequest

	if err = context.ShouldBindJSON(&request); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), p.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, p.GetLog())
		return
	}

	if quiz, err = p.GetRepository().GetFullQuizByID(request.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, p.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if err = p.GetRepository().PublishQuiz(quiz); err != nil {
		error_handler.HandleInternalServerError(context, err, p.GetLog())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success"})
}
