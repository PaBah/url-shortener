package main

import (
	"flag"
	"github.com/PaBah/url-shortener.git/internal/config"
	"os"
)

func ParseFlags(options *config.Options) {
	var specified bool
	var serverAddress, baseURL string
	flag.StringVar(&options.ServerAddress, "a", ":8080", "host:port on which server run")
	flag.StringVar(&options.BaseURL, "b", "http://localhost:8080", "URL for of shortened URLs hosting")
	flag.Parse()

	serverAddress, specified = os.LookupEnv("SERVER_ADDRESS")
	if specified {
		options.ServerAddress = serverAddress
	}

	baseURL, specified = os.LookupEnv("BASE_URL")
	if specified {
		options.BaseURL = baseURL
	}
}
