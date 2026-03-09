// (C) Copyright 2023 Hewlett Packard Enterprise Development LP
package drivers

import (
	"context"
	"errors"
	"testing"

	"github.hpe.com/cloud/go-service-template/internal/adapters/rest"
	"github.hpe.com/cloud/go-service-template/internal/adapters"
	"github.com/stretchr/testify/assert"
	configmocks "github.hpe.com/cloud/go-gadgets/x/config/mocks"
	"github.hpe.com/cloud/go-gadgets/x/logging/assertlogging"
	"github.hpe.com/cloud/go-gadgets/x/restserver"
	"github.hpe.com/cloud/go-gadgets/x/restserver/defaulthandlers"
	mtracing "github.hpe.com/cloud/go-gadgets/x/tracing/mocks"
	"github.hpe.com/cloud/go-gadgets/x/tracing"
)

var restPort = adapters.EnvVarToConfigKeys[adapters.EnvVarRestPort]

func TestProvideRESTServer_HappyPath(t *testing.T) {
	logger := assertlogging.NewLogger(t)
	notFoundHandler, err := defaulthandlers.NewNotFoundHandler(logger)
	assert.NoError(t, err)
	methodNotAllowedHandler, err := defaulthandlers.NewMethodNotAllowedHandler(logger)
	assert.NoError(t, err)
	conf := &configmocks.Config{}

	fallbackGenerator := new(mtracing.FallbackGenerator)

	propagatorLogger := assertlogging.NewLogger(t)
	propagator := tracing.NewB3Propagator(fallbackGenerator, propagatorLogger)

	router := rest.ProvideChiRouter(
		logger,
		notFoundHandler,
		methodNotAllowedHandler,
		propagator,
		rest.NewReadinessCheck(),
	)

	conf.On("Get", restPort).Return(8080, nil)

	server, err := ProvideRESTServer(context.Background(), conf, logger, router)

	assert.NoError(t,err)
	assert.IsType(t, &restserver.Server{}, server)

	assert.NoError(t, server.Shutdown(context.Background()))
}

func TestProvideRESTServer_ConfigError(t *testing.T) {
	logger := assertlogging.NewLogger(t)
	notFoundHandler, err := defaulthandlers.NewNotFoundHandler(logger)
	assert.NoError(t,err)
	methodNotAllowedHandler, err := defaulthandlers.NewMethodNotAllowedHandler(logger)
	assert.NoError(t,err)
	conf := &configmocks.Config{}

	fallbackGenerator := new(mtracing.FallbackGenerator)

	propagatorLogger := assertlogging.NewLogger(t)
	propagator := tracing.NewB3Propagator(fallbackGenerator, propagatorLogger)

	router := rest.ProvideChiRouter(
		logger,
		notFoundHandler,
		methodNotAllowedHandler,
		propagator,
		rest.NewReadinessCheck(),
	)

	conf.On("Get", restPort).Return(nil, errors.New("test error"))
	expectedErr := errors.New("getting rest.port from config: config was not found with key: rest.port")
	server, err := ProvideRESTServer(context.Background(), conf, logger, router)

	assert.Error(t, err)

	assert.EqualError(t, expectedErr, err.Error())
	assert.Nil(t, server)
}

func TestProvideRESTServer_NewServerError(t *testing.T) {
	logger := assertlogging.NewLogger(t)
	notFoundHandler, err := defaulthandlers.NewNotFoundHandler(logger)
	assert.NoError(t, err)
	methodNotAllowedHandler, err := defaulthandlers.NewMethodNotAllowedHandler(logger)
	assert.NoError(t, err)
	conf := &configmocks.Config{}

	fallbackGenerator := new(mtracing.FallbackGenerator)

	propagatorLogger := assertlogging.NewLogger(t)
	propagator := tracing.NewB3Propagator(fallbackGenerator, propagatorLogger)

	router := rest.ProvideChiRouter(
		logger,
		notFoundHandler,
		methodNotAllowedHandler,
		propagator,
		rest.NewReadinessCheck(),
	)

	conf.On("Get", restPort).Return(80, nil)
	expectedErr := errors.New("provide rest server: invalid parameter port: cannot be less than 1024")

	server, err := ProvideRESTServer(context.Background(), conf, logger, router)

	assert.Error(t, err)
	assert.EqualError(t, expectedErr, err.Error())

	assert.Nil(t, server)
}
