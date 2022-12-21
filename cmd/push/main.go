package main

import (
	"context"
	"os"
)

var (
	message = []byte("Merry Christmas!")
)

func main() {
	app, err := initialize()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err := app.client.Publish(ctx, message); err != nil {
		app.logger.Error(ctx, "failed to publish message", app.logger.Field("err", err))
		os.Exit(1)
	}

	app.logger.Info(ctx, "published")
}
