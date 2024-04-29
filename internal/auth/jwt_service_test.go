package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWTService(t *testing.T) {
	userID := GetUserID("1")
	assert.Equal(t, "", userID)
}
