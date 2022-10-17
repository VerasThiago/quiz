package option

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type UpdateRequest struct {
	*models.Option
	Token *auth.JWTClaim
}

func (u *UpdateRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"id":          []string{"required", "uuid"},
		"value":       []string{},
		"questionid":  []string{"uuid"},
		"correctness": []string{"bool"},
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
		return nil
	}

	if quiz.UserID != u.Token.ID {
		errors = append(errors, "user can't execute this operation on options from other users")
	}

	if quiz.IsPublished {
		errors = append(errors, "user can't update option from a published quiz")
	}

	return errors
}
