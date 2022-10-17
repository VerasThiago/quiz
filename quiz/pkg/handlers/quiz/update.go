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

type UpdateQuizAPI interface {
	Handler(context *gin.Context)
}

type UpdateQuizHandler struct {
	builder.Builder
}

func (u *UpdateQuizHandler) InitFromBuilder(builder builder.Builder) *UpdateQuizHandler {
	u.Builder = builder
	return u
}

func (u *UpdateQuizHandler) Handler(context *gin.Context) {
	var err error
	var quiz *models.Quiz
	var request *validator.UpdateRequest

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

	if quiz, err = u.GetRepository().GetQuizByID(request.Quiz.ID); err != nil {
		error_handler.HandleInternalServerError(context, err, u.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if err = u.GetRepository().UpdateQuiz(request.Quiz); err != nil {
		error_handler.HandleInternalServerError(context, err, u.GetLog())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success"})
}
