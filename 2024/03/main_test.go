package main

import (
	"testing"
)

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestFixCorruptedLine(t *testing.T) {
	tests := []struct {
		input    string
		expected [][]int
	}{
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}},
	}

	for _, test := range tests {
		result := FixCorruptedLine(test.input)
		if !equal(result, test.expected) {
			t.Errorf("FixCorruptedLine(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestFixCorruptedLineConditional(t *testing.T) {
	tests := []struct {
		input    string
		expected [][]int
	}{
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))mul ( 2 , 4 )don't()xxx23mul(8,5)", [][]int{{2, 4}, {8, 5}}}}

	for _, test := range tests {
		result := FixCorruptedLineConditional(test.input)
		if !equal(result, test.expected) {
			t.Errorf("FixCorruptedLineConditional(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
