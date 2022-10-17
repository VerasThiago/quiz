package published

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	"github.com/verasthiago/quiz/quiz/pkg/validator/published"
	"github.com/verasthiago/quiz/shared/auth"
	error_handler "github.com/verasthiago/quiz/shared/errors"
	"github.com/verasthiago/quiz/shared/models"
)

type OpenPublishedQuizAPI interface {
	Handler(context *gin.Context)
}

type OpenPublishedQuizHandler struct {
	builder.Builder
}

func (o *OpenPublishedQuizHandler) InitFromBuilder(builder builder.Builder) *OpenPublishedQuizHandler {
	o.Builder = builder
	return o
}

func (o *OpenPublishedQuizHandler) Handler(context *gin.Context) {
	var err error
	var response []byte
	var quiz *models.Quiz
	var submissionExist bool
	request := published.OpenRequest{
		QuizID: context.Param("quizid"),
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), o.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, o.GetLog())
		return
	}

	if submissionExist, err = o.GetRepository().CheckAlreadySubmittedQuizByUser(request.Token.ID, request.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, o.GetLog())
		return
	}

	if errors := request.ValidateSemantic(submissionExist); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if quiz, err = o.GetRepository().GetFullQuizByID(request.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, o.GetLog())
		return
	}

	if response, err = quiz.ByteArray([]string{"api"}); err != nil {
		error_handler.HandleInternalServerError(context, err, o.GetLog())
		return
	}

	context.Data(http.StatusOK, "application/json", response)
}
