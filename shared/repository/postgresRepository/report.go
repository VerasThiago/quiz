package postgresrepository

import (
	"math"

	"github.com/verasthiago/quiz/shared/models"
)

func (p *PostgresRepository) GetFullReportBySubmissionID(submissionID string) (*models.Report, error) {
	var report models.Report
	if record := p.db.Preload("ReportQuestionList").Where("submission_id = ?", submissionID).First(&report); record.Error != nil {
		return nil, record.Error
	}
	return &report, nil
}

func (p *PostgresRepository) CreateReportBySubmissionID(submissionID string) error {
	var err error
	var quizBase *models.QuizBase
	var submission *models.Submission
	var subOptionsMap map[string]bool
	var reportQuestionList []*models.ReportQuestion
	var report models.Report = models.Report{SubmissionID: submissionID}
	var totalScore float64 = 0
	var totalQuestions float64

	if submission, err = p.GetFullSubmissionByID(submissionID); err != nil {
		return err
	}

	if quizBase, err = p.GetQuizBaseByID(submission.QuizID); err != nil {
		return err
	}

	if subOptionsMap, err = submission.GetOptionsMap(); err != nil {
		return err
	}

	for questionID, questionBase := range quizBase.QuestionMap {
		var questionScore float64 = 0
		for optionID, correctness := range questionBase.OptionMap {
			if _, ok := subOptionsMap[optionID]; ok {
				if correctness {
					questionScore += *questionBase.RightAnswerWeight
				} else {
					questionScore -= *questionBase.WrongAnswerWeight
				}
			}
		}
		questionScore = math.Round(questionScore*100) / 100
		reportQuestionList = append(reportQuestionList,
			&models.ReportQuestion{
				QuestionID:    questionID,
				QuestionScore: questionScore,
			},
		)
		totalScore += questionScore
	}

	totalQuestions = float64(len(quizBase.QuestionMap))
	report.TotalScore = totalScore / totalQuestions
	report.TotalScore = math.Round(report.TotalScore*100) / 100

	if err := p.db.Create(&report).Error; err != nil {
		return err
	}

	for _, reportQuestion := range reportQuestionList {
		reportQuestion.ReportID = report.ID
	}

	return p.CreateReportQuestionList(reportQuestionList)
}

func (p *PostgresRepository) CreateReportQuestionList(reportQuestion []*models.ReportQuestion) error {
	return p.db.Create(reportQuestion).Error
}
