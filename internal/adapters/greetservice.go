// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

package adapters

import (
	"context"
	"github.hpe.com/cloud/go-gadgets/x/grpcserver/interceptors"
	"fmt"
	"github.hpe.com/cloud/go-gadgets/x/logging"
	pb "github.hpe.com/cloud/storage-proto/go/example/v1"
	"google.golang.org/grpc"
)

// GreetService implements the Greet ProtoBuf handlers, it provides the register function to be used with go-gadets grpc server.
type GreetService struct {
	pb.UnimplementedGreetServiceServer
	logger logging.Logger
}

// NewGreetService creates a new instance of the Greet service
func NewGreetService(logger logging.Logger) *GreetService {
	return &GreetService{logger: logger}
}

// Register binds the handlers of the GreetService with the supplied gRPC server
func (service *GreetService) Register(server *grpc.Server) {
	pb.RegisterGreetServiceServer(server, service)
}

// Greet is the handler of the greet person operatiion
func (service *GreetService) Greet(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	ctxlogger, err := interceptors.LoggerFromContext(ctx)
	if err != nil{
		return nil, fmt.Errorf("logger not provided")
	}
	ctxlogger.WithField("name", request.Name).Info("Received Greet request")
	message := fmt.Sprintf("Hello, %s!", request.Name)
	return &pb.GreetResponse{Message: message}, nil
}


