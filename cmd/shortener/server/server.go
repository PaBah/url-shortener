package server

import (
	"fmt"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

type Server struct {
	options *config.Options
	storage *storage.Repository
}

func (s Server) getShortURLHandle(res http.ResponseWriter, req *http.Request) {
	shortID := chi.URLParam(req, "id")
	responseMessage, _ := (*s.storage).FindByID(shortID)
	http.Redirect(res, req, responseMessage, http.StatusTemporaryRedirect)
}

func (s Server) postURLHandle(res http.ResponseWriter, req *http.Request) {
	var responseMessage []byte

	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		responseMessage = []byte("Invalid body")
	}

	shortURL := (*s.storage).Store(string(body))
	shortenedURL := fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL)
	res.Header().Set("Content-Type", "")
	res.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))
	res.WriteHeader(http.StatusCreated)
	responseMessage = []byte(shortenedURL)
	_, err = res.Write(responseMessage)
	if err != nil {
		fmt.Printf("Can not send response from postURLHandle: %s", err)
	}
}

func NewRouter(options *config.Options, storage *storage.Repository) *chi.Mux {
	r := chi.NewRouter()

	s := Server{
		options: options,
		storage: storage,
	}

	r.Post("/", s.postURLHandle)
	r.Get("/{id}", s.getShortURLHandle)
	r.MethodNotAllowed(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusBadRequest)
	})
	return r
}
