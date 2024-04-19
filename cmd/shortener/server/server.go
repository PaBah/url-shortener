package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/dto"
	"github.com/PaBah/url-shortener.git/internal/logger"
	"github.com/PaBah/url-shortener.git/internal/middlewares"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

type Server struct {
	options *config.Options
	storage storage.Repository
}

func (s Server) getShortURLHandle(res http.ResponseWriter, req *http.Request) {
	shortID := chi.URLParam(req, "id")

	responseMessage, _ := s.storage.FindByID(req.Context(), shortID)
	res.Header().Set("Location", responseMessage)
	http.Redirect(res, req, responseMessage, http.StatusTemporaryRedirect)
}

func (s Server) postURLHandle(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	shortURL, err := s.storage.Store(req.Context(), string(body))
	shortenedURL := fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))

	if errors.Is(err, storage.ErrConflict) {
		res.WriteHeader(http.StatusConflict)
	} else {
		res.WriteHeader(http.StatusCreated)
	}

	_, err = res.Write([]byte(shortenedURL))
	if err != nil {
		logger.Log().Error("Can not send response from postURLHandle:", zap.Error(err))
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) apiShortenHandle(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	requestData := &dto.ShortenRequest{}
	err = json.Unmarshal(body, requestData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	shortURL, err := s.storage.Store(req.Context(), requestData.URL)

	if errors.Is(err, storage.ErrConflict) {
		res.WriteHeader(http.StatusConflict)
	} else {
		res.WriteHeader(http.StatusCreated)
	}

	responseData := dto.ShortenResponse{
		Result: fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL),
	}

	res.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(responseData)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = res.Write(response)
	if err != nil {
		logger.Log().Error("Can not send response from apiShortenHandle:", zap.Error(err))
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) pingHandle(res http.ResponseWriter, req *http.Request) {
	dbStorage, ok := s.storage.(*storage.DBStorage)
	if !ok {
		http.Error(res, "Service working not on top DB storage", http.StatusInternalServerError)
		return
	}

	if err := dbStorage.Ping(req.Context()); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s Server) apiShortenBatchHandle(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	requestData := &[]dto.BatchShortenRequest{}
	err = json.Unmarshal(body, requestData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	urlsMapToShortify := make(map[string]string)
	for _, batchRequest := range *requestData {
		urlsMapToShortify[batchRequest.CorrelationID] = batchRequest.URL
	}

	shortenedURLs, err := s.storage.StoreBatch(req.Context(), urlsMapToShortify)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var responseData []dto.BatchShortenResponse
	for correlationID, shortenedURL := range shortenedURLs {
		responseData = append(responseData, dto.BatchShortenResponse{
			CorrelationID: correlationID,
			ShortURL:      fmt.Sprintf("%s/%s", s.options.BaseURL, shortenedURL),
		})
	}

	res.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(responseData)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
	_, err = res.Write(response)
	if err != nil {
		logger.Log().Error("Can not send response from apiShortenHandle:", zap.Error(err))
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// NewRouter Creates router
func NewRouter(options *config.Options, storage *storage.Repository) *chi.Mux {
	r := chi.NewRouter()

	s := Server{
		options: options,
		storage: *storage,
	}
	r.Use(middlewares.GzipMiddleware)
	r.Use(logger.LoggerMiddleware)

	r.Post("/", s.postURLHandle)
	r.Get("/{id}", s.getShortURLHandle)
	r.Get("/ping", s.pingHandle)
	r.Post("/api/shorten", s.apiShortenHandle)
	r.Post("/api/shorten/batch", s.apiShortenBatchHandle)
	r.MethodNotAllowed(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusBadRequest)
		},
	)
	return r
}
