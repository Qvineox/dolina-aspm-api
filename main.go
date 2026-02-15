package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"gitlab.domsnail.ru/dolina/dolina-aspm-api/internal/app"
)

func main() {
	// stop signal handling
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()

	// running applications
	go func() {
		err := app.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()

	slog.Info("applications started.")
	<-ctx.Done()
}
