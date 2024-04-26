package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthorizedMiddleware(t *testing.T) {
	handler := AuthorizedMiddleware(mock.NewHandlerMock(``, http.StatusUnauthorized))

	srv := httptest.NewServer(handler)
	defer srv.Close()

	t.Run("authorized_middleware", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, srv.URL+"/", nil)
		r.RequestURI = ""

		resp, err := http.DefaultClient.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		_, err = r.Cookie("Authorization")
		require.Error(t, err)

		defer resp.Body.Close()
	})
}

func TestAuthorizedMiddlewareSuccess(t *testing.T) {
	handler := AuthorizedMiddleware(mock.NewHandlerMock(``, http.StatusOK))

	srv := httptest.NewServer(handler)
	defer srv.Close()

	r := httptest.NewRequest(http.MethodGet, srv.URL+"/", nil)
	JWTToken, _ := BuildJWTString("1")
	r.Header.Set("Cookie", "Authorization="+JWTToken)
	r.RequestURI = ""

	resp, err := http.DefaultClient.Do(r)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	cookie, err := r.Cookie("Authorization")
	require.NoError(t, err)

	userID := GetUserID(cookie.Value)
	assert.Equal(t, "1", userID)

	defer resp.Body.Close()
}
