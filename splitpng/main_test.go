package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScale(t *testing.T) {
	assert.Equal(t, 1.0, scale(1, 0, 1))
}
