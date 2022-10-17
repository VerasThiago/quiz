package builder

import (
	shared "github.com/verasthiago/quiz/shared/flags"
	"github.com/verasthiago/quiz/shared/repository"
	postgresrepository "github.com/verasthiago/quiz/shared/repository/postgresRepository"
)

type ServerBuilder struct {
	*shared.SharedFlags
	Repository repository.Repository
}

func (s *ServerBuilder) GetSharedFlags() *shared.SharedFlags {
	return s.SharedFlags
}

func (s *ServerBuilder) GetRepository() repository.Repository {
	return s.Repository
}

func (s *ServerBuilder) InitBuilder() Builder {

	sharedflags, err := new(shared.SharedFlags).InitFromViper()
	if err != nil {
		panic(err)
	}

	s.SharedFlags = sharedflags
	s.Repository = new(postgresrepository.PostgresRepository).InitFromFlags(s.SharedFlags)

	return s
}
