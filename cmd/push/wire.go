//go:build wireinject
// +build wireinject

package main

import (
	"context"

	domainclient "github.com/Pranc1ngPegasus/pubsub-push/domain/client"
	domainlogger "github.com/Pranc1ngPegasus/pubsub-push/domain/logger"
	"github.com/Pranc1ngPegasus/pubsub-push/infra/client"
	"github.com/Pranc1ngPegasus/pubsub-push/infra/configuration"
	"github.com/Pranc1ngPegasus/pubsub-push/infra/logger"
	"github.com/google/wire"
)

type app struct {
	logger domainlogger.Logger
	client domainclient.PubSub
}

func initialize() (*app, error) {
	wire.Build(
		context.Background,

		configuration.NewConfigurationSet,
		logger.NewLoggerSet,
		client.NewPubSubSet,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
