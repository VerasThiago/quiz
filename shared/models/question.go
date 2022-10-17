package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/liip/sheriff"
	"gorm.io/gorm"
)

const (
	SINGLE_ANS   int = iota + 1 // Idx 1
	MULTIPLE_ANS                // Idx 2
)

type Question struct {
	CreatedAt         time.Time      `groups:"private"`
	UpdatedAt         time.Time      `groups:"private"`
	DeletedAt         gorm.DeletedAt `gorm:"index" groups:"private"`
	ID                string         `json:"id" gorm:"primary_key" groups:"api"`
	Name              string         `json:"name" groups:"api"`
	Description       string         `json:"description" groups:"api"`
	Type              int            `json:"type" groups:"api"`
	WrongAnswerWeight *float64       `json:"wronganswerweight,omitempty" groups:"private"`
	RightAnswerWeight *float64       `json:"rightanswerweight,omitempty" groups:"private"`
	QuizID            string         `gorm:"index" json:"quizid" groups:"api"`
	OptionList        *[]Option      `json:"optionlist,omitempty" groups:"api"`
}

type QuestionBase struct {
	Type              int             `json:"type" groups:"api"`
	WrongAnswerWeight *float64        `json:"wronganswerweight,omitempty" groups:"private"`
	RightAnswerWeight *float64        `json:"rightanswerweight,omitempty" groups:"private"`
	OptionMap         map[string]bool `json:"optionmap" groups:"api"`
}

func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.New().String()
	return nil
}

func (q *Question) SetAnswerWeight() error {
	var right float64 = 0
	var wrong float64 = 0

	if q.OptionList == nil {
		return fmt.Errorf("option list can't be nil")
	}
	for _, option := range *q.OptionList {
		if *option.Correctness {
			right++
		} else {
			wrong++
		}
	}

	right = 1 / right
	wrong = 1 / wrong

	q.RightAnswerWeight = &right
	q.WrongAnswerWeight = &wrong

	return nil
}

func (q *Question) ByteArray(groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, q)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}

func QuestionListByteArray(questionList []*Question, groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, questionList)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}
