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

type DeleteQuestionAPI interface {
	Handler(context *gin.Context)
}

type DeleteQuestionHandler struct {
	builder.Builder
}

func (d *DeleteQuestionHandler) InitFromBuilder(builder builder.Builder) *DeleteQuestionHandler {
	d.Builder = builder
	return d
}

type DeleteQuestionRequest struct {
	QuestionID string `json:"questionid"`
}

func (d *DeleteQuestionHandler) Handler(context *gin.Context) {
	var err error
	var quiz *models.Quiz
	var request question.DeleteRequest

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

	if quiz, err = d.GetRepository().GetQuizByQuestionID(request.QuestionID); err != nil {
		error_handler.HandleInternalServerError(context, err, d.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if err = d.GetRepository().DeleteQuestionByID(request.QuestionID); err != nil {
		error_handler.HandleInternalServerError(context, err, d.GetLog())
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success"})
}
