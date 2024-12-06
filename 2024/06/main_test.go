package main

import (
	"testing"
)

func TestCountPositions(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected int
	}{
		{[][]string{
			{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
			{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		}, 41},
	}

	for _, test := range tests {
		result := CountPositions(test.input)

		if result != test.expected {
			t.Errorf("CountPositions(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestCountLoopObstructions(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected int
	}{
		{[][]string{
			{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
			{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		}, 6},
	}

	for _, test := range tests {
		result := CountLoopObstructions(test.input)

		if result != test.expected {
			t.Errorf("TestCountObstructions(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
