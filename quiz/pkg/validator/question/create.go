package question

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type CreateRequest struct {
	Question   *models.Question `json:"question"`
	OptionList *[]models.Option `json:"optionlist"`
	Token      *auth.JWTClaim
}

func (c *CreateRequest) ValidateSyntax() []string {
	return append(checkQuestion(c.Question), CheckOptionList(c.Question.Type, c.OptionList)...)
}

func (c *CreateRequest) ValidateSemantic(quiz *models.Quiz) []string {
	var errors []string

	if c.Token.IsAdmin {
		return errors
	}

	if quiz.UserID != c.Token.ID {
		errors = append(errors, "user can't execute this operation on quizzes from other users")
	}

	return errors
}

func checkQuestion(question *models.Question) []string {
	if question == nil {
		return []string{"The question field is required"}
	}

	rules := govalidator.MapData{
		"name":        []string{"required"},
		"description": []string{"required"},
		"type":        []string{"required", "numeric_between:1,2"},
		"quizid":      []string{"required", "uuid"},
	}

	options := govalidator.Options{
		Data:  question,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	return validator.MergeUrlValues(validator.GetRulesKey(rules), values)
}

func CheckOptionList(questionType int, optionList *[]models.Option) []string {
	if optionList == nil {
		return []string{}
	}

	var errors []string
	var trueAnswers int = 0

	for _, option := range *optionList {
		rules := govalidator.MapData{
			"value":       []string{"required"},
			"correctness": []string{"bool"},
		}

		options := govalidator.Options{
			Data:  &option,
			Rules: rules,
		}

		values := govalidator.New(options).ValidateStruct()
		errors = append(errors, validator.MergeUrlValues(validator.GetRulesKey(rules), values)...)

		if option.Correctness == nil {
			errors = append(errors, "The correctness field is required")
		} else if *option.Correctness {
			trueAnswers++
		}
	}

	if trueAnswers == 0 {
		errors = append(errors, "The question must have at least 1 right option (correctness = true)")
	}

	if questionType == models.SINGLE_ANS && trueAnswers > 1 {
		errors = append(errors, "Single answer questions must have only 1 right option (correctness = true)")
	}

	return errors
}
