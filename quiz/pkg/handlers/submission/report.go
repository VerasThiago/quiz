package submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	validator "github.com/verasthiago/quiz/quiz/pkg/validator/submission"
	"github.com/verasthiago/quiz/shared/auth"
	error_handler "github.com/verasthiago/quiz/shared/errors"
	"github.com/verasthiago/quiz/shared/models"
	"gorm.io/gorm"
)

type GetSubmissionReportAPI interface {
	Handler(context *gin.Context)
}

type GetSubmissionReportHandler struct {
	builder.Builder
}

func (g *GetSubmissionReportHandler) InitFromBuilder(builder builder.Builder) *GetSubmissionReportHandler {
	g.Builder = builder
	return g
}

func (g *GetSubmissionReportHandler) Handler(context *gin.Context) {
	var err error
	var response []byte
	var report *models.Report
	var quiz *models.Quiz
	var submission *models.Submission
	request := validator.GetReportRequest{
		SubmissionID: context.Param("submissionid"),
	}

	if request.Token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), g.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, g.GetLog())
		return
	}

	if errors := request.ValidateSyntax(); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if submission, err = g.GetRepository().GetSubmissionByID(request.SubmissionID); err != nil {
		error_handler.HandleInternalServerError(context, err, g.GetLog())
		return
	}

	if quiz, err = g.GetRepository().GetQuizByID(submission.QuizID); err != nil {
		error_handler.HandleInternalServerError(context, err, g.GetLog())
		return
	}

	if errors := request.ValidateSemantic(submission, quiz); len(errors) > 0 {
		error_handler.HandleBadRequestErrors(context, errors)
		return
	}

	if report, err = g.GetRepository().GetFullReportBySubmissionID(request.SubmissionID); err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusAccepted, gin.H{"status": "report is beign generated"})
			context.Abort()
			return
		} else {
			error_handler.HandleInternalServerError(context, err, g.GetLog())
			return
		}

	}

	if response, err = report.ByteArray([]string{"api"}); err != nil {
		error_handler.HandleInternalServerError(context, err, g.GetLog())
		return
	}

	context.Data(http.StatusOK, "application/json", response)
}
