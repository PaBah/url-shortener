package server

import (
	"bytes"
	"compress/gzip"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGzipCompression(t *testing.T) {
	options := &config.Options{
		ServerAddress: ":8080",
		BaseURL:       "http://localhost:8080",
	}

	ctrl := gomock.NewController(t)
	var store storage.Repository
	rm := storage.NewMockRepository(ctrl)

	store = rm

	rm.
		EXPECT().
		Store(gomock.Eq("https://practicum.yandex.ru/")).
		Return("2187b119").
		AnyTimes()
	rm.
		EXPECT().
		FindByID("2187b119").
		Return("https://practicum.yandex.ru/", nil).
		AnyTimes()

	s := Server{
		options: options,
		storage: &store,
	}

	handler := GzipMiddleware(s.apiShortenHandle)

	srv := httptest.NewServer(handler)
	defer srv.Close()

	requestBody := `{"url":"https://practicum.yandex.ru/"}`

	successBody := `{"result":"http://localhost:8080/2187b119"}`

	t.Run("sends_gzip", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		zb := gzip.NewWriter(buf)
		_, err := zb.Write([]byte(requestBody))
		require.NoError(t, err)
		err = zb.Close()
		require.NoError(t, err)

		r := httptest.NewRequest("POST", srv.URL+"/api/shorten", buf)
		r.RequestURI = ""
		r.Header.Set("Content-Encoding", "gzip")
		r.Header.Set("Accept-Encoding", "")
		r.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode)

		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		require.JSONEq(t, successBody, string(b))
	})

	t.Run("accepts_gzip", func(t *testing.T) {
		buf := bytes.NewBufferString(requestBody)
		r := httptest.NewRequest("POST", srv.URL+"/api/shorten", buf)
		r.RequestURI = ""
		r.Header.Set("Accept-Encoding", "gzip")
		r.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode)

		defer resp.Body.Close()

		zr, err := gzip.NewReader(resp.Body)
		require.NoError(t, err)

		b, err := io.ReadAll(zr)
		require.NoError(t, err)

		require.JSONEq(t, successBody, string(b))
	})
}
