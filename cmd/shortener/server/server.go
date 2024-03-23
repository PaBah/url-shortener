package server

import (
	"fmt"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/logger"
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
	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}

	shortURL := (*s.storage).Store(string(body))
	shortenedURL := fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL)
	res.Header().Set("Content-Type", "")
	res.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))
	res.WriteHeader(http.StatusCreated)
	_, err = res.Write([]byte(shortenedURL))
	if err != nil {
		fmt.Printf("Can not send response from postURLHandle: %s", err)
	}
}

// NewRouter Creates router
func NewRouter(options *config.Options, storage *storage.Repository) *chi.Mux {
	r := chi.NewRouter()

	s := Server{
		options: options,
		storage: storage,
	}

	r.Post("/", logger.RequestLogger(s.postURLHandle))
	r.Get("/{id}", logger.RequestLogger(s.getShortURLHandle))
	r.MethodNotAllowed(
		logger.RequestLogger(
			func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusBadRequest)
			},
		),
	)
	return r
}
