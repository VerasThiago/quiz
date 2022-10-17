package question

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type ListRequest struct {
	QuizID string `json:"id"`
	Token  *auth.JWTClaim
}

func (l *ListRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"id": []string{"required", "uuid"},
	}

	options := govalidator.Options{
		Data:  l,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	return validator.MergeUrlValues(validator.GetRulesKey(rules), values)
}

func (l *ListRequest) ValidateSemantic(quiz *models.Quiz) []string {
	var errors []string

	if l.Token.IsAdmin {
		return errors
	}

	if quiz.UserID != l.Token.ID {
		errors = append(errors, "user can't execute this operation on quizzes from other users")
	}

	return errors
}
