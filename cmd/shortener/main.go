package main

import (
	"github.com/PaBah/url-shortener.git/cmd/shortener/server"
	"github.com/PaBah/url-shortener.git/cmd/shortener/storage"
	"net/http"
)

func main() {
	cs := storage.CringeStorage{}
	newServer := server.NewServer(&cs)
	http.ListenAndServe(":8080", *newServer)
}
