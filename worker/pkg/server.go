package pkg

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/verasthiago/quiz/shared/task/asynctask"
	"github.com/verasthiago/quiz/worker/pkg/builder"
	"github.com/verasthiago/quiz/worker/pkg/handlers"
)

type Server struct {
	builder.Builder

	ReportCreateAPI handlers.CreateReportAPI
}

func (s *Server) InitFromBuilder(builder builder.Builder) *Server {
	s.Builder = builder

	s.ReportCreateAPI = new(handlers.CreateReportHandler).InitFromBuilder(builder)
	return s
}

func (s *Server) Run() error {
	dsn := fmt.Sprintf("%+v:%+v", s.GetSharedFlags().QueueHost, s.GetSharedFlags().QueuePort)
	redisConnection := asynq.RedisClientOpt{
		Addr: dsn,
	}

	worker := asynq.NewServer(redisConnection, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})

	mux := asynq.NewServeMux()

	mux.HandleFunc(
		asynctask.TypeReportCreate,
		s.ReportCreateAPI.Handler,
	)

	return worker.Run(mux)
}
