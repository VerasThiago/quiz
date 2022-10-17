package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/liip/sheriff"
	"gorm.io/gorm"
)

type Option struct {
	CreatedAt   time.Time      `groups:"private"`
	UpdatedAt   time.Time      `groups:"private"`
	DeletedAt   gorm.DeletedAt `gorm:"index" groups:"private"`
	ID          string         `json:"id" gorm:"primary_key" groups:"api"`
	Value       string         `json:"value" groups:"api"`
	Correctness *bool          `json:"correctness,omitempty" groups:"user"`
	QuestionID  string         `gorm:"index" json:"questionid" groups:"api"`
}

func (q *Option) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.New().String()
	return nil
}

func (o *Option) ByteArray(groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, o)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}

func OptionListByteArray(optionList []*Option, groups []string) ([]byte, error) {
	opts := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(opts, optionList)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}
