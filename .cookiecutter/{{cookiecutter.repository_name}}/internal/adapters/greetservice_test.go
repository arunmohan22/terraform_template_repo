// {{cookiecutter.__copyright}}

package adapters

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.hpe.com/cloud/go-gadgets/x/grpcserver/interceptors"
	"github.hpe.com/cloud/go-gadgets/x/logging/assertlogging"
	pb "github.hpe.com/cloud/storage-proto/go/example/v1"
	"testing"
)

func Test_Service_Greet(t *testing.T) {
	req := &pb.GreetRequest{Name: "Sam"}
	ctx := context.Background()
	logger := assertlogging.NewLogger(t)
	ctxWithLogger := interceptors.ContextWithLogger(ctx, logger)

	service := NewGreetService(logger)
	response, err := service.Greet(ctxWithLogger, req)
	require.NoError(t, err)

	assert.Equal(t, "Hello, Sam!", response.Message)

	logger.ExpectInfo("Received Greet request").WithField("name", assertlogging.Equal("Sam"))
}
