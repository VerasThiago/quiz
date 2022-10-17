package published

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	error_handler "github.com/verasthiago/quiz/shared/errors"
	"github.com/verasthiago/quiz/shared/models"
)

type ListPublishedQuizAPI interface {
	Handler(context *gin.Context)
}

type ListPublishedQuizHandler struct {
	builder.Builder
}

func (l *ListPublishedQuizHandler) InitFromBuilder(builder builder.Builder) *ListPublishedQuizHandler {
	l.Builder = builder
	return l
}

func (l *ListPublishedQuizHandler) Handler(context *gin.Context) {
	var err error
	var response []byte
	var quizList []*models.Quiz

	if quizList, err = l.GetRepository().GetPublishedQuizList(); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if response, err = models.QuizListByteArray(quizList, []string{"api"}); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	context.Data(http.StatusOK, "application/json", response)
}
