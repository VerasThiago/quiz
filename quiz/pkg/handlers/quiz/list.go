package quiz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	"github.com/verasthiago/quiz/shared/auth"
	error_handler "github.com/verasthiago/quiz/shared/errors"
	"github.com/verasthiago/quiz/shared/models"
)

type ListQuizAPI interface {
	Handler(context *gin.Context)
}

type ListQuizHandler struct {
	builder.Builder
}

func (l *ListQuizHandler) InitFromBuilder(builder builder.Builder) *ListQuizHandler {
	l.Builder = builder
	return l
}

func (l *ListQuizHandler) Handler(context *gin.Context) {
	var err error
	var response []byte
	var token *auth.JWTClaim
	var quizList []*models.Quiz

	if token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), l.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if quizList, err = l.GetRepository().GetQuizListByUserID(token.ID); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if response, err = models.QuizListByteArray(quizList, []string{"api"}); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	context.Data(http.StatusOK, "application/json", response)
}
