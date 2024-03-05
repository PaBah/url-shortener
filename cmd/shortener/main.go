package main

import (
	"flag"
	"fmt"
	"github.com/PaBah/url-shortener.git/cmd/shortener/config"
	"github.com/PaBah/url-shortener.git/cmd/shortener/server"
	"github.com/PaBah/url-shortener.git/cmd/shortener/storage"
	"net/http"
)

func parseFlags() {
	flag.StringVar(&config.Settings.ServerHost, "a", ":8080", "host:port on which server run")
	flag.StringVar(&config.Settings.ShortURLHost, "b", "http://localhost:8080", "URL for of shortened URLs hosting")
	flag.Parse()
}

func main() {
	parseFlags()
	cs := storage.CringeStorage{}
	newServer := server.NewServer(&cs)
	fmt.Printf("Running server on [%s]\n", config.Settings.ServerHost)
	http.ListenAndServe(config.Settings.ServerHost, newServer)
}
