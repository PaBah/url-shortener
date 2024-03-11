package main

import (
	"fmt"
	"github.com/PaBah/url-shortener.git/cmd/shortener/server"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"net/http"
)

func main() {
	options := &config.Options{}
	ParseFlags(options)

	var store storage.Repository
	inFileStore := storage.InFileStorage{}
	store = &inFileStore

	newServer := server.NewRouter(options, &store)
	fmt.Printf("Running server on [%s]\n", options.ServerAddress)
	http.ListenAndServe(options.ServerAddress, newServer)
}
