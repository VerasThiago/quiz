package question

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	"github.com/verasthiago/quiz/quiz/pkg/validator/question"
	"github.com/verasthiago/quiz/shared/auth"
	error_handler "github.com/verasthiago/quiz/shared/errors"
	"github.com/verasthiago/quiz/shared/models"
)

type UpdateQuestionAPI interface {
	Handler(context *gin.Context)
}

type UpdateQuestionHandler struct {
	builder.Builder
}

func (u *UpdateQuestionHandler) InitFromBuilder(builder builder.Builder) *UpdateQuestionHandler {
	u.Builder = builder
	return u
}

func (u *UpdateQuestionHandler) Handler(context *gin.Context) {
	var err error
	var quiz *models.Quiz
	var request *question.UpdateRequest

	if err = context.ShouldBindJSON(request); err != nil {
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

	if quiz, err = u.GetRepository().GetQuizByQuestionID(request.Question.ID); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if err = u.GetRepository().UpdateQuestion(request.Question); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success"})
}
