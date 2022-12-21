// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/Pranc1ngPegasus/pubsub-push/adapter/handler"
	"github.com/Pranc1ngPegasus/pubsub-push/adapter/server"
	logger2 "github.com/Pranc1ngPegasus/pubsub-push/domain/logger"
	"github.com/Pranc1ngPegasus/pubsub-push/infra/configuration"
	"github.com/Pranc1ngPegasus/pubsub-push/infra/logger"
	"net/http"
)

// Injectors from wire.go:

func initialize() (*app, error) {
	loggerLogger, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	contextContext := context.Background()
	configurationConfiguration, err := configuration.NewConfiguration()
	if err != nil {
		return nil, err
	}
	handlerHandler := handler.NewHandler(loggerLogger)
	httpServer := server.NewServer(contextContext, loggerLogger, configurationConfiguration, handlerHandler)
	mainApp := &app{
		logger: loggerLogger,
		server: httpServer,
	}
	return mainApp, nil
}

// wire.go:

type app struct {
	logger logger2.Logger
	server *http.Server
}
