// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.hpe.com/cloud/go-gadgets/x/headers"
	"github.hpe.com/cloud/go-gadgets/x/logging"
	"github.hpe.com/cloud/go-gadgets/x/restserver/defaulthandlers"
	"github.hpe.com/cloud/go-gadgets/x/restserver/middleware"
	"go.opentelemetry.io/otel/propagation"
)

const readinessPath = "/readyz"

func ProvideChiRouter(
	logger logging.Logger,
	notFoundHandler *defaulthandlers.NotFoundHandler,
	methodNotAllowedHandler *defaulthandlers.MethodNotAllowedHandler,
	propagator propagation.TextMapPropagator,
	readinessCheck *ReadinessCheck,
) chi.Router {
	router := chi.NewRouter()

	// global middleware
	router.Use(middleware.PanicMiddleware(logger))
	router.Use(middleware.LoggerMiddleware(logger, middleware.WithRequestFields()))
	router.Use(middleware.LogHeaderMiddleware(logger, middleware.OptionalLogHeaders(headers.XDSCCTestname, headers.UserAgent)))
	router.Use(middleware.TracingMiddleware(logger, propagator))

	// default handlers
	router.NotFound(notFoundHandler.ServeHTTP)
	router.MethodNotAllowed(methodNotAllowedHandler.ServeHTTP)

	router.Method(http.MethodGet, readinessPath, readinessCheck)

	router.Group(func(router chi.Router) {
		// This is a basic existence check handler, refer to restserver/defaulthandlers for examples of handlers that
		// use net/http.Handler interface
		router.MethodFunc(http.MethodGet, "/ping", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("pong"))
		})
	})

	return router
}
