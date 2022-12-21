package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	domain "github.com/Pranc1ngPegasus/pubsub-push/domain/client"
	"github.com/Pranc1ngPegasus/pubsub-push/domain/configuration"
	"github.com/google/wire"
)

var _ domain.PubSub = (*PubSub)(nil)

var NewPubSubSet = wire.NewSet(
	wire.Bind(new(domain.PubSub), new(*PubSub)),
	NewPubSub,
)

type PubSub struct {
	client *pubsub.Client
	config *configuration.GCP
}

func NewPubSub(
	ctx context.Context,
	cfg configuration.Configuration,
) (*PubSub, error) {
	config := cfg.GCP()

	client, err := pubsub.NewClient(ctx, config.GCPProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize pubsub client: %w", err)
	}

	return &PubSub{
		client: client,
		config: config,
	}, nil
}

func (c *PubSub) Publish(ctx context.Context, data []byte) error {
	topic := c.client.Topic(c.config.Topic)
	defer topic.Stop()

	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	if _, err := result.Get(ctx); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
