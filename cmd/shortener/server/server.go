package server

import (
	"fmt"
	"github.com/PaBah/url-shortener.git/cmd/shortener/storage"
	"io"
	"net/http"
	"strconv"
)

type Server struct {
	Storage storage.Repository
}

func NewServer(storage storage.Repository) *Server {
	newServer := Server{storage}
	return &newServer
}

func (srv Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var responseMessage string
	if req.Method != http.MethodGet && req.Method != http.MethodPost {
		res.WriteHeader(http.StatusBadRequest)
		responseMessage = "Unsupported HTTP Method"
	}

	if req.Method == http.MethodPost {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			responseMessage = "Invalid body"
		}

		shortURL := srv.addURL(string(body))
		shortenedURL := fmt.Sprintf("http://localhost:8080/%s", shortURL)
		res.Header().Set("Content-Type", "")
		res.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))
		res.WriteHeader(http.StatusCreated)
		responseMessage = shortenedURL
	}

	if req.Method == http.MethodGet {
		responseMessage, _ = srv.findURL(req.URL.EscapedPath()[1:])
		http.Redirect(res, req, responseMessage, http.StatusTemporaryRedirect)
	}

	_, err := res.Write([]byte(responseMessage))
	if err != nil {
		panic("Can not send response!")
	}
}

func (srv Server) addURL(Value string) (shortURL string) {
	shortURL = srv.Storage.Store(Value)
	return
}

func (srv Server) findURL(shorURL string) (string, error) {
	result, _ := srv.Storage.FindByID(shorURL)
	return result, nil
}
