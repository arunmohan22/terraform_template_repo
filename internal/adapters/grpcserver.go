// (C) Copyright 2023 Hewlett Packard Enterprise Development LP
package adapters

import (
	"github.hpe.com/cloud/go-gadgets/x/config"
	"github.hpe.com/cloud/go-gadgets/x/grpcserver"
	"github.hpe.com/cloud/go-gadgets/x/logging"
	"google.golang.org/grpc"
)

// NewGRPCServer creates a new internal gRPC server. To do this, it fetches the gRPC port from the config using the key 'grpc.port'
func NewGRPCServer(
	logger logging.Logger,
	conf config.Config,
	services []grpcserver.Service,
	options []grpc.ServerOption,
) (*grpcserver.Server,error) {
	port, portErr := config.Get[int](conf, EnvVarToConfigKeys[EnvVarGrpcPort])
	if portErr != nil {
		logger.WithError(portErr).Error("Error getting gRPC port")
		return nil, portErr
	}

	server, serverErr := grpcserver.NewServer(uint16(port), services, options...)
	if serverErr != nil {
		logger.WithError(serverErr).Error("Error creating a new gRPC server")
		return nil, serverErr
	}

	return server, nil
}
