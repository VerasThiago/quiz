package published

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/validator"
)

type OpenRequest struct {
	QuizID string `json:"id"`
	Token  *auth.JWTClaim
}

func (o *OpenRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"id": []string{"required", "uuid"},
	}

	options := govalidator.Options{
		Data:  o,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	return validator.MergeUrlValues(validator.GetRulesKey(rules), values)
}

func (o *OpenRequest) ValidateSemantic(alreadySubmitted bool) []string {
	var errors []string

	if o.Token.IsAdmin {
		return nil
	}

	if alreadySubmitted {
		errors = append(errors, "user already solved that quiz")
	}

	return errors
}
