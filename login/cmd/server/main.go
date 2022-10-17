package cmd

import (
	"github.com/verasthiago/quiz/login/pkg"
	"github.com/verasthiago/quiz/login/pkg/builder"
)

func Execute() {
	builder := new(builder.ServerBuilder).InitBuilder()
	server := new(pkg.Server).InitFromBuilder(builder)

	if err := server.Run(); err != nil {
		panic(err)
	}
}
