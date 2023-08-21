package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		w1  string
		w2  string
		exp bool
	}{
		{"cabbba", "abbccc", true},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.exp, closeStrings(tt.w1, tt.w2))
	}
}
