// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

const (
	ExitCodeOk    = 0
	ExitCodeError = 1
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	code := Run(ctx, stop, ProvideService, ProvideLogger)
	os.Exit(code)
}

// Run is the unit testable entrypoint that starts and stops the service
func Run(ctx context.Context, stop context.CancelFunc, provideService provideServiceType, provideLogger provideLoggerType) int {
	service, serviceErr := provideService(ctx, stop)
	if serviceErr != nil {
		logger, loggerErr := provideLogger()
		if loggerErr != nil {
			panic(fmt.Errorf("error providing logger (%v) occurred after error providing service occurred (%v) ",
				loggerErr,
				serviceErr,
			))
		}
		logger.Error(fmt.Errorf("starting service: %v", serviceErr).Error())
		return ExitCodeError
	}
	//nolint: errcheck // This is sloppy but will be fixed in next PR when the mains are split out
	go service.GRPCServer.Run()
	go service.RESTServer.Start(service.Logger)
	service.Logger.Info("Service has started")

	<-ctx.Done()

	stop()

	service.Logger.Info("Service is stopping")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := service.GRPCServer.Shutdown(ctx); err != nil {
		service.Logger.WithError(err).Error("Error stopping the gRPC Server")
	}

	service.Logger.Info("gRPC server has completed graceful shutdown")
	ctx, cleanup := context.WithTimeout(context.Background(), service.ShutdownTimeout)
	defer cleanup()

	if err := service.RESTServer.Shutdown(ctx); err != nil {
		service.Logger.WithError(err).Error("error shutting down rest server")
	}

	service.Logger.Info("REST server has completed graceful shutdown")

	return ExitCodeOk
}
