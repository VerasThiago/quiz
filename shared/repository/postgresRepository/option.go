package postgresrepository

import (
	"fmt"

	"github.com/verasthiago/quiz/shared/models"
)

func (p *PostgresRepository) CreateOption(option *models.Option) error {
	return p.db.Create(option).Error
}

func (p *PostgresRepository) CreateOptionList(optionList *[]models.Option, questionID string) error {
	for _, option := range *optionList {
		option.QuestionID = questionID
		if err := p.CreateOption(&option); err != nil {
			return err
		}
	}
	return nil
}

func (p *PostgresRepository) GetOptionListByQuestionID(questionID string) ([]*models.Option, error) {
	var optionList []*models.Option
	if record := p.db.Where("question_id = ?", questionID).Find(&optionList); record.Error != nil {
		return nil, record.Error
	}
	return optionList, nil
}

func (p *PostgresRepository) GetOptionByID(questionID string) (*models.Option, error) {
	var option models.Option
	if record := p.db.Where("id = ?", questionID).First(&option); record.Error != nil {
		return nil, record.Error
	}
	return &option, nil
}

func (p *PostgresRepository) DeleteOptionByID(optionID string) error {
	return p.db.Where("id = ?", optionID).Delete(&models.Option{}).Error
}

func (p *PostgresRepository) DeleteOptionsByQuestionID(questionID string) error {
	fmt.Println("")
	return p.db.Where("question_id = ?", questionID).Delete(&models.Option{}).Error
}

func (p *PostgresRepository) UpdateOption(option *models.Option) error {
	return p.db.Model(option).Updates(option).Error
}
