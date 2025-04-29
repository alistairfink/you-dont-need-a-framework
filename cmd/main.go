package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alistairfink/you-dont-need-a-framework/cmd/di"
	"github.com/alistairfink/you-dont-need-a-framework/handlers"
)

const (
	ExitOK    = 0
	ExitError = 1
)

func main() {
	// Since os.Exit can't handle defer we create a separate main
	os.Exit(realMain(os.Args))
}

func realMain(args []string) int {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	env := di.GetEnv()

	container, err := di.NewContainer(env)
	if err != nil {
		logger.Error("Failed to initialize di", slog.Any("error", err))
		return ExitError
	}

	httpServer, err := container.HttpServer()
	if err != nil {
		logger.Error("Failed to setup http server", slog.Any("error", err))
		return ExitError
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func(httpServer *handlers.HttpServer) {
		httpServer.ListenAndServe()
	}(httpServer)

	sig := <-signalCh
	logger.Info(fmt.Sprintf("Received signal: %v", sig))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("error", err))
		return ExitError
	}

	logger.Info("Server shutdown gracefully")
	return ExitOK
}
