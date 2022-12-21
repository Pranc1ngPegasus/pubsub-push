package client

import "context"

type PubSub interface {
	Publish(context.Context, []byte) error
}
