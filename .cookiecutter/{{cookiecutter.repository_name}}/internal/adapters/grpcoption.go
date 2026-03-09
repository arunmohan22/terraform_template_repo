// {{cookiecutter.__copyright}}

package adapters

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.hpe.com/cloud/go-gadgets/x/grpcserver/interceptors"
	"github.hpe.com/cloud/go-gadgets/x/logging"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
)

// Provide standard options, such as logging, tracing, panicing interceptors.
// These standard options can be added to or replaced where needed
func ProvideGRPCOptions(logger logging.Logger, propagator propagation.TextMapPropagator,) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptors.PanicUnaryServerInterceptor(logger),
			interceptors.LoggerUnaryServerInterceptor(logger),
			interceptors.TracingUnaryServerInterceptor(logger, propagator),
		),
	),
	}
}