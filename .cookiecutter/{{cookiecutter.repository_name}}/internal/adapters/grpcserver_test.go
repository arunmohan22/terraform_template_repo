// {{cookiecutter.__copyright}}

package adapters

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	configMocks "github.hpe.com/cloud/go-gadgets/x/config/mocks"
	"github.hpe.com/cloud/go-gadgets/x/logging/assertlogging"
	"google.golang.org/grpc"
	"testing"
)

const (
	gRPCPortValue      = 5000
	gRPCWrongPortValue = 80
)

var gRPCPortKey = EnvVarToConfigKeys["APP_GRPC_PORT"]

func TestNewGRPCServer_HappyPath(t *testing.T) {
	logger := assertlogging.NewLogger(t)
	conf :=	configMocks.NewConfig(t)
	service,_ := NewGRPCServices(logger)
	options := make([]grpc.ServerOption, 0)

	conf.On("Get", gRPCPortKey).Return(gRPCPortValue, nil)

	_, serverErr := NewGRPCServer(logger, conf, service, options)
	require.NoError(t, serverErr)
}

func TestNewGRPCServer_ConfigError(t *testing.T) {
	logger := assertlogging.NewLogger(t)
	conf :=	configMocks.NewConfig(t)
	service,_ := NewGRPCServices(logger)
	options := make([]grpc.ServerOption, 0)

	portErr := errors.New("error getting gRPC port")
	conf.On("Get", gRPCPortKey).Return(nil, portErr)
	expected := "config was not found with key: grpc.port"

	_, serverErr := NewGRPCServer(logger, conf, service, options)
	assert.EqualError(t, serverErr, expected)

	logger.ExpectError("Error getting gRPC port").WithError(assertlogging.EqualError(expected))
}

func TestNewGRPCServer_WrongPort(t *testing.T) {
	logger := assertlogging.NewLogger(t)
	conf :=	configMocks.NewConfig(t)
	service ,_:= NewGRPCServices(logger)
	options := make([]grpc.ServerOption, 0)

	conf.On("Get", gRPCPortKey).Return(80, nil)
	expected := "invalid parameter port: cannot be less than 1024"

	_, serverErr := NewGRPCServer(logger, conf, service, options)
	assert.EqualError(t, serverErr, expected)

	logger.ExpectError("Error creating a new gRPC server").WithError(assertlogging.EqualError(expected))
}
