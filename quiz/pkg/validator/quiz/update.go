package quiz

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type UpdateRequest struct {
	*models.Quiz
	Token *auth.JWTClaim
}

func (u *UpdateRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"id":          []string{"required", "uuid"},
		"name":        []string{"alpha_space"},
		"description": []string{},
		"userid":      []string{"uuid"},
	}

	options := govalidator.Options{
		Data:  u,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	return validator.MergeUrlValues(validator.GetRulesKey(rules), values)
}

func (u *UpdateRequest) ValidateSemantic(quiz *models.Quiz) []string {
	var errors []string

	if u.Token.IsAdmin {
		return errors
	}

	if quiz.UserID != u.Token.ID {
		errors = append(errors, "user can't execute this operation on quizzes from other users")
	}

	if quiz.IsPublished {
		errors = append(errors, "user can't update published quiz")
	}

	if u.IsPublished {
		errors = append(errors, "user can't publish quiz through update")
	}

	return errors
}
