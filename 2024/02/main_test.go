package main

import (
	"testing"
)

func TestCheckIfSafe(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		result := CheckIfSafe(test.input)
		if result != test.expected {
			t.Errorf("CheckIfSafe(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsSafeReportWithTolerance(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		result := IsSafeReportWithTolerance(test.input)
		if result != test.expected {
			t.Errorf("IsSafeReportWithTolerance(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
