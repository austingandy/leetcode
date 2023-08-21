package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{
			in:       []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, pivotIndex(tt.in))
	}
}
