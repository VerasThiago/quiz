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

type SubmitPublishedQuizAPI interface {
	Handler(context *gin.Context)
}

type SubmitPublishedQuizHandler struct {
	builder.Builder
}

func (s *SubmitPublishedQuizHandler) InitFromBuilder(builder builder.Builder) *SubmitPublishedQuizHandler {
	s.Builder = builder
	return s
}

func (s *SubmitPublishedQuizHandler) Handler(context *gin.Context) {
	var err error
	var alreadySubmitted bool
	var quizBase *models.QuizBase
	var submission *models.Submission
	var request *published.SubmitRequest

	if err = context.ShouldBindJSON(&request); err != nil {
		error_handler.HandleBadRequestError(context, err)
		return
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if quizBase, err = s.GetRepository().GetQuizBaseByID(request.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, s.GetLog())
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), s.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, s.GetLog())
		return
	}

	if alreadySubmitted, err = s.GetRepository().CheckAlreadySubmittedQuizByUser(request.Token.ID, request.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, s.GetLog())
		return
	}

	if errors := request.ValidateSemantic(quizBase, alreadySubmitted); len(errors) > 0 {
		context.JSON(http.StatusUnauthorized, gin.H{"error": errors})
		context.Abort()
		return
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), s.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, s.GetLog())
		return
	}

	if submission, err = request.GetSubmissionBase(quizBase); err != nil {
		error_handler.HandleInternalServerError(context, err, s.GetLog())
		return
	}

	submission.UserID = request.Token.ID
	submission.QuizID = request.QuizID

	if err = s.GetRepository().CreateSubmission(submission); err != nil {
		error_handler.HandleInternalServerError(context, err, s.GetLog())
		return
	}

	if err = s.GetTask().CreateReportAync(submission.ID); err != nil {
		error_handler.HandleInternalServerError(context, err, s.GetLog())
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": "success"})
}
