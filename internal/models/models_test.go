package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildID(t *testing.T) {
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

func Test_NewShortURL(t *testing.T) {
	tests := []struct {
		name      string
		value     []string
		wantValue ShortenURL
	}{
		{
			name:      "Short URL Generated in expected way",
			value:     []string{"1", "2"},
			wantValue: ShortenURL{UUID: "050c5d2e", OriginalURL: "1", UserID: "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewShortURL(tt.value[0], tt.value[1]), tt.wantValue, "NewShortURL generated incorrect value")
		})
	}
}
