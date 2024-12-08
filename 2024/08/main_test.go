package main

import (
	"testing"
)

func TestGetAntinodesPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"............",
				"........0...",
				".....0......",
				".......0....",
				"....0.......",
				"......A.....",
				"............",
				"............",
				"........A...",
				".........A..",
				"............",
				"............",
			}, 14},
	}

	for _, test := range tests {

		antennas := ConvertToAntennas(test.input)
		result, _ := GetAntinodes(antennas, len(test.input[0]), len(test.input))

		if len(result) != test.expected {
			t.Errorf("TestGetAntinodesPart1(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestGetAntinodesPart2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"............",
				"........0...",
				".....0......",
				".......0....",
				"....0.......",
				"......A.....",
				"............",
				"............",
				"........A...",
				".........A..",
				"............",
				"............",
			}, 34},
	}

	for _, test := range tests {

		antennas := ConvertToAntennas(test.input)
		_, result := GetAntinodes(antennas, len(test.input[0]), len(test.input))

		if len(result) != test.expected {
			t.Errorf("TestGetAntinodesPart2(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
