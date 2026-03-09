// {{cookiecutter.__copyright}}

package main

import (
	{%- if cookiecutter.rest == "True" or cookiecutter.grpc == "True"%}
	"time"
	"context"
	{%- endif %}
	"github.hpe.com/cloud/go-gadgets/x/logging"
)

{%- if cookiecutter.grpc == "True" %}
// GRPCServer is the interface to the internal gRPC service.
type GRPCServer interface {
	Run() error
	Shutdown(ctx context.Context) error
}
{%- endif %}


{%- if cookiecutter.rest == "True" or cookiecutter.grpc == "True" %}
// 30s is the time between kubernetes sending SIGTERM and SIGKILL when shutting down a pod
const shutdownTimeout = time.Second * 30
{%- endif %}

{%- if cookiecutter.rest == "True" %}
type RESTServer interface {
	Start(logging.Logger)
	Shutdown(context.Context) error
  }
{%- endif %}

// {{cookiecutter.__serviceStructName}} represents an instance of the service. It is provided by the ProvideService function and contains
// all of the necessary dependencies for the service to function.
type {{cookiecutter.__serviceStructName}} struct {
	Logger logging.Logger
	{%- if cookiecutter.grpc == "True" %}
	GRPCServer GRPCServer
	{%- endif %}
	{%- if cookiecutter.rest == "True" %}
	RESTServer RESTServer
	ShutdownTimeout time.Duration
	{%- endif %}
}

// NewService creates a new top level application struct, which is a root of the dependency tree
{%- if cookiecutter.rest == "False" and  cookiecutter.grpc == "False"  %}
func NewService(logger logging.Logger) *{{cookiecutter.__serviceStructName}}  {
	return &{{cookiecutter.__serviceStructName}} {
		Logger:          logger,
	}
}
{%- endif %}

{%- if cookiecutter.rest == "True" and  cookiecutter.grpc == "False" %}
func NewService(restserver RESTServer, logger logging.Logger) *{{cookiecutter.__serviceStructName}}  {
	return &{{cookiecutter.__serviceStructName}} {
		Logger:          logger,
		RESTServer:		 restserver,
		ShutdownTimeout: shutdownTimeout,
	}
}
{%- endif %}


{%- if cookiecutter.grpc == "True" and  cookiecutter.rest == "False"%}
func NewService(grpcserver GRPCServer, logger logging.Logger) *{{cookiecutter.__serviceStructName}}  {
	return &{{cookiecutter.__serviceStructName}} {
		Logger:          logger,
		GRPCServer:		 grpcserver,
	}
}
{%- endif %}

{%- if cookiecutter.grpc == "True" and  cookiecutter.rest == "True"%}
func NewService(grpcserver GRPCServer,restserver RESTServer ,logger logging.Logger) *{{cookiecutter.__serviceStructName}}  {
	return &{{cookiecutter.__serviceStructName}} {
		Logger:          logger,
		GRPCServer:		 grpcserver,
		RESTServer:		 restserver,
		ShutdownTimeout: shutdownTimeout,
	}
}
{%- endif %}
