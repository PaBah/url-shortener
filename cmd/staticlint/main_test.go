package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCheckers(t *testing.T) {
	assert.Equal(t, len(getCheckers()), 133)
}
