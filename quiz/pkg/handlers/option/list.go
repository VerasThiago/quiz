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

type ListOptionAPI interface {
	Handler(context *gin.Context)
}

type ListOptionHandler struct {
	builder.Builder
}

func (l *ListOptionHandler) InitFromBuilder(builder builder.Builder) *ListOptionHandler {
	l.Builder = builder
	return l
}

func (l *ListOptionHandler) Handler(context *gin.Context) {
	var err error
	var response []byte
	var quiz *models.Quiz
	var optionList []*models.Option
	request := option.ListRequest{
		QuestionID: context.Param("questionid"),
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), l.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if quiz, err = l.GetRepository().GetQuizByQuestionID(request.QuestionID); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if optionList, err = l.GetRepository().GetOptionListByQuestionID(request.QuestionID); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if response, err = models.OptionListByteArray(optionList, []string{"api", "user"}); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	context.Data(http.StatusOK, "application/json", response)
}
