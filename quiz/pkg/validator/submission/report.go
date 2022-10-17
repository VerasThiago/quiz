package submission

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type GetReportRequest struct {
	SubmissionID string `json:"id"`
	Token        *auth.JWTClaim
}

func (g *GetReportRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"id": []string{"required", "uuid"},
	}

	options := govalidator.Options{
		Data:  g,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	return validator.MergeUrlValues(validator.GetRulesKey(rules), values)
}

func (l *GetReportRequest) ValidateSemantic(submission *models.Submission, quiz *models.Quiz) []string {
	var errors []string

	if l.Token.IsAdmin {
		return errors
	}

	if quiz.UserID == l.Token.ID {
		return errors
	}

	if submission.UserID != l.Token.ID {
		errors = append(errors, "user can't execute this operation on submissions from other users")
	}

	return errors
}
