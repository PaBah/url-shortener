package main

import (
	"context"
	"net/http"

	"github.com/PaBah/url-shortener.git/cmd/shortener/server"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/storage"
)

func Example() {
	options := &config.Options{}

	var store storage.Repository
	dbStore, _ := storage.NewDBStorage(context.Background(), options.DatabaseDSN)
	store = &dbStore
	defer dbStore.Close()

	newServer := server.NewRouter(options, &store)

	_ = http.ListenAndServe(options.ServerAddress, newServer)
}
