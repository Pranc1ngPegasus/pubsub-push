//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"net/http"

	"github.com/Pranc1ngPegasus/pubsub-push/adapter/handler"
	"github.com/Pranc1ngPegasus/pubsub-push/adapter/server"
	domainlogger "github.com/Pranc1ngPegasus/pubsub-push/domain/logger"
	"github.com/Pranc1ngPegasus/pubsub-push/infra/configuration"
	"github.com/Pranc1ngPegasus/pubsub-push/infra/logger"
	"github.com/google/wire"
)

type app struct {
	logger domainlogger.Logger
	server *http.Server
}

func initialize() (*app, error) {
	wire.Build(
		context.Background,

		configuration.NewConfigurationSet,
		logger.NewLoggerSet,
		handler.NewHandlerSet,
		server.NewServer,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
