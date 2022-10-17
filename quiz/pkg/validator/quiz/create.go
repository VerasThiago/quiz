package quiz

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type CreateRequest struct {
	*models.Quiz
	Token *auth.JWTClaim
}

func (c *CreateRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"name":        []string{"required", "alpha_space"},
		"description": []string{"required"},
	}

	options := govalidator.Options{
		Data:  c,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	return validator.MergeUrlValues(validator.GetRulesKey(rules), values)
}
