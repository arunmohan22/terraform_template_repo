// {{cookiecutter.__copyright}}

//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"

	"github.hpe.com/cloud/go-gadgets/x/logging"

	//"{{cookiecutter.__repository_url}}/internal/adapters/rest"
)

func makeApplication(ctx context.Context) (*application, error) {
	wire.Build(
		newLogger,
		wire.Bind(new(logging.Logger), new(*logging.ZapJSONLogger)),

		routes.NewRouter,

		newServer,

		newApplication,
	)

	return &application{}, nil
}
