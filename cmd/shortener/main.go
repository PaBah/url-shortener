package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/PaBah/url-shortener.git/cmd/shortener/server"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/logger"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"go.uber.org/zap"
)

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	options := &config.Options{}
	ParseFlags(options)

	if err := logger.Initialize(options.LogsLevel); err != nil {
		fmt.Printf("Logger can not be initialized %s", err)
		return
	}

	var store storage.Repository
	dbStore, err := storage.NewDBStorage(context.Background(), options.DatabaseDSN)
	if err != nil {
		logger.Log().Error("Database error with start", zap.Error(err))
		inFileStore := storage.NewInFileStorage(options.FileStoragePath)
		store = &inFileStore

		defer inFileStore.Close()
	} else {
		store = &dbStore
		defer dbStore.Close()
	}

	newServer := server.NewRouter(options, &store)

	logger.Log().Info("Start server on", zap.String("address", options.ServerAddress))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		err := http.ListenAndServe(options.ServerAddress, newServer)

		if err != nil {
			logger.Log().Error("Server crashed with error: ", zap.Error(err))
		}
	}()

	<-ctx.Done()
}
