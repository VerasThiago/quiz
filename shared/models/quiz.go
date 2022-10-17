package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/liip/sheriff"
	"gorm.io/gorm"
)

type Quiz struct {
	CreatedAt    time.Time      `groups:"private"`
	UpdatedAt    time.Time      `groups:"private"`
	DeletedAt    gorm.DeletedAt `gorm:"index" groups:"private"`
	ID           string         `json:"id" gorm:"primary_key" groups:"api"`
	Name         string         `json:"name" groups:"api"`
	Description  string         `json:"description" groups:"api"`
	UserID       string         `json:"userid" groups:"api"`
	IsPublished  bool           `json:"ispublished" groups:"api"`
	QuestionList *[]Question    `json:"questionList,omitempty" groups:"api"`
}

type QuizBase struct {
	ID          string                  `json:"id" groups:"api"`
	IsPublished bool                    `json:"ispublished" groups:"api"`
	QuestionMap map[string]QuestionBase `json:"questionmap" groups:"api"`
}

func (q *Quiz) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.New().String()
	q.IsPublished = false
	return nil
}

func (q *Quiz) GetQuizBase() *QuizBase {
	var quizBase QuizBase
	quizBase.QuestionMap = make(map[string]QuestionBase)
	quizBase.IsPublished = q.IsPublished

	for _, question := range *q.QuestionList {
		var questionBase QuestionBase
		questionBase.Type = question.Type
		questionBase.OptionMap = make(map[string]bool)
		questionBase.RightAnswerWeight = question.RightAnswerWeight
		questionBase.WrongAnswerWeight = question.WrongAnswerWeight

		for _, option := range *question.OptionList {
			questionBase.OptionMap[option.ID] = *option.Correctness
		}
		quizBase.QuestionMap[question.ID] = questionBase
	}
	return &quizBase
}

func (q *Quiz) ByteArray(groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, q)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}

func QuizListByteArray(quizList []*Quiz, groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, quizList)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}
