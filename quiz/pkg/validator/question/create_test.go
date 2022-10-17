package question_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/verasthiago/quiz/quiz/pkg/validator/question"
	"github.com/verasthiago/quiz/shared/models"
)

var (
	TRUE  = true
	FALSE = false
)

func TestSyntax(t *testing.T) {
	Convey("Syntax analisys", t, func() {
		dataList := []struct {
			Input          question.CreateRequest
			ExpectedOutput []string
		}{
			{
				Input: question.CreateRequest{
					Question: &models.Question{
						Name:        "1 + 1?",
						Description: "Hard question",
						QuizID:      "8af4f04f-434d-4570-8727-69c4fac340c6",
						Type:        1,
					},
				},
				ExpectedOutput: []string(nil),
			},
			{
				Input: question.CreateRequest{
					Question: &models.Question{
						Name:        "1 + 1?",
						Description: "Hard question",
						QuizID:      "8af4f04f-434d-4570-8727-6940c6",
						Type:        1,
					},
				},
				ExpectedOutput: []string{"The quizid field must contain valid UUID"},
			},
			{
				Input: question.CreateRequest{
					Question: &models.Question{
						Name:        "1 + 1?",
						Description: "Hard question",
						QuizID:      "8af4f04f-434d-4570-8727-69c4fac340c6",
						Type:        1,
					},
					OptionList: &[]models.Option{
						{
							Value:       "1",
							Correctness: &TRUE,
						},
						{
							Value:       "2",
							Correctness: &FALSE,
						},
						{
							Value:       "3",
							Correctness: &TRUE,
						},
					},
				},
				ExpectedOutput: []string{"Single answer questions must have only 1 right option (correctness = true)"},
			},
			{
				Input: question.CreateRequest{
					Question: &models.Question{
						Name:        "1 + 1?",
						Description: "Hard question",
						QuizID:      "8af4f04f-434d-4570-8727-69c4fac340c6",
						Type:        1,
					},
					OptionList: &[]models.Option{
						{
							Value:       "1",
							Correctness: &TRUE,
						},
						{
							Value:       "2",
							Correctness: &FALSE,
						},
						{
							Value:       "3",
							Correctness: &FALSE,
						},
						{
							Value:       "4",
							Correctness: &FALSE,
						},
						{
							Value:       "5",
							Correctness: &FALSE,
						},
						{
							Value:       "6",
							Correctness: &FALSE,
						},
					},
				},
				ExpectedOutput: []string(nil),
			},
		}

		for _, data := range dataList {
			So(data.Input.ValidateSyntax(), ShouldResemble, data.ExpectedOutput)
		}

	})
}
