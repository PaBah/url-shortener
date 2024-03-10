package main

import (
	"flag"
	"fmt"
	"github.com/PaBah/url-shortener.git/cmd/shortener/config"
	"github.com/PaBah/url-shortener.git/cmd/shortener/server"
	"github.com/PaBah/url-shortener.git/cmd/shortener/storage"
	"net/http"
	"os"
)

func parseFlags() {
	var specified bool
	var serverAddress, baseURL string
	flag.StringVar(&config.Options.ServerAddress, "a", ":8080", "host:port on which server run")
	flag.StringVar(&config.Options.BaseURL, "b", "http://localhost:8080", "URL for of shortened URLs hosting")
	flag.Parse()

	serverAddress, specified = os.LookupEnv("SERVER_ADDRESS")
	if specified {
		config.Options.ServerAddress = serverAddress
	}

	baseURL, specified = os.LookupEnv("BASE_URL")
	if specified {
		config.Options.BaseURL = baseURL
	}
}

func main() {
	parseFlags()
	cs := &storage.InFileStorage{}
	newServer := server.NewServer(cs)
	fmt.Printf("Running server on [%s]\n", config.Options.ServerAddress)
	http.ListenAndServe(config.Options.ServerAddress, newServer)
}
