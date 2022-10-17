package question

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type DeleteRequest struct {
	QuestionID string `json:"id"`
	Token      *auth.JWTClaim
}

func (d *DeleteRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"id": []string{"required", "uuid"},
	}

	options := govalidator.Options{
		Data:  d,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	return validator.MergeUrlValues(validator.GetRulesKey(rules), values)
}

func (d *DeleteRequest) ValidateSemantic(quiz *models.Quiz) []string {
	var errors []string

	if d.Token.IsAdmin {
		return errors
	}

	if quiz.UserID != d.Token.ID {
		errors = append(errors, "user can't execute this operation on questions from other users")
	}

	if quiz.IsPublished {
		errors = append(errors, "user can't update question from a published quiz")
	}

	return errors
}
