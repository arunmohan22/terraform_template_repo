// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

package main

import (
	"context"
	"fmt"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
	"github.hpe.com/cloud/go-gadgets/x/logging"
	"github.hpe.com/cloud/go-gadgets/x/logging/assertlogging"
)

type runTestSuite struct {
	testSuite.Suite
	ctx                  context.Context
	stop                 context.CancelFunc
	logger               *assertlogging.Logger
	service              *Example-Service
}

// TestRun runs the runTestSuite
func TestRun(t *testing.T) {
	testSuite.Run(t, new(runTestSuite))
}

func provideServiceWith(service *Example-Service , err error, funcCalled *bool) provideServiceType {
	return func(ctx context.Context, stop context.CancelFunc) (*Example-Service , error) {
		*funcCalled = true
		return service , err
	}
}

func provideLoggerWith(logger logging.Logger, err error, funcCalled *bool) provideLoggerType {
	return func() (logging.Logger, error) {
		*funcCalled = true
		return logger, err
	}
}
type GRPCServerMock struct{}

func (r *GRPCServerMock) Run() error{
	return nil}

func (r *GRPCServerMock) Shutdown( ctx context.Context) error {
	return nil
}
type RESTServerMock struct{}

func (r *RESTServerMock) Start( logger logging.Logger){}

func (r *RESTServerMock) Shutdown( ctx context.Context) error {
	return nil
}

func (ts *runTestSuite) BeforeTest(_, _ string) {
	ctx, stop := context.WithCancel(context.Background())
	ts.ctx = ctx
	ts.stop = stop
	ts.logger = assertlogging.NewLogger(ts.T())
	ts.service = &Example-Service {Logger: ts.logger}
	restserver := &RESTServerMock{}
	ts.service.RESTServer = restserver
	ts.service.ShutdownTimeout = shutdownTimeout
	grpcserver := &GRPCServerMock{}
	ts.service.GRPCServer = grpcserver
}

func (ts *runTestSuite) TestHappyPath() {
	// Given
	ts.logger.ExpectInfo("Service has started")
	ts.logger.ExpectInfo("Service is stopping")
	ts.logger.ExpectInfo("gRPC server has completed graceful shutdown")
	ts.logger.ExpectInfo("REST server has completed graceful shutdown")

	provideServiceCalled := false
	provideService := provideServiceWith(ts.service, nil, &provideServiceCalled)
	provideLoggerCalled := false
	provideLogger := provideLoggerWith(ts.logger, nil, &provideLoggerCalled)
	ts.stop()

	// When
	code := Run(ts.ctx, ts.stop, provideService, provideLogger)

	// Then
	ts.Equal(ExitCodeOk, code)
	ts.True(provideServiceCalled)
}

func (ts *runTestSuite) TestProvideServiceError() {
	// Given
	err := fmt.Errorf("expected error")
	ts.logger.ExpectError(fmt.Errorf("starting service: %v", err).Error())
	ts.service.Logger = nil
	provideServiceCalled := false
	provideService := provideServiceWith(ts.service, err, &provideServiceCalled)
	provideLoggerCalled := false
	provideLogger := provideLoggerWith(ts.logger, nil, &provideLoggerCalled)
	ts.stop()

	// When
	code := Run(ts.ctx, ts.stop, provideService, provideLogger)

	// Then
	ts.Equal(ExitCodeError, code)
	ts.True(provideLoggerCalled)
	ts.True(provideLoggerCalled)
}

func (ts *runTestSuite) TestProvideLoggerError() {
	// Given
	serviceErr := fmt.Errorf("expected service error")
	loggerErr := fmt.Errorf("expected logger error")
	ts.service.Logger = nil
	provideServiceCalled := false
	provideService := provideServiceWith(ts.service, serviceErr, &provideServiceCalled)
	provideLoggerCalled := false
	provideLogger := provideLoggerWith(ts.logger, loggerErr, &provideLoggerCalled)
	ts.stop()

	// When and Then
	ts.PanicsWithError(
		fmt.Errorf(
			"error providing logger (%v) occurred after error providing service occurred (%v) ",
			loggerErr,
			serviceErr,
		).Error(),
		func() { Run(ts.ctx, ts.stop, provideService, provideLogger) },
	)
	ts.True(provideServiceCalled)
	ts.True(provideLoggerCalled)
}
