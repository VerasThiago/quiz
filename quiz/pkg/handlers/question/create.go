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

type CreateQuestionAPI interface {
	Handler(context *gin.Context)
}

type CreateQuestionHandler struct {
	builder.Builder
}

func (c *CreateQuestionHandler) InitFromBuilder(builder builder.Builder) *CreateQuestionHandler {
	c.Builder = builder
	return c
}

func (c *CreateQuestionHandler) Handler(context *gin.Context) {
	var err error
	var quiz *models.Quiz
	var request question.CreateRequest

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

	if quiz, err = c.GetRepository().GetQuizByID(request.Question.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, c.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if err = c.GetRepository().CreateQuestion(request.Question); err != nil {
		error_handler.HandleInternalServerError(context, err, c.GetLog())
		return
	}

	optionIDList := []string{}

	if request.OptionList != nil {
		if err = c.GetRepository().CreateOptionList(request.OptionList, request.Question.ID); err != nil {
			error_handler.HandleInternalServerError(context, err, c.GetLog())
			return
		}

		for _, option := range *request.OptionList {
			optionIDList = append(optionIDList, option.ID)
		}
	}

	context.JSON(http.StatusCreated, gin.H{"id": request.Question.ID, "optionidlist": optionIDList})
}
