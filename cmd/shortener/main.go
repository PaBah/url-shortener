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

func main() {
	options := &config.Options{}
	ParseFlags(options)

	if err := logger.Initialize(options.LogsLevel); err != nil {
		fmt.Printf("Logger can not be initialized %s", err)
		return
	}

	var store storage.Repository
	inFileStore := storage.NewInFileStorage(options.FileStoragePath)
	defer inFileStore.Close()

	store = &inFileStore
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
