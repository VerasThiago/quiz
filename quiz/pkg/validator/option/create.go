package option

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type CreateRequest struct {
	*models.Option
	Token *auth.JWTClaim
}

func (c *CreateRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"value":       []string{"required"},
		"questionid":  []string{"required", "uuid"},
		"correctness": []string{"bool"},
	}

	options := govalidator.Options{
		Data:  c,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	errors := []string{}

	if c.Correctness == nil {
		errors = append(errors, "The correctness field is required")
	}

	return append(validator.MergeUrlValues(validator.GetRulesKey(rules), values), errors...)
}

func (c *CreateRequest) ValidateSemantic(quiz *models.Quiz) []string {
	var errors []string

	if c.Token.IsAdmin {
		return errors
	}

	if quiz.UserID != c.Token.ID {
		errors = append(errors, "user can't execute this operation on questions from other users")
	}

	return errors
}
