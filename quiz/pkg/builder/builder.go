package builder

import (
	shared "github.com/verasthiago/quiz/shared/flags"
	"github.com/verasthiago/quiz/shared/repository"
	"github.com/verasthiago/quiz/shared/task"
	"go.uber.org/zap"
)

type Builder interface {
	GetRepository() repository.Repository
	GetFlags() *Flags
	GetLog() *zap.Logger
	GetTask() task.Task
	GetSharedFlags() *shared.SharedFlags
	InitBuilder() Builder
}
