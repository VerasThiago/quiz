package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/verasthiago/quiz/login/pkg/builder"
	"github.com/verasthiago/quiz/login/pkg/handlers"
	"github.com/verasthiago/quiz/login/pkg/middlewares"
)

type Server struct {
	builder.Builder

	LoginAPI  handlers.LoginUserAPI
	CreateAPI handlers.CreateUserAPI
	DeleteAPI handlers.DeleteUserAPI
	UpdateAPI handlers.UpdateUserAPI

	AdminAPI middlewares.AuthUserAPI
}

func (s *Server) InitFromBuilder(builder builder.Builder) *Server {
	s.Builder = builder
	s.LoginAPI = new(handlers.LoginUserHandler).InitFromBuilder(builder)
	s.CreateAPI = new(handlers.CreateUserHandler).InitFromBuilder(builder)
	s.DeleteAPI = new(handlers.DeleteUserHandler).InitFromBuilder(builder)
	s.UpdateAPI = new(handlers.UpdateUserHandler).InitFromBuilder(builder)

	s.AdminAPI = new(middlewares.AuthUserHandler).InitFromFlags(builder.GetFlags(), builder.GetSharedFlags())
	return s
}

func (s *Server) Run() error {

	app := gin.Default()
	api := app.Group("/login")
	{
		apiV0 := api.Group("/v0")
		{
			apiV0User := apiV0.Group("/user")
			{
				apiV0User.POST("register", s.CreateAPI.Handler)
				apiV0User.POST("login", s.LoginAPI.Handler)
			}
			apiV0Admin := apiV0.Group("/admin").Use(s.AdminAPI.Handler())
			{
				apiV0Admin.DELETE("delete", s.DeleteAPI.Handler)
				apiV0Admin.PUT("update", s.UpdateAPI.Handler)
			}
		}
	}
	return app.Run(":" + s.GetFlags().Port)
}
