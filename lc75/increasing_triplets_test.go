package main

import "testing"

func TestIncreasingTriplets(t *testing.T) {
	tests := []struct {
		in       []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
		{[]int{6, 7, 1, 2}, false},
		{[]int{20, 100, 10, 12, 5, 13}, true},
	}
	for _, tt := range tests {
		if actual := increasingTriplet(tt.in); actual != tt.expected {
			t.Errorf("increasingTriplet(%v): expected %v, actual %v", tt.in, tt.expected, actual)
		}
		if actual := dumbIncreasingTriplet(tt.in); actual != tt.expected {
			t.Errorf("dumbIncreasingTriplet(%v): expected %v, actual %v", tt.in, tt.expected, actual)
		}
	}
}
