package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"withoutforget/cider/internal/config"
	"withoutforget/cider/internal/server"
)

func main() {

	config_path := os.Getenv("CONFIG_PATH")

	cfg := config.GetConfig(config_path)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM,
		syscall.SIGABRT,
		syscall.SIGINT)
	defer cancel()

	srv := server.NewServer(ctx, cfg)

	go srv.Run()

	<-ctx.Done()
}
