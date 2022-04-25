package main

import (
	"math"
	"testing"
)

func TestFindAverge(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput float64
	}{
		{[]int{10, 30, 20}, 20.0},
		{[]int{}, 0},
		{[]int{34}, 34},
		{[]int{math.MaxInt, math.MaxInt}, math.MaxInt},
	}
	for _, currentPair := range tests {
		actualResult := findAverage(currentPair.input)
		if actualResult != currentPair.expectedOutput {
			t.Error("test failed")
		}
	}
}
