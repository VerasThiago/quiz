package quiz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	"github.com/verasthiago/quiz/quiz/pkg/validator/quiz"
	"github.com/verasthiago/quiz/shared/auth"
	error_handler "github.com/verasthiago/quiz/shared/errors"
)

type CreateQuizAPI interface {
	Handler(context *gin.Context)
}

type CreateQuizHandler struct {
	builder.Builder
}

func (c *CreateQuizHandler) InitFromBuilder(builder builder.Builder) *CreateQuizHandler {
	c.Builder = builder
	return c
}

func (c *CreateQuizHandler) Handler(context *gin.Context) {
	var err error
	var request *quiz.CreateRequest

	if err = context.ShouldBindJSON(&request); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), c.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, c.GetLog())
		return
	}

	request.UserID = request.Token.ID

	if err = c.GetRepository().CreateQuiz(request.Quiz); err != nil {
		error_handler.HandleInternalServerError(context, err, c.GetLog())
		return
	}

	context.JSON(http.StatusCreated, gin.H{"name": request.Name, "id": request.ID})
}
