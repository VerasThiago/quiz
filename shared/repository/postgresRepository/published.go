package postgresrepository

import "github.com/verasthiago/quiz/shared/models"

func (p *PostgresRepository) GetPublishedQuizList() ([]*models.Quiz, error) {
	var quizList []*models.Quiz
	if record := p.db.Where("is_published=true").Find(&quizList); record.Error != nil {
		return nil, record.Error
	}
	return quizList, nil
}
