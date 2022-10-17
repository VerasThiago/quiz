package quiz

import (
	"github.com/thedevsaddam/govalidator"
	questionValidator "github.com/verasthiago/quiz/quiz/pkg/validator/question"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/validator"
)

type PublishRequest struct {
	QuizID string `json:"id"`
	Token  *auth.JWTClaim
}

func (p *PublishRequest) ValidateSyntax() []string {
	rules := govalidator.MapData{
		"id": []string{"required", "uuid"},
	}

	options := govalidator.Options{
		Data:  p,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	return validator.MergeUrlValues(validator.GetRulesKey(rules), values)
}

func (p *PublishRequest) ValidateSemantic(quiz *models.Quiz) []string {
	var errors []string

	if p.Token.IsAdmin {
		return errors
	}

	if quiz.UserID != p.Token.ID {
		errors = append(errors, "user can't execute this operation on quizzes from other users")
	}

	if quiz.IsPublished {
		errors = append(errors, "user can't publish published quiz")
	}

	if quiz.QuestionList == nil {
		return append(errors, "quiz questions can't be empty")
	}

	if len(*quiz.QuestionList) < 1 || len(*quiz.QuestionList) > 10 {
		return append(errors, "quiz must contain [1-10] questions")
	}

	for _, question := range *quiz.QuestionList {
		if question.OptionList == nil {
			errors = append(errors, "options for a quiz question can't be empty")
		} else if len(*question.OptionList) < 1 || len(*question.OptionList) > 5 {
			errors = append(errors, "options for a quiz question must contain [1-5] options")
		} else {
			errors = append(errors, questionValidator.CheckOptionList(question.Type, question.OptionList)...)
		}
	}

	return errors
}
