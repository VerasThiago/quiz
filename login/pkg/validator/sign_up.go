package validator

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type SignUpRequest struct {
	*models.User
}

func (s *SignUpRequest) Validate() []string {
	rules := govalidator.MapData{
		"name":     []string{"required", "alpha_dash"},
		"username": []string{"required", "alpha_dash"},
		"email":    []string{"required", "email"},
		"password": []string{"required"},
	}

	options := govalidator.Options{
		Data:  s,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	return validator.MergeUrlValues([]string{"name", "username", "email", "password"}, values)
}
