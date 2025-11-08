package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"withoutforget/cider/internal/config"
	"withoutforget/cider/internal/logging"
	"withoutforget/cider/internal/server"
)

func main() {
	config_path := os.Getenv("CONFIG_PATH")

	cfg := config.GetConfig(config_path)

	logging.InitLogger(cfg.Logging)

	slog.Info("Logger initializated")
	fmt.Println(cfg.String())

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM,
		syscall.SIGABRT,
		syscall.SIGINT)

	srv := server.NewServer(ctx, cfg)

	go srv.Run()

	<-ctx.Done()
	slog.Info("Global ctx is done... Shutting down")
	cancel()
	srv.Shutdown()
	slog.Info("Server has been shut down")
}
