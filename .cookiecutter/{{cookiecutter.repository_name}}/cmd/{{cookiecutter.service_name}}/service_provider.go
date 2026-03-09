//go:build wireinject
// +build wireinject

// {{cookiecutter.__copyright}}

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"context"

	"github.com/google/wire"
	"github.hpe.com/cloud/go-gadgets/x/logging"
	"{{cookiecutter.repoURL}}/internal/adapters"
	"github.hpe.com/cloud/go-gadgets/x/config"
	{%- if cookiecutter.rest == "True" or cookiecutter.grpc == "True" %}
	"github.hpe.com/cloud/go-gadgets/x/tracing"
	{%- endif %}
	{%- if cookiecutter.rest == "True" %}
	"{{cookiecutter.repoURL}}/internal/adapters/rest"
	"{{cookiecutter.repoURL}}/internal/drivers"
	"github.hpe.com/cloud/go-gadgets/x/restserver/defaulthandlers"
	"github.hpe.com/cloud/go-gadgets/x/restserver"
	{%- endif %}
	{%- if cookiecutter.rest == "True" or cookiecutter.grpc == "True"%}
	"go.opentelemetry.io/otel/propagation"
	{%- endif %}
	{%- if cookiecutter.grpc == "True" %}
	"github.hpe.com/cloud/go-gadgets/x/grpcserver"
	{%- endif %}

)

type provideServiceType func(context.Context, context.CancelFunc) (*{{cookiecutter.__serviceStructName}} , error)

// ProvideService returns a running instance of the service struct {{cookiecutter.__serviceStructName}}.
func ProvideService(ctx context.Context, stop context.CancelFunc) (*{{cookiecutter.__serviceStructName}} , error) {
	wire.Build(
		adapters.ProvideArgs,
		adapters.ProvideFlags,
		adapters.ProvideConfig,
		wire.Bind(new(config.Config), new(*config.KoanfConfig)),
		adapters.ProvideLogger,
		{%- if cookiecutter.rest == "True" %}
		defaulthandlers.NewNotFoundHandler,
		defaulthandlers.NewMethodNotAllowedHandler,
		{%- endif %}
		{%- if cookiecutter.grpc == "True" or cookiecutter.rest == "True" %}
		tracing.NewRandomFallbackGenerator,
		wire.Bind(new(tracing.FallbackGenerator), new(*tracing.RandomFallbackGenerator)),
		tracing.NewB3Propagator,
		wire.Bind(new(propagation.TextMapPropagator), new(*tracing.B3Propagator)),
		{%- endif %}
		{%- if cookiecutter.rest == "True" %}
		rest.NewReadinessCheck,
		rest.ProvideChiRouter,
		drivers.ProvideRESTServer,
		wire.Bind(new(RESTServer), new(*restserver.Server)),
		{%- endif %}
		{%- if cookiecutter.grpc == "True" %}
		adapters.NewGRPCServices,
		adapters.ProvideGRPCOptions,
		adapters.NewGRPCServer,
		wire.Bind(new(GRPCServer), new(*grpcserver.Server)),
		{%- endif %}
		NewService,
	)
	return &{{cookiecutter.__serviceStructName}} {}, nil
}

type provideLoggerType func() (logging.Logger, error)

// ProvideLogger provides a new ZapJsonLogger
func ProvideLogger() (logging.Logger, error) {
	wire.Build(
		adapters.ProvideArgs,
		adapters.ProvideFlags,
		adapters.ProvideConfig,
		wire.Bind(new(config.Config), new(*config.KoanfConfig)),
		adapters.ProvideLogger,
	)
	return &logging.ZapJSONLogger{}, nil
}
