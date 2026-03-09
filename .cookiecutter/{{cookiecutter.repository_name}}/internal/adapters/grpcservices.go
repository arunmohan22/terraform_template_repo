// {{cookiecutter.__copyright}}

package adapters

import (
	"fmt"

	"github.hpe.com/cloud/go-gadgets/x/grpcserver"
	"github.hpe.com/cloud/go-gadgets/x/logging"
)

// NewGRPCServices provides specific services to be handled by the gRPC server
func NewGRPCServices(logger logging.Logger) ([]grpcserver.Service ,error) {
	if logger == nil{
		return nil, fmt.Errorf("logger not provided")
	}
	return []grpcserver.Service{NewGreetService(logger)} , nil
}
