package published

import (
	"fmt"

	"github.com/thedevsaddam/govalidator"
	"github.com/verasthiago/quiz/shared/auth"
	"github.com/verasthiago/quiz/shared/models"
	"github.com/verasthiago/quiz/shared/utils"
	"github.com/verasthiago/quiz/shared/validator"
)

type Question struct {
	QuestionID string    `json:"id"`
	OptionList *[]string `json:"optionlist,omitempty"`
}

type SubmitRequest struct {
	QuizID       string      `json:"id"`
	QuestionList *[]Question `json:"questionList,omitempty"`
	Token        *auth.JWTClaim
}

type SubmitRequestMap map[string][]string

func (s *SubmitRequest) ValidateSyntax() []string {
	return append(s.validateBase(), s.validateQuestionList()...)
}

func (s *SubmitRequest) ValidateSemantic(quizBase *models.QuizBase, alreadySubmitted bool) []string {
	var err error
	var errors []string
	var requestMap *SubmitRequestMap

	if !quizBase.IsPublished {
		errors = append(errors, "user can't submit on unpublished quiz")
	}

	if alreadySubmitted {
		errors = append(errors, "user already submitted that quiz")
	}

	if requestMap, err = s.GetFilteredSubmitRequestMap(quizBase); err != nil {
		return append(errors, err.Error())
	}

	for questionID, question := range quizBase.QuestionMap {
		if question.Type == models.SINGLE_ANS && len((*requestMap)[questionID]) > 1 {
			errors = append(errors, "user can't submit multiple answers for single answer questions")
		}
	}

	return errors
}

func (s *SubmitRequest) validateBase() []string {
	rules := govalidator.MapData{
		"id": []string{"required", "uuid"},
	}

	options := govalidator.Options{
		Data:  s,
		Rules: rules,
	}

	values := govalidator.New(options).ValidateStruct()
	errors := validator.MergeUrlValues(validator.GetRulesKey(rules), values)

	if s.QuestionList == nil {
		errors = append(errors, "The questionList field is required")
	}

	return errors
}

func (s *SubmitRequest) validateQuestionList() []string {
	var errors []string

	if s.QuestionList == nil {
		return errors
	}

	for _, question := range *s.QuestionList {
		if question.QuestionID == "" {
			errors = append(errors, "The question id field is required")
		} else if !utils.IsUUID(question.QuestionID) {
			errors = append(errors, "The question id field must contain valid UUID")
		}

		if question.OptionList == nil {
			errors = append(errors, "The optionlist field is required")
		} else {
			for _, optionID := range *question.OptionList {
				if !utils.IsUUID(optionID) {
					errors = append(errors, "The option id field must contain valid UUID")
				}
			}
		}
	}

	return errors
}

func (s *SubmitRequest) GetSubmitRequestMap() (*SubmitRequestMap, error) {
	var requestMap SubmitRequestMap = make(SubmitRequestMap)

	if s.QuestionList == nil {
		return nil, fmt.Errorf("questionlist can't be nil")
	}

	for _, question := range *s.QuestionList {
		if question.QuestionID == "" {
			return nil, fmt.Errorf("question id can't be nil")
		}
		if _, ok := requestMap[question.QuestionID]; ok {
			return nil, fmt.Errorf("duplicated question id")
		}
		if question.OptionList == nil {
			return nil, fmt.Errorf("optionlist can't be nil")
		}
		requestMap[question.QuestionID] = *question.OptionList
	}

	return &requestMap, nil
}

func (s *SubmitRequest) GetFilteredSubmitRequestMap(quizBase *models.QuizBase) (*SubmitRequestMap, error) {
	var err error
	var requestMap *SubmitRequestMap
	var visitedOption map[string]bool = make(map[string]bool)

	if requestMap, err = s.GetSubmitRequestMap(); err != nil {
		return nil, err
	}

	for questionID, OptionList := range *requestMap {
		var filteredOptions []string

		if _, ok := quizBase.QuestionMap[questionID]; !ok {
			return nil, fmt.Errorf("question id not found")
		}

		for _, optionID := range OptionList {
			if _, ok := quizBase.QuestionMap[questionID].OptionMap[optionID]; ok {
				if _, ok := visitedOption[optionID]; ok {
					return nil, fmt.Errorf("duplicated option id")
				}
				visitedOption[optionID] = true
				filteredOptions = append(filteredOptions, optionID)
			} else {
				return nil, fmt.Errorf("option id not found")
			}
		}
		(*requestMap)[questionID] = filteredOptions
	}

	return requestMap, nil
}

func (s *SubmitRequest) GetSubmissionBase(quizBase *models.QuizBase) (*models.Submission, error) {
	var err error
	var requestMap *SubmitRequestMap
	var submission models.Submission
	var optionList []models.Option

	if requestMap, err = s.GetFilteredSubmitRequestMap(quizBase); err != nil {
		return nil, err
	}

	for _, OptionList := range *requestMap {
		for _, optionID := range OptionList {
			optionList = append(optionList, models.Option{ID: optionID})
		}
	}

	if len(optionList) > 0 {
		submission.OptionList = &optionList
	}

	return &submission, nil
}
