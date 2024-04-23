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
