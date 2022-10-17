package postgresrepository

import "github.com/verasthiago/quiz/shared/models"

func (p *PostgresRepository) CreateQuestion(question *models.Question) error {
	return p.db.Create(question).Error
}

func (p *PostgresRepository) GetQuestionListByQuizID(quizID string) ([]*models.Question, error) {
	var questionList []*models.Question
	if record := p.db.Where("quiz_id = ?", quizID).Find(&questionList); record.Error != nil {
		return nil, record.Error
	}
	return questionList, nil
}

func (p *PostgresRepository) DeleteQuestionsByQuizID(quizID string) error {
	var err error
	var questionList []*models.Question
	if questionList, err = p.GetQuestionListByQuizID(quizID); err != nil {
		return err
	}
	for idx := range questionList {
		if err := p.DeleteQuestionByID(questionList[idx].ID); err != nil {
			return err
		}
	}
	return nil
}

func (p *PostgresRepository) DeleteQuestionByID(questionID string) error {
	if err := p.db.Where("id = ?", questionID).Delete(&models.Question{}).Error; err != nil {
		return nil
	}
	return p.DeleteOptionsByQuestionID(questionID)
}

func (p *PostgresRepository) UpdateQuestion(question *models.Question) error {
	return p.db.Model(question).Updates(question).Error
}

func (p *PostgresRepository) GetQuestionByID(questionID string) (*models.Question, error) {
	var question models.Question
	if record := p.db.Where("id = ?", questionID).First(&question); record.Error != nil {
		return nil, record.Error
	}
	return &question, nil
}
