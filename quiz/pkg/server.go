package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/quiz/pkg/builder"
	"github.com/verasthiago/quiz/quiz/pkg/handlers/option"
	"github.com/verasthiago/quiz/quiz/pkg/handlers/published"
	"github.com/verasthiago/quiz/quiz/pkg/handlers/question"
	"github.com/verasthiago/quiz/quiz/pkg/handlers/quiz"
	"github.com/verasthiago/quiz/quiz/pkg/handlers/submission"
	"github.com/verasthiago/quiz/quiz/pkg/middlewares"
)

type Server struct {
	builder.Builder

	CreateUserQuizAPI  quiz.CreateQuizAPI
	ListUserQuizAPI    quiz.CreateQuizAPI
	DeleteUserQuizAPI  quiz.DeleteQuizAPI
	UpdateUserQuizAPI  quiz.UpdateQuizAPI
	PublishUserQuizAPI quiz.PublishQuizAPI

	CreateUserQuestionAPI question.CreateQuestionAPI
	ListUserQuestionAPI   question.ListQuestionAPI
	DeleteUserQuestionAPI question.DeleteQuestionAPI
	UpdateUserQuestionAPI question.UpdateQuestionAPI

	CreateUserOptionAPI option.CreateOptionAPI
	ListUserOptionAPI   option.ListOptionAPI
	DeleteUserOptionAPI option.DeleteOptionAPI
	UpdateUserOptionAPI option.UpdateOptionAPI

	ListPublishedQuizAPI   published.ListPublishedQuizAPI
	OpenPublishedQuizAPI   published.OpenPublishedQuizAPI
	SubmitPublishedQuizAPI published.SubmitPublishedQuizAPI

	ListUserSubmissionAPI  submission.ListUserSubmissionAPI
	ListQuizSubmissionAPI  submission.ListQuizSubmissionAPI
	GetSubmissionReportAPI submission.GetSubmissionReportAPI

	AuthAPI middlewares.AuthUserAPI
}

func (s *Server) InitFromBuilder(builder builder.Builder) *Server {
	s.Builder = builder

	s.CreateUserQuizAPI = new(quiz.CreateQuizHandler).InitFromBuilder(builder)
	s.ListUserQuizAPI = new(quiz.ListQuizHandler).InitFromBuilder(builder)
	s.DeleteUserQuizAPI = new(quiz.DeleteQuizHandler).InitFromBuilder(builder)
	s.UpdateUserQuizAPI = new(quiz.UpdateQuizHandler).InitFromBuilder(builder)
	s.PublishUserQuizAPI = new(quiz.PublishQuizHandler).InitFromBuilder(builder)

	s.CreateUserQuestionAPI = new(question.CreateQuestionHandler).InitFromBuilder(builder)
	s.ListUserQuestionAPI = new(question.ListQuestionHandler).InitFromBuilder(builder)
	s.DeleteUserQuestionAPI = new(question.DeleteQuestionHandler).InitFromBuilder(builder)
	s.UpdateUserQuestionAPI = new(question.UpdateQuestionHandler).InitFromBuilder(builder)

	s.CreateUserOptionAPI = new(option.CreateOptionHandler).InitFromBuilder(builder)
	s.ListUserOptionAPI = new(option.ListOptionHandler).InitFromBuilder(builder)
	s.DeleteUserOptionAPI = new(option.DeleteOptionHandler).InitFromBuilder(builder)
	s.UpdateUserOptionAPI = new(option.UpdateOptionHandler).InitFromBuilder(builder)

	s.ListPublishedQuizAPI = new(published.ListPublishedQuizHandler).InitFromBuilder(builder)
	s.OpenPublishedQuizAPI = new(published.OpenPublishedQuizHandler).InitFromBuilder(builder)
	s.SubmitPublishedQuizAPI = new(published.SubmitPublishedQuizHandler).InitFromBuilder(builder)

	s.ListUserSubmissionAPI = new(submission.ListUserSubmissionHandler).InitFromBuilder(builder)
	s.ListQuizSubmissionAPI = new(submission.ListQuizSubmissionHandler).InitFromBuilder(builder)
	s.GetSubmissionReportAPI = new(submission.GetSubmissionReportHandler).InitFromBuilder(builder)

	s.AuthAPI = new(middlewares.AuthUserHandler).InitFromFlags(builder.GetFlags(), builder.GetSharedFlags())
	return s
}

func (s *Server) Run() error {
	app := gin.Default()
	api := app.Group("/api")
	{
		apiV0 := api.Group("/v0")
		{
			apiV0Quiz := apiV0.Group("/quiz")
			{
				apiV0QuizUser := apiV0Quiz.Group("/user").Use(s.AuthAPI.Handler())
				{
					apiV0QuizUser.GET("/list", s.ListUserQuizAPI.Handler)
					apiV0QuizUser.POST("/create", s.CreateUserQuizAPI.Handler)
					apiV0QuizUser.DELETE("/delete", s.DeleteUserQuizAPI.Handler)
					apiV0QuizUser.PUT("/update", s.UpdateUserQuizAPI.Handler)
					apiV0QuizUser.PUT("/publish", s.PublishUserQuizAPI.Handler)
				}
				apiV0QuizPublished := apiV0Quiz.Group("/published").Use(s.AuthAPI.Handler())
				{
					apiV0QuizPublished.GET("/list", s.ListPublishedQuizAPI.Handler)
					apiV0QuizPublished.GET("/open/:quizid", s.OpenPublishedQuizAPI.Handler)
					apiV0QuizPublished.POST("/submit", s.SubmitPublishedQuizAPI.Handler)
				}
			}

			apiV0Question := apiV0.Group("/question")
			{
				apiV0QuestionUser := apiV0Question.Group("/user").Use(s.AuthAPI.Handler())
				{
					apiV0QuestionUser.POST("/create", s.CreateUserQuestionAPI.Handler)
					apiV0QuestionUser.GET("/list/:quizid", s.ListUserQuestionAPI.Handler)
					apiV0QuestionUser.DELETE("/delete", s.DeleteUserQuestionAPI.Handler)
					apiV0QuestionUser.PUT("/update", s.UpdateUserQuestionAPI.Handler)
				}
			}

			apiV0Option := apiV0.Group("/option")
			{
				apiV0OptionUser := apiV0Option.Group("/user").Use(s.AuthAPI.Handler())
				{
					apiV0OptionUser.POST("/create", s.CreateUserOptionAPI.Handler)
					apiV0OptionUser.GET("/list/:questionid", s.ListUserOptionAPI.Handler)
					apiV0OptionUser.DELETE("/delete", s.DeleteUserOptionAPI.Handler)
					apiV0OptionUser.PUT("/update", s.UpdateUserOptionAPI.Handler)
				}
			}

			apiV0Submission := apiV0.Group("/submission")
			{
				apiV0SubmissionUser := apiV0Submission.Group("/user").Use(s.AuthAPI.Handler())
				{
					apiV0SubmissionUser.GET("/list", s.ListUserSubmissionAPI.Handler)
					apiV0SubmissionUser.GET("/report/:submissionid", s.GetSubmissionReportAPI.Handler)
				}
				apiV0Submission.GET("/list/:quizid", s.ListQuizSubmissionAPI.Handler)
			}

		}
	}
	return app.Run(":" + s.GetFlags().Port)
}
