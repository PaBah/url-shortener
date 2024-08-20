package middlewares

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/mock"
	"github.com/stretchr/testify/require"
)

func TestWhitelisting(t *testing.T) {

	t.Run("allowed_request", func(t *testing.T) {
		testMessage := `{"test":"test"}`
		_, trustedNet, _ := net.ParseCIDR("0.0.0.0/0")
		middlewareInstance := IPWhiteListMiddleware(trustedNet)(mock.NewHandlerMock(testMessage, http.StatusCreated))

		srv := httptest.NewServer(middlewareInstance)
		defer srv.Close()

		requestBody := `{"url":"https://practicum.yandex.ru/"}`

		buf := bytes.NewBufferString(requestBody)
		r := httptest.NewRequest("POST", srv.URL+"/api/shorten", buf)
		r.RequestURI = ""
		r.Header.Set("X-Real-IP", "127.0.0.1")
		r.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode)

		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		require.JSONEq(t, testMessage, string(b))
	})

	t.Run("not_allowed_request", func(t *testing.T) {
		testMessage := `{"test":"test"}`
		_, trustedNet, _ := net.ParseCIDR("192.0.2.1/24")
		middlewareInstance := IPWhiteListMiddleware(trustedNet)(mock.NewHandlerMock(testMessage, http.StatusCreated))

		srv := httptest.NewServer(middlewareInstance)
		defer srv.Close()

		requestBody := `{"url":"https://practicum.yandex.ru/"}`

		buf := bytes.NewBufferString(requestBody)
		r := httptest.NewRequest("POST", srv.URL+"/api/shorten", buf)
		r.RequestURI = ""
		r.Header.Set("X-Real-IP", "65.21.233.48")
		r.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusForbidden, resp.StatusCode)

		defer resp.Body.Close()
	})
}
