package main

import (
	"testing"
)

func TestFindWord(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			}, 18},
	}

	for _, test := range tests {
		result := FindWord(test.input, "XMAS")
		if result != test.expected {
			t.Errorf("TestFindWord(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestFindPattern(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			}, 9},
	}

	for _, test := range tests {
		result := FindXPattern(test.input, "MAS")
		if result != test.expected {
			t.Errorf("FindXPattern(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
