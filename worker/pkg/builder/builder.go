package builder

import (
	shared "github.com/verasthiago/quiz/shared/flags"
	"github.com/verasthiago/quiz/shared/repository"
)

type Builder interface {
	GetRepository() repository.Repository
	GetSharedFlags() *shared.SharedFlags
	InitBuilder() Builder
}
