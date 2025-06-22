package main

import (
	"context"
	"gitlab.domsnail.ru/templates/go-clean-template/internal/app"
	"log/slog"
	"os/signal"
	"syscall"
)

func main() {
	// stop signal handling
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()

	// running application
	go func() {
		err := app.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()

	slog.Info("application started.")
	<-ctx.Done()
}
