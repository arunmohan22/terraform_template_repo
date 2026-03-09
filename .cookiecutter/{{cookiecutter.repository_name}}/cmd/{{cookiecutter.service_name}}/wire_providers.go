// {{cookiecutter.__copyright}}

package main

import (
	"github.com/go-chi/chi/v5"

	"github.hpe.com/cloud/go-gadgets/x/logging"
	"github.hpe.com/cloud/go-gadgets/x/restserver"
	"github.hpe.com/cloud/go-gadgets/x/restserver/chiserver"
)

func newApplication(
	logger logging.Logger,
	server *restserver.Server,
) *application {
	return &application{
		logger: logger,
		server: server,
	}
}

func newLogger() (*logging.ZapJSONLogger, error) {
	// TODO - Add support for loading config values.
	return logging.NewZapJSONLogger("info")
}

func newServer(router chi.Router) (*restserver.Server, error) {
	// TODO - Add support for loading config values.
	return chiserver.NewServerFromRouter(router, 5000)
}
