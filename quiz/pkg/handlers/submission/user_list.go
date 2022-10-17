package submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	"github.com/verasthiago/quiz/shared/auth"
	error_handler "github.com/verasthiago/quiz/shared/errors"
	"github.com/verasthiago/quiz/shared/models"
)

type ListUserSubmissionAPI interface {
	Handler(context *gin.Context)
}

type ListUserSubmissionHandler struct {
	builder.Builder
}

func (l *ListUserSubmissionHandler) InitFromBuilder(builder builder.Builder) *ListUserSubmissionHandler {
	l.Builder = builder
	return l
}

func (l *ListUserSubmissionHandler) Handler(context *gin.Context) {
	var err error
	var response []byte
	var token *auth.JWTClaim
	var submissionList []*models.Submission

	if token, err = auth.GetJWTClaimFromToken(context.GetHeader("Authorization"), l.GetSharedFlags().JwtKey); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if submissionList, err = l.GetRepository().GetSubmissionListByUserID(token.ID); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	if response, err = models.SubmissionListByteArray(submissionList, []string{"api"}); err != nil {
		error_handler.HandleInternalServerError(context, err, l.GetLog())
		return
	}

	context.Data(http.StatusOK, "application/json", response)
}
