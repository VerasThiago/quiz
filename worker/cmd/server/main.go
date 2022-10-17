package server

import (
	"github.com/verasthiago/quiz/worker/pkg"
	"github.com/verasthiago/quiz/worker/pkg/builder"
)

func Execute() {
	builder := new(builder.ServerBuilder).InitBuilder()
	server := new(pkg.Server).InitFromBuilder(builder)

	if err := server.Run(); err != nil {
		panic(err)
	}
}
