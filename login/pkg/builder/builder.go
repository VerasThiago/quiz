package builder

import (
	shared "github.com/verasthiago/quiz/shared/flags"
	"github.com/verasthiago/quiz/shared/repository"
	"go.uber.org/zap"
)

type Builder interface {
	GetRepository() repository.Repository
	GetFlags() *Flags
	GetLog() *zap.Logger
	GetSharedFlags() *shared.SharedFlags
	InitBuilder() Builder
}
