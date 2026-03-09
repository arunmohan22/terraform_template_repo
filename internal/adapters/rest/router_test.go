// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.hpe.com/cloud/go-gadgets/x/logging/assertlogging"
	"github.hpe.com/cloud/go-gadgets/x/restserver/defaulthandlers"
	"github.hpe.com/cloud/go-gadgets/x/tracing"
)

func TestProvideRouter(t *testing.T) {
	logger := assertlogging.NewLogger(t)
	notFoundHandler, err := defaulthandlers.NewNotFoundHandler(logger)
	require.NoError(t, err)
	methodNotAllowedHandler, err := defaulthandlers.NewMethodNotAllowedHandler(logger)
	require.NoError(t, err)

	fallbackGenerator, err := tracing.NewRandomFallbackGenerator()
	require.NoError(t, err)

	propagatorLogger := assertlogging.NewLogger(t)
	propagator := tracing.NewB3Propagator(fallbackGenerator, propagatorLogger)

	router := ProvideChiRouter(
		logger,
		notFoundHandler,
		methodNotAllowedHandler,
		propagator,
		NewReadinessCheck(),
	)

	assert.NotNil(t, router)
}
