package quiz_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/verasthiago/quiz/quiz/pkg/validator/quiz"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
)

var (
	TRUE  = true
	FALSE = false
)

func TestSemantic(t *testing.T) {
	Convey("Semantic analisys", t, func() {
		dataList := []struct {
			Input          quiz.PublishRequest
			Quiz           *models.Quiz
			ExpectedOutput []string
		}{
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				Quiz: &models.Quiz{
					UserID:       "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
					IsPublished:  false,
					QuestionList: &[]models.Question{},
				},
				ExpectedOutput: []string{"quiz must contain [1-10] questions"},
			},
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				Quiz: &models.Quiz{
					UserID:       "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
					IsPublished:  true,
					QuestionList: &[]models.Question{},
				},
				ExpectedOutput: []string{
					"user can't publish published quiz",
					"quiz must contain [1-10] questions",
				},
			},
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				Quiz: &models.Quiz{
					UserID:       "79ce8ef1-bd1e-4c7d-bbce-1233123122",
					IsPublished:  true,
					QuestionList: &[]models.Question{},
				},
				ExpectedOutput: []string{
					"user can't execute this operation on quizzes from other users",
					"user can't publish published quiz",
					"quiz must contain [1-10] questions",
				},
			},
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: true,
					},
				},
				Quiz: &models.Quiz{
					UserID:       "79ce8ef1-bd1e-4c7d-bbce-1233123122",
					IsPublished:  true,
					QuestionList: &[]models.Question{},
				},
				ExpectedOutput: []string(nil),
			},
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				Quiz: &models.Quiz{
					UserID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
					IsPublished: false,
					QuestionList: &[]models.Question{
						{
							Type:       1,
							OptionList: &[]models.Option{},
						},
					},
				},
				ExpectedOutput: []string{"options for a quiz question must contain [1-5] options"},
			},
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				Quiz: &models.Quiz{
					UserID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
					IsPublished: false,
					QuestionList: &[]models.Question{
						{
							Name: "1 + 1:",
							Type: 1,
							OptionList: &[]models.Option{
								{
									Value:       "2",
									Correctness: &FALSE,
								},
							},
						},
					},
				},
				ExpectedOutput: []string{"The question must have at least 1 right option (correctness = true)"},
			},
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				Quiz: &models.Quiz{
					UserID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
					IsPublished: false,
					QuestionList: &[]models.Question{
						{
							Name: "1 + 1:",
							Type: 1,
							OptionList: &[]models.Option{
								{
									Value:       "2",
									Correctness: &FALSE,
								},
								{
									Value:       "3",
									Correctness: &TRUE,
								},
								{
									Value:       "4",
									Correctness: &TRUE,
								},
							},
						},
					},
				},
				ExpectedOutput: []string{"Single answer questions must have only 1 right option (correctness = true)"},
			},
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				Quiz: &models.Quiz{
					UserID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
					IsPublished: false,
					QuestionList: &[]models.Question{
						{
							Name: "1 + 1:",
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "2",
									Correctness: &FALSE,
								},
								{
									Value:       "3",
									Correctness: &TRUE,
								},
								{
									Value:       "4",
									Correctness: &TRUE,
								},
							},
						},
					},
				},
				ExpectedOutput: []string(nil),
			},
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				Quiz: &models.Quiz{
					UserID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
					IsPublished: false,
					QuestionList: &[]models.Question{
						{
							Name: "1 + 1:",
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &FALSE,
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
									Correctness: &TRUE,
								},
								{
									Value:       "5",
									Correctness: &TRUE,
								},
								{
									Value:       "6",
									Correctness: &FALSE,
								},
							},
						},
					},
				},
				ExpectedOutput: []string{"options for a quiz question must contain [1-5] options"},
			},
			{
				Input: quiz.PublishRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				Quiz: &models.Quiz{
					UserID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
					IsPublished: false,
					QuestionList: &[]models.Question{
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
						{
							Type: 2,
							OptionList: &[]models.Option{
								{
									Value:       "1",
									Correctness: &TRUE,
								},
							},
						},
					},
				},
				ExpectedOutput: []string{"quiz must contain [1-10] questions"},
			},
		}

		for _, data := range dataList {
			So(data.Input.ValidateSemantic(data.Quiz), ShouldResemble, data.ExpectedOutput)
		}

	})
}
