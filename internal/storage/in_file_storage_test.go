package storage

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInFileStorage_FindByID(t *testing.T) {
	tests := []struct {
		name     string
		state    map[string]string
		ID       string
		wantData string
		wantErr  bool
	}{
		{
			name:     "Successfully found value",
			state:    map[string]string{"2187b119": "https://practicum.yandex.ru/"},
			ID:       "2187b119",
			wantData: "https://practicum.yandex.ru/",
			wantErr:  false,
		},
		{
			name:     "No value in store",
			state:    nil,
			ID:       "2187b119",
			wantData: "",
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
		value    string
		wantData string
	}{
		{
			name:     "With initialed store",
			state:    map[string]string{"2187b119": "https://practicum.yandex.ru/"},
			value:    "https://practicum.yandex.ru/",
			wantData: "2187b119",
		},

		{
			name:     "Empty store",
			state:    make(map[string]string),
			value:    "https://practicum.yandex.ru/",
			wantData: "2187b119",
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &InFileStorage{
				state: tt.state,
			}
			ID, _ := cs.Store(ctx, tt.value)
			assert.Equal(t, ID, tt.wantData)
			assert.Equal(t, cs.state[tt.wantData], tt.value, "Результат после добавления не совпадает с ожидаемым")
		})
	}
}

func TestInFileStorage_buildID(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantValue string
	}{
		{
			name:      "ID Generated in expected way",
			value:     "https://practicum.yandex.ru/",
			wantValue: "2187b119",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, buildID(tt.value), tt.wantValue, "Сгенерированный и ожидаемый ID не совпадают")
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

	shortURLs, err := fs.StoreBatch(context.Background(), map[string]string{"test1": "test", "test2": "test"})

	assert.NoError(t, err, "Batch value insertion not failed")
	assert.Equal(t, map[string]string{"test1": "bc2c0be9", "test2": "bc2c0be9"}, shortURLs, "All batch stored")
	_ = os.Remove("/tmp/.test_store")
}
