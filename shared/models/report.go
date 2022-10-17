package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/liip/sheriff"
	"gorm.io/gorm"
)

type Report struct {
	ID                 string            `json:"id" gorm:"primary_key" groups:"api"`
	CreatedAt          time.Time         `groups:"private"`
	UpdatedAt          time.Time         `groups:"private"`
	DeletedAt          gorm.DeletedAt    `gorm:"index" groups:"private"`
	SubmissionID       string            `json:"submissionid" groups:"api"`
	TotalScore         float64           `json:"totalscore" groups:"api"`
	ReportQuestionList *[]ReportQuestion `json:"reportquestionlist,omitempty" groups:"api"`
}

type ReportQuestion struct {
	ReportID      string  `json:"reportid" groups:"full"`
	QuestionID    string  `json:"questionid" groups:"api"`
	QuestionScore float64 `json:"questionscore" groups:"api"`
}

func (r *Report) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New().String()
	return nil
}

func (r *Report) ByteArray(groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, r)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}

func ReportListByteArray(reportList []*Report, groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, reportList)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}
