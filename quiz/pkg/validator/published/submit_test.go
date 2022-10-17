package published_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/verasthiago/quiz/quiz/pkg/validator/published"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
)

var (
	TRUE   = true
	FALSE  = false
	FLOAT1 = 0.3
	FLOAT2 = 1.0
)

func TestSyntax(t *testing.T) {
	Convey("Syntax analisys", t, func() {
		dataList := []struct {
			Input          published.SubmitRequest
			ExpectedOutput []string
		}{
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				ExpectedOutput: []string{"The questionList field is required"},
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{},
				},
				ExpectedOutput: []string(nil),
			},
		}

		for _, data := range dataList {
			So(data.Input.ValidateSyntax(), ShouldResemble, data.ExpectedOutput)
		}

	})
}

func TestSemantic(t *testing.T) {
	Convey("Semantic analisys", t, func() {
		dataList := []struct {
			Input            published.SubmitRequest
			QuizBase         *models.QuizBase
			AlreadySubmitted bool
			ExpectedOutput   []string
		}{
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
				},
				QuizBase: &models.QuizBase{
					IsPublished: true,
				},
				AlreadySubmitted: false,
				ExpectedOutput:   []string{"questionlist can't be nil"},
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{},
				},
				QuizBase: &models.QuizBase{
					IsPublished: true,
				},
				AlreadySubmitted: false,
				ExpectedOutput:   []string(nil),
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{},
				},
				QuizBase: &models.QuizBase{
					IsPublished: true,
				},
				AlreadySubmitted: true,
				ExpectedOutput:   []string{"user already submitted that quiz"},
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{},
				},
				QuizBase: &models.QuizBase{
					IsPublished: false,
				},
				AlreadySubmitted: false,
				ExpectedOutput:   []string{"user can't submit on unpublished quiz"},
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{
						{},
					},
				},
				QuizBase: &models.QuizBase{
					IsPublished: true,
				},
				AlreadySubmitted: false,
				ExpectedOutput:   []string{"question id can't be nil"},
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{
						{
							QuestionID: "41242ef1-bd1e-4c7d-bbce-38d58a32924c",
						},
					},
				},
				QuizBase: &models.QuizBase{
					IsPublished: true,
				},
				AlreadySubmitted: false,
				ExpectedOutput:   []string{"optionlist can't be nil"},
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{
						{
							QuestionID: "41242ef1-bd1e-4c7d-bbce-38d58a32924c",
							OptionList: &[]string{
								"8633c6e0-9174-45e0-a4a4-3c9053e1e8b0",
								"75b55ead-1d13-46cd-b8d0-f75f1226f333",
								"a6302e6d-dca8-41b2-9f99-584a4ac196fa",
								"6c9e0caf-6531-49ad-b1cb-e11ee02f8673",
								"47a4ede1-5d5f-4085-980d-64e70906ce5c",
							},
						},
					},
				},
				QuizBase: &models.QuizBase{
					IsPublished: true,
					QuestionMap: map[string]models.QuestionBase{
						"41242ef1-bd1e-4c7d-bbce-38d58a32924c": {
							Type:              1,
							WrongAnswerWeight: &FLOAT1,
							RightAnswerWeight: &FLOAT2,
							OptionMap: map[string]bool{
								"8633c6e0-9174-45e0-a4a4-3c9053e1e8b0": false,
								"75b55ead-1d13-46cd-b8d0-f75f1226f333": false,
								"a6302e6d-dca8-41b2-9f99-584a4ac196fa": false,
								"6c9e0caf-6531-49ad-b1cb-e11ee02f8673": true,
								"47a4ede1-5d5f-4085-980d-64e70906ce5c": false,
							},
						},
					},
				},
				AlreadySubmitted: false,
				ExpectedOutput:   []string{"user can't submit multiple answers for single answer questions"},
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{
						{
							QuestionID: "41242ef1-bd1e-4c7d-bbce-38d58a32924c",
							OptionList: &[]string{
								"47a4ede1-5d5f-4085-980d-64e70906ce2c",
							},
						},
					},
				},
				QuizBase: &models.QuizBase{
					IsPublished: true,
					QuestionMap: map[string]models.QuestionBase{
						"41242ef1-bd1e-4c7d-bbce-38d58a32924c": {
							Type:              1,
							WrongAnswerWeight: &FLOAT1,
							RightAnswerWeight: &FLOAT2,
							OptionMap: map[string]bool{
								"8633c6e0-9174-45e0-a4a4-3c9053e1e8b0": false,
								"75b55ead-1d13-46cd-b8d0-f75f1226f333": false,
								"a6302e6d-dca8-41b2-9f99-584a4ac196fa": false,
								"6c9e0caf-6531-49ad-b1cb-e11ee02f8673": true,
								"47a4ede1-5d5f-4085-980d-64e70906ce5c": false,
							},
						},
					},
				},
				AlreadySubmitted: false,
				ExpectedOutput:   []string{"option id not found"},
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{
						{
							QuestionID: "41242ef1-bd1e-4c7d-bbce-38d58a329243",
							OptionList: &[]string{
								"75b55ead-1d13-46cd-b8d0-f75f1226f333",
							},
						},
					},
				},
				QuizBase: &models.QuizBase{
					IsPublished: true,
					QuestionMap: map[string]models.QuestionBase{
						"41242ef1-bd1e-4c7d-bbce-38d58a32924c": {
							Type:              1,
							WrongAnswerWeight: &FLOAT1,
							RightAnswerWeight: &FLOAT2,
							OptionMap: map[string]bool{
								"8633c6e0-9174-45e0-a4a4-3c9053e1e8b0": false,
								"75b55ead-1d13-46cd-b8d0-f75f1226f333": false,
								"a6302e6d-dca8-41b2-9f99-584a4ac196fa": false,
								"6c9e0caf-6531-49ad-b1cb-e11ee02f8673": true,
								"47a4ede1-5d5f-4085-980d-64e70906ce5c": false,
							},
						},
					},
				},
				AlreadySubmitted: false,
				ExpectedOutput:   []string{"question id not found"},
			},
			{
				Input: published.SubmitRequest{
					QuizID: "8af4f04f-434d-4570-8727-69c4fac340c6",
					Token: &auth.JWTClaim{
						ID:      "79ce8ef1-bd1e-4c7d-bbce-38d58a32924c",
						IsAdmin: false,
					},
					QuestionList: &[]published.Question{
						{
							QuestionID: "41242ef1-bd1e-4c7d-bbce-38d58a32924c",
							OptionList: &[]string{
								"75b55ead-1d13-46cd-b8d0-f75f1226f333",
							},
						},
					},
				},
				QuizBase: &models.QuizBase{
					IsPublished: true,
					QuestionMap: map[string]models.QuestionBase{
						"41242ef1-bd1e-4c7d-bbce-38d58a32924c": {
							Type:              1,
							WrongAnswerWeight: &FLOAT1,
							RightAnswerWeight: &FLOAT2,
							OptionMap: map[string]bool{
								"8633c6e0-9174-45e0-a4a4-3c9053e1e8b0": false,
								"75b55ead-1d13-46cd-b8d0-f75f1226f333": false,
								"a6302e6d-dca8-41b2-9f99-584a4ac196fa": false,
								"6c9e0caf-6531-49ad-b1cb-e11ee02f8673": true,
								"47a4ede1-5d5f-4085-980d-64e70906ce5c": false,
							},
						},
					},
				},
				AlreadySubmitted: false,
				ExpectedOutput:   []string(nil),
			},
		}

		for _, data := range dataList {
			So(data.Input.ValidateSemantic(data.QuizBase, data.AlreadySubmitted), ShouldResemble, data.ExpectedOutput)
		}

	})
}
