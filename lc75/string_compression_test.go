package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		in       []byte
		expected int
	}{
		{
			in:       []byte("aabbeeee"),
			expected: 6,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, compress(tt.in), tt.expected)
	}
}
