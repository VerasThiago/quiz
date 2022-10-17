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

type ListQuestionAPI interface {
	Handler(context *gin.Context)
}

type ListQuestionHandler struct {
	builder.Builder
}

func (l *ListQuestionHandler) InitFromBuilder(builder builder.Builder) *ListQuestionHandler {
	l.Builder = builder
	return l
}

func (l *ListQuestionHandler) Handler(context *gin.Context) {
	var err error
	var response []byte
	var quiz *models.Quiz
	var questionList []*models.Question
	request := question.ListRequest{
		QuizID: context.Param("quizid"),
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), l.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if quiz, err = l.GetRepository().GetQuizByID(request.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if questionList, err = l.GetRepository().GetQuestionListByQuizID(request.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if response, err = models.QuestionListByteArray(questionList, []string{"api"}); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	context.Data(http.StatusOK, "application/json", response)
}
