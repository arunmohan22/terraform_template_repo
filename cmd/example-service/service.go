// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

package main

import (
	"time"
	"context"
	"github.hpe.com/cloud/go-gadgets/x/logging"
)
// GRPCServer is the interface to the internal gRPC service.
type GRPCServer interface {
	Run() error
	Shutdown(ctx context.Context) error
}
// 30s is the time between kubernetes sending SIGTERM and SIGKILL when shutting down a pod
const shutdownTimeout = time.Second * 30
type RESTServer interface {
	Start(logging.Logger)
	Shutdown(context.Context) error
  }

// Example-Service represents an instance of the service. It is provided by the ProvideService function and contains
// all of the necessary dependencies for the service to function.
type Example-Service struct {
	Logger logging.Logger
	GRPCServer GRPCServer
	RESTServer RESTServer
	ShutdownTimeout time.Duration
}

// NewService creates a new top level application struct, which is a root of the dependency tree
func NewService(grpcserver GRPCServer,restserver RESTServer ,logger logging.Logger) *Example-Service  {
	return &Example-Service {
		Logger:          logger,
		GRPCServer:		 grpcserver,
		RESTServer:		 restserver,
		ShutdownTimeout: shutdownTimeout,
	}
}
