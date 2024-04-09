package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

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
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	shortURL := (*s.storage).Store(string(body))
	shortenedURL := fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL)
	res.Header().Set("Content-Type", "")
	res.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))
	res.WriteHeader(http.StatusCreated)
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

	shortURL := (*s.storage).Store(requestData.URL)
	responseData := dto.ShortenResponse{
		Result: fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL),
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

func (s Server) pingHandler(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("pgx", s.options.DatabaseDSN)
	if err != nil {
		logger.Log().Error("Server can not connect to DB ", zap.Error(err))
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// NewRouter Creates router
func NewRouter(options *config.Options, storage *storage.Repository) *chi.Mux {
	r := chi.NewRouter()

	s := Server{
		options: options,
		storage: storage,
	}
	r.Use(middlewares.GzipMiddleware)
	r.Use(logger.LoggerMiddleware)

	r.Post("/", s.postURLHandle)
	r.Get("/{id}", s.getShortURLHandle)
	r.Get("/ping", s.pingHandler)
	r.Post("/api/shorten", s.apiShortenHandle)
	r.MethodNotAllowed(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusBadRequest)
		},
	)
	return r
}
