//go:build wireinject
// +build wireinject

// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"context"

	"github.com/google/wire"
	"github.hpe.com/cloud/go-gadgets/x/logging"
	"github.hpe.com/cloud/go-service-template/internal/adapters"
	"github.hpe.com/cloud/go-gadgets/x/config"
	"github.hpe.com/cloud/go-gadgets/x/tracing"
	"github.hpe.com/cloud/go-service-template/internal/adapters/rest"
	"github.hpe.com/cloud/go-service-template/internal/drivers"
	"github.hpe.com/cloud/go-gadgets/x/restserver/defaulthandlers"
	"github.hpe.com/cloud/go-gadgets/x/restserver"
	"go.opentelemetry.io/otel/propagation"
	"github.hpe.com/cloud/go-gadgets/x/grpcserver"

)

type provideServiceType func(context.Context, context.CancelFunc) (*Example-Service , error)

// ProvideService returns a running instance of the service struct Example-Service.
func ProvideService(ctx context.Context, stop context.CancelFunc) (*Example-Service , error) {
	wire.Build(
		adapters.ProvideArgs,
		adapters.ProvideFlags,
		adapters.ProvideConfig,
		wire.Bind(new(config.Config), new(*config.KoanfConfig)),
		adapters.ProvideLogger,
		defaulthandlers.NewNotFoundHandler,
		defaulthandlers.NewMethodNotAllowedHandler,
		tracing.NewRandomFallbackGenerator,
		wire.Bind(new(tracing.FallbackGenerator), new(*tracing.RandomFallbackGenerator)),
		tracing.NewB3Propagator,
		wire.Bind(new(propagation.TextMapPropagator), new(*tracing.B3Propagator)),
		rest.NewReadinessCheck,
		rest.ProvideChiRouter,
		drivers.ProvideRESTServer,
		wire.Bind(new(RESTServer), new(*restserver.Server)),
		adapters.NewGRPCServices,
		adapters.ProvideGRPCOptions,
		adapters.NewGRPCServer,
		wire.Bind(new(GRPCServer), new(*grpcserver.Server)),
		NewService,
	)
	return &Example-Service {}, nil
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
