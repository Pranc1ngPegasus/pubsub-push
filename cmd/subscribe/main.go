package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app, err := initialize()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err := app.server.ListenAndServe(); err != nil {
		app.logger.Error(ctx, "failed to start server", app.logger.Field("err", err))
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.server.Shutdown(ctx); err != nil {
		app.logger.Error(ctx, "failed to shutdown server", app.logger.Field("err", err))
	}
}
