package main

import (
	"fmt"
	"net/http"

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

	logger.Log.Info("Start server on", zap.String("address", options.ServerAddress))

	err := http.ListenAndServe(options.ServerAddress, newServer)

	if err != nil {
		logger.Log.Error("Server crashed with error: ", zap.Error(err))
	}
}
