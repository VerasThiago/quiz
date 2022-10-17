package postgresrepository

import (
	"github.com/verasthiago/quiz/shared/models"
	"gorm.io/gorm"
)

func (p *PostgresRepository) CreateSubmission(submission *models.Submission) error {

	if submission.OptionList == nil {
		return p.db.Create(submission).Error
	}

	submissionOptionList := []*models.SubmissionOption{}
	optionList := *submission.OptionList
	submission.OptionList = nil

	if err := p.db.Create(submission).Error; err != nil {
		return err
	}

	for _, option := range optionList {
		submissionOptionList = append(submissionOptionList, &models.SubmissionOption{SubmissionID: submission.ID, OptionID: option.ID})
	}

	return p.CreateSubmissionOptionList(submissionOptionList)
}

func (p *PostgresRepository) CreateSubmissionOptionList(submissionOptionList []*models.SubmissionOption) error {
	return p.db.Create(submissionOptionList).Error
}

func (p *PostgresRepository) GetSubmissionListByUserID(userID string) ([]*models.Submission, error) {
	var submissionList []*models.Submission
	if record := p.db.Where("user_id = ?", userID).Find(&submissionList); record.Error != nil {
		return nil, record.Error
	}
	return submissionList, nil
}

func (p *PostgresRepository) GetSubmissionListByQuizID(quizID string) ([]*models.Submission, error) {
	var submissionList []*models.Submission
	if record := p.db.Where("quiz_id = ?", quizID).Find(&submissionList); record.Error != nil {
		return nil, record.Error
	}
	return submissionList, nil
}

func (p *PostgresRepository) GetSubmissionByID(submissionID string) (*models.Submission, error) {
	var submission models.Submission
	if record := p.db.Where("id = ?", submissionID).First(&submission); record.Error != nil {
		return nil, record.Error
	}
	return &submission, nil
}

func (p *PostgresRepository) GetFullSubmissionByID(submissionID string) (*models.Submission, error) {
	var submission models.Submission
	if record := p.db.Preload("OptionList").Where("id = ?", submissionID).First(&submission); record.Error != nil {
		return nil, record.Error
	}
	return &submission, nil
}

func (p *PostgresRepository) CheckAlreadySubmittedQuizByUser(userID, quizID string) (bool, error) {
	var submission models.Submission
	record := p.db.Where("user_id = ? AND quiz_id = ?", userID, quizID).First(&submission)

	if record.Error == nil {
		return true, nil
	}

	if record.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	return false, record.Error
}
