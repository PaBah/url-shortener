package storage

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/auth"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestInFileStorage_FindByID(t *testing.T) {
	tests := []struct {
		name     string
		state    map[string]models.ShortenURL
		ID       string
		wantData models.ShortenURL
		wantErr  bool
	}{
		{
			name:     "Successfully found value",
			state:    map[string]models.ShortenURL{"2187b119": models.NewShortURL("https://practicum.yandex.ru/", "1")},
			ID:       "2187b119",
			wantData: models.NewShortURL("https://practicum.yandex.ru/", "1"),
			wantErr:  false,
		},
		{
			name:     "No value in store",
			state:    nil,
			ID:       "2187b119",
			wantData: models.ShortenURL{UUID: "", UserID: "", OriginalURL: ""},
			wantErr:  true,
		},
	}
	ctx := context.WithValue(context.Background(), auth.ContextUserKey, 1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &InFileStorage{
				state: tt.state,
			}
			gotData, err := cs.FindByID(ctx, tt.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantData, gotData, "Не совпадает ожидаемое значение")
		})
	}
}

func TestInFileStorage_Store(t *testing.T) {
	tests := []struct {
		name     string
		state    map[string]models.ShortenURL
		wantData string
		value    models.ShortenURL
	}{
		{
			name:     "With initialed store",
			state:    map[string]models.ShortenURL{"2a49568d": models.NewShortURL("https://practicum.yandex.kz/", "1")},
			value:    models.NewShortURL("https://practicum.yandex.ru/", "1"),
			wantData: "2187b119",
		},

		{
			name:     "Empty store",
			state:    make(map[string]models.ShortenURL),
			value:    models.NewShortURL("https://practicum.yandex.ru/", "1"),
			wantData: "2187b119",
		},
	}
	ctx := context.WithValue(context.Background(), auth.ContextUserKey, "1")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &InFileStorage{
				state: tt.state,
			}
			err := cs.Store(ctx, tt.value)
			assert.NoError(t, err)
			assert.Equal(t, cs.state[tt.wantData], tt.value, "Результат после добавления не совпадает с ожидаемым")
		})
	}
}

func TestWorkWithFile(t *testing.T) {
	fs := NewInFileStorage("/tmp/.test_store")
	defer fs.Close()

	fs.state = map[string]models.ShortenURL{"bc2c0be9": models.NewShortURL("test", "1")}

	err := fs.writeBackup()
	assert.NoError(t, err, "data had been written with error")

	fs.state = nil
	fs.initialize("/tmp/.test_store")

	assert.Equal(t, fs.state, map[string]models.ShortenURL{"bc2c0be9": models.NewShortURL("test", "1")}, "data had been read with error")
	_ = os.Remove("/tmp/.test_store")
}

func TestInFileStorage_StoreBatch(t *testing.T) {
	fs := NewInFileStorage("/tmp/.test_store")
	defer fs.Close()
	shortURLs := map[string]models.ShortenURL{
		"test1": models.NewShortURL("test", "1"),
		"test2": models.NewShortURL("test", "1"),
	}
	err := fs.StoreBatch(context.Background(), shortURLs)

	assert.NoError(t, err, "Batch value insertion not failed")
	_ = os.Remove("/tmp/.test_store")
}

func TestInFileStorage_GetAllUsers(t *testing.T) {
	fs := NewInFileStorage("/tmp/.test_store")
	defer fs.Close()
	shortURLs := map[string]models.ShortenURL{
		"test1": models.NewShortURL("test", "1"),
		"test2": models.NewShortURL("test", "1"),
	}
	fs.state = shortURLs

	ctx := context.WithValue(context.Background(), auth.ContextUserKey, "1")
	data, err := fs.GetAllUsers(ctx)
	assert.NoError(t, err)
	assert.Equal(t, []models.ShortenURL{models.NewShortURL("test", "1"), models.NewShortURL("test", "1")}, data, "Found message scanned correctly")
	assert.NoError(t, err, "Batch value insertion not failed")
	_ = os.Remove("/tmp/.test_store")
}

func TestInFileStorage_AsyncCheckURLsUserID(t *testing.T) {
	fs := NewInFileStorage("/tmp/.test_store")
	defer fs.Close()
	shortURLs := map[string]models.ShortenURL{
		"bc2c0be9": models.NewShortURL("test", "test"),
	}
	fs.state = shortURLs
	shortURLCh := make(chan string)
	res := fs.AsyncCheckURLsUserID("test", shortURLCh)
	shortURLCh <- "bc2c0be9"
	assert.Equal(t, <-res, "bc2c0be9", "check was successful")
	_ = os.Remove("/tmp/.test_store")
}

func TestInFileStorage_DeleteShortURLs(t *testing.T) {
	fs := NewInFileStorage("/tmp/.test_store")
	defer fs.Close()
	shortURLs := map[string]models.ShortenURL{
		"bc2c0be9": models.NewShortURL("test", "test"),
	}
	fs.state = shortURLs
	err := fs.DeleteShortURLs(context.Background(), []string{"test"})
	assert.NoError(t, err, "successfully deleted urls")
	_ = os.Remove("/tmp/.test_store")
}

func BenchmarkInFileStorage_StoreBatch(b *testing.B) {
	fs := NewInFileStorage("/tmp/.test_store")
	defer fs.Close()
	for i := 0; i < b.N; i++ {
		testFirst, testSecond := fmt.Sprintf("%d%s", i, "test"), fmt.Sprintf("%d%s", i, "2test")
		shortURLs := map[string]models.ShortenURL{
			testFirst:  models.NewShortURL("test", "1"),
			testSecond: models.NewShortURL("test", "1"),
		}
		_ = fs.StoreBatch(context.Background(), shortURLs)
	}
}
