package question

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type UpdateRequest struct {
	*models.Question
	Token *auth.JWTClaim
}

func (u *UpdateRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"id":          []string{"required", "uuid"},
		"name":        []string{"alpha_space"},
		"description": []string{},
		"type":        []string{},
		"quizid":      []string{"uuid"},
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
		errors = append(errors, "user can't execute this operation on questions from other users")
	}

	if quiz.IsPublished {
		errors = append(errors, "user can't update question from a published quiz")
	}

	return errors
}
