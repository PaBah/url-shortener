package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/mock"
	"github.com/stretchr/testify/require"
)

func TestPublicAuthorizationMiddleware(t *testing.T) {
	handler := PublicAuthorizationMiddleware(mock.NewHandlerMock(``, http.StatusOK))

	srv := httptest.NewServer(handler)
	defer srv.Close()

	t.Run("public_auth", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, srv.URL+"/", nil)
		r.RequestURI = ""

		resp, err := http.DefaultClient.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		_, err = r.Cookie("Authorization")
		require.Error(t, err)

		defer resp.Body.Close()
	})
}
