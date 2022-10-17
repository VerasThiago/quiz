package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/liip/sheriff"
	"gorm.io/gorm"
)

type Submission struct {
	CreatedAt  time.Time      `groups:"private"`
	UpdatedAt  time.Time      `groups:"private"`
	DeletedAt  gorm.DeletedAt `gorm:"index" groups:"private"`
	ID         string         `json:"id" gorm:"primary_key" groups:"api"`
	UserID     string         `json:"userid" groups:"api"`
	QuizID     string         `json:"quizid" groups:"api"`
	OptionList *[]Option      `json:"optionlist,omitempty" groups:"api" gorm:"many2many:submission_options;"`
}

type SubmissionOption struct {
	SubmissionID string `json:"submissionid" groups:"api"`
	OptionID     string `json:"optionid" groups:"api"`
}

func (s *Submission) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return nil
}

func (s *Submission) GetOptionsMap() (map[string]bool, error) {
	var optionsMap map[string]bool = make(map[string]bool)

	if s.OptionList == nil {
		return nil, fmt.Errorf("OptionList can't be nil")
	}

	for _, option := range *s.OptionList {
		optionsMap[option.ID] = true
	}
	return optionsMap, nil
}

func (s *Submission) ByteArray(groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, s)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}

func SubmissionListByteArray(submissionList []*Submission, groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, submissionList)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}
