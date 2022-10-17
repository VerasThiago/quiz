package postgresrepository

import (
	"github.com/verasthiago/quiz/shared/models"
)

func (p *PostgresRepository) CreateQuiz(quiz *models.Quiz) error {
	return p.db.Create(quiz).Error
}

func (p *PostgresRepository) GetQuizListByUserID(userID string) ([]*models.Quiz, error) {
	var quizList []*models.Quiz
	if record := p.db.Where("user_id = ?", userID).Find(&quizList); record.Error != nil {
		return nil, record.Error
	}
	return quizList, nil
}

func (p *PostgresRepository) GetQuizByID(quizID string) (*models.Quiz, error) {
	var quiz models.Quiz
	if record := p.db.Where("id = ?", quizID).First(&quiz); record.Error != nil {
		return nil, record.Error
	}
	return &quiz, nil
}

func (p *PostgresRepository) GetFullQuizByID(quizID string) (*models.Quiz, error) {
	var quiz models.Quiz
	if record := p.db.Preload("QuestionList").Preload("QuestionList.OptionList").Where("id = ?", quizID).First(&quiz); record.Error != nil {
		return nil, record.Error
	}
	return &quiz, nil
}

func (p *PostgresRepository) GetQuizBaseByID(quizID string) (*models.QuizBase, error) {
	var err error
	var quiz *models.Quiz
	if quiz, err = p.GetFullQuizByID(quizID); err != nil {
		return nil, err
	}
	return quiz.GetQuizBase(), nil
}

func (p *PostgresRepository) GetQuizByQuestionID(questionID string) (*models.Quiz, error) {
	var quiz models.Quiz
	if record := p.db.Joins("LEFT JOIN questions ON quizzes.id = questions.quiz_id").Where("questions.id = ?", questionID).First(&quiz); record.Error != nil {
		return nil, record.Error
	}
	return &quiz, nil
}

func (p *PostgresRepository) GetQuizByOptionID(optionID string) (*models.Quiz, error) {
	var quiz models.Quiz
	if record := p.db.Joins("LEFT JOIN questions ON quizzes.id = questions.quiz_id").Joins("LEFT JOIN options ON options.question_id = questions.id").Where("options.id = ?", optionID).First(&quiz); record.Error != nil {
		return nil, record.Error
	}
	return &quiz, nil
}

func (p *PostgresRepository) DeleteQuizByID(quizID string) error {
	if err := p.db.Where("id = ?", quizID).Delete(&models.Quiz{}).Error; err != nil {
		return nil
	}
	return p.DeleteQuestionsByQuizID(quizID)
}

func (p *PostgresRepository) DeleteQuizzesByUserID(userID string) error {
	quizList, err := p.GetQuizListByUserID(userID)
	if err != nil {
		return err
	}
	for idx := range quizList {
		if err := p.DeleteQuizByID(quizList[idx].ID); err != nil {
			return err
		}
	}
	return nil
}

func (p *PostgresRepository) UpdateQuiz(quiz *models.Quiz) error {
	return p.db.Model(quiz).Updates(quiz).Error
}

func (p *PostgresRepository) PublishQuiz(quiz *models.Quiz) error {
	for _, question := range *quiz.QuestionList {
		if err := question.SetAnswerWeight(); err != nil {
			return err
		}
		question.OptionList = nil
		if err := p.db.Save(question).Error; err != nil {
			return err
		}
	}
	return p.db.Model(&models.Quiz{}).Where("id = ?", quiz.ID).Update("is_published", true).Error
}
