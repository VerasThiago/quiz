package builder

import (
	shared "github.com/verasthiago/quiz/shared/flags"
	"github.com/verasthiago/quiz/shared/repository"
	postgresrepository "github.com/verasthiago/quiz/shared/repository/postgresRepository"
	"github.com/verasthiago/quiz/shared/task"
	asynqtask "github.com/verasthiago/quiz/shared/task/asynctask"
	"go.uber.org/zap"
)

type ServerBuilder struct {
	*Flags
	*shared.SharedFlags
	Repository repository.Repository
	Log        *zap.Logger
	Task       task.Task
}

func (s *ServerBuilder) GetFlags() *Flags {
	return s.Flags
}

func (s *ServerBuilder) GetTask() task.Task {
	return s.Task
}

func (s *ServerBuilder) GetSharedFlags() *shared.SharedFlags {
	return s.SharedFlags
}

func (s *ServerBuilder) GetRepository() repository.Repository {
	return s.Repository
}

func (s *ServerBuilder) GetLog() *zap.Logger {
	return s.Log
}

func (s *ServerBuilder) InitBuilder() Builder {

	flags, err := new(Flags).InitFromViper()
	if err != nil {
		panic(err)
	}
	s.Flags = flags

	sharedflags, err := new(shared.SharedFlags).InitFromViper()
	if err != nil {
		panic(err)
	}

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	s.SharedFlags = sharedflags
	s.Log = log
	s.Repository = new(postgresrepository.PostgresRepository).InitFromFlags(s.SharedFlags)
	s.Task = new(asynqtask.AsyncQueue).InitFromFlags(s.SharedFlags)

	return s
}
