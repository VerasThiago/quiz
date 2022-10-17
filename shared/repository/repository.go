package repository

import (
	"github.com/verasthiago/quiz/shared/models"
)

type Repository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	DeleteUser(userID string) error
	UpdateUser(user *models.User) error

	CreateQuiz(quiz *models.Quiz) error
	GetQuizListByUserID(userID string) ([]*models.Quiz, error)
	GetQuizByID(quizID string) (*models.Quiz, error)
	GetQuizByOptionID(optionID string) (*models.Quiz, error)
	GetQuizBaseByID(quizID string) (*models.QuizBase, error)
	GetQuizByQuestionID(questionID string) (*models.Quiz, error)
	GetFullQuizByID(quizID string) (*models.Quiz, error)
	DeleteQuizByID(quizID string) error
	UpdateQuiz(quiz *models.Quiz) error
	PublishQuiz(quiz *models.Quiz) error

	GetPublishedQuizList() ([]*models.Quiz, error)

	GetSubmissionListByUserID(userID string) ([]*models.Submission, error)
	GetSubmissionListByQuizID(quizID string) ([]*models.Submission, error)
	GetSubmissionByID(submissionID string) (*models.Submission, error)
	GetFullSubmissionByID(submissionID string) (*models.Submission, error)
	CreateSubmission(submission *models.Submission) error
	CreateSubmissionOptionList(submissionOptionList []*models.SubmissionOption) error
	CheckAlreadySubmittedQuizByUser(userID, quizID string) (bool, error)

	GetFullReportBySubmissionID(submissionID string) (*models.Report, error)
	CreateReportBySubmissionID(submissionID string) error

	CreateQuestion(question *models.Question) error
	DeleteQuestionsByQuizID(quizID string) error
	DeleteQuestionByID(quizID string) error
	UpdateQuestion(question *models.Question) error
	GetQuestionListByQuizID(quizID string) ([]*models.Question, error)
	GetQuestionByID(questionID string) (*models.Question, error)

	CreateOption(option *models.Option) error
	CreateOptionList(optionList *[]models.Option, questionID string) error
	GetOptionListByQuestionID(questionID string) ([]*models.Option, error)
	GetOptionByID(optionID string) (*models.Option, error)
	DeleteOptionByID(optionID string) error
	UpdateOption(option *models.Option) error
}
