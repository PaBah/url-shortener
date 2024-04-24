package storage

import (
	"context"
	"os"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestInFileStorage_FindByID(t *testing.T) {
	tests := []struct {
		name     string
		state    map[string]string
		ID       string
		wantData models.ShortenURL
		wantErr  bool
	}{
		{
			name:     "Successfully found value",
			state:    map[string]string{"2187b119": "https://practicum.yandex.ru/"},
			ID:       "2187b119",
			wantData: models.NewShortURL("https://practicum.yandex.ru/"),
			wantErr:  false,
		},
		{
			name:     "No value in store",
			state:    nil,
			ID:       "2187b119",
			wantData: models.NewShortURL(""),
			wantErr:  true,
		},
	}
	ctx := context.Background()
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
			assert.Equal(t, gotData, tt.wantData, "Не совпадает ожидаемое значение")
		})
	}
}

func TestInFileStorage_Store(t *testing.T) {
	tests := []struct {
		name     string
		state    map[string]string
		value    models.ShortenURL
		wantData string
	}{
		{
			name:     "With initialed store",
			state:    make(map[string]string),
			value:    models.NewShortURL("https://practicum.yandex.ru/"),
			wantData: "2187b119",
		},

		{
			name:     "Empty store",
			state:    make(map[string]string),
			value:    models.NewShortURL("https://practicum.yandex.ru/"),
			wantData: "2187b119",
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &InFileStorage{
				state: tt.state,
			}
			err := cs.Store(ctx, tt.value)
			assert.NoError(t, err)
			assert.Equal(t, cs.state[tt.wantData], tt.value.OriginalURL, "Результат после добавления не совпадает с ожидаемым")
		})
	}
}

func TestWorkWithFile(t *testing.T) {
	fs := NewInFileStorage("/tmp/.test_store")
	defer fs.Close()

	fs.state = map[string]string{"test": "test"}

	err := fs.writeBackup()
	assert.NoError(t, err, "data had been written with error")

	fs.state = nil
	fs.initialize("/tmp/.test_store")

	assert.Equal(t, fs.state, map[string]string{"test": "test"}, "data had been read with error")
	_ = os.Remove("/tmp/.test_store")
}

func TestInFileStorage_StoreBatch(t *testing.T) {
	fs := NewInFileStorage("/tmp/.test_store")
	defer fs.Close()
	shortURLs := map[string]models.ShortenURL{
		"test1": models.NewShortURL("test"),
		"test2": models.NewShortURL("test"),
	}
	err := fs.StoreBatch(context.Background(), shortURLs)

	assert.NoError(t, err, "Batch value insertion not failed")
	_ = os.Remove("/tmp/.test_store")
}
