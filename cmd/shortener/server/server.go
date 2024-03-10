package server

import (
	"fmt"
	"github.com/PaBah/url-shortener.git/cmd/shortener/config"
	"github.com/PaBah/url-shortener.git/cmd/shortener/storage"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

type Server struct{}

func (s Server) getShortURLHandle(store *storage.Repository) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		shortID := chi.URLParam(req, "id")
		responseMessage, _ := (*store).FindByID(shortID)
		http.Redirect(res, req, responseMessage, http.StatusTemporaryRedirect)
	}
}

func (s Server) postURLHandle(store *storage.Repository) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var responseMessage string

		body, err := io.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			responseMessage = "Invalid body"
		}

		shortURL := (*store).Store(string(body))
		shortenedURL := fmt.Sprintf("%s/%s", config.Options.BaseURL, shortURL)
		res.Header().Set("Content-Type", "")
		res.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))
		res.WriteHeader(http.StatusCreated)
		responseMessage = shortenedURL
		res.Write([]byte(responseMessage))
	}
}

func NewServer(storage storage.Repository) *chi.Mux {
	r := chi.NewRouter()
	s := Server{}
	r.Post("/", s.postURLHandle(&storage))
	r.Get("/{id}", s.getShortURLHandle(&storage))
	r.MethodNotAllowed(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusBadRequest)
	})
	return r
}
