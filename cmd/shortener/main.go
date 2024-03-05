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
	flag.StringVar(&config.Settings.ServerAddress, "a", ":8080", "host:port on which server run")
	flag.StringVar(&config.Settings.BaseURL, "b", "http://localhost:8080", "URL for of shortened URLs hosting")
	flag.Parse()

	serverAddress, specified = os.LookupEnv("SERVER_ADDRESS")
	if specified {
		config.Settings.ServerAddress = serverAddress
	}

	baseURL, specified = os.LookupEnv("BASE_URL")
	if specified {
		config.Settings.BaseURL = baseURL
	}
}

func main() {
	parseFlags()
	cs := storage.CringeStorage{}
	newServer := server.NewServer(&cs)
	fmt.Printf("Running server on [%s]\n", config.Settings.ServerAddress)
	http.ListenAndServe(config.Settings.ServerAddress, newServer)
}
