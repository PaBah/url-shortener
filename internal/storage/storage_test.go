package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &InMemoryStorage{
				state: tt.state,
			}
			gotData, err := cs.FindByID(tt.ID)
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
			state:    nil,
			value:    "https://practicum.yandex.ru/",
			wantData: "2187b119",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &InMemoryStorage{
				state: tt.state,
			}
			cs.Store(tt.value)
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
			cs := &InMemoryStorage{}
			assert.Equal(t, cs.buildID(tt.value), tt.wantValue, "Сгенерированный и ожидаемый ID не совпадают")
		})
	}
}
