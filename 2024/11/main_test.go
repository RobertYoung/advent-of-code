package main

import (
	"testing"
)

func equal(result []int, i []int) bool {
	if len(result) != len(i) {
		return false
	}

	for j := 0; j < len(result); j++ {
		if result[j] != i[j] {
			return false
		}
	}

	return true
}

func TestSplitStringToNumbers(t *testing.T) {
	test := []struct {
		input    string
		expected []int
	}{
		{"0 1 10 99 999", []int{0, 1, 10, 99, 999}},
	}

	for _, test := range test {
		result, _ := SplitStringToNumbers(test.input)

		if !equal(result, test.expected) {
			t.Errorf("SplitStringToNumbers(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestBlinkNumTimes(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`125 17`, 22},
	}

	for _, test := range tests {
		numbers, _ := SplitStringToNumbers(test.input)
		result, _ := CountStones(numbers, 6)

		if result != test.expected {
			t.Errorf("TestBlinkNumTimes(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
