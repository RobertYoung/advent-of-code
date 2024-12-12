package util

import "testing"

func TestConvertToPointsInt(t *testing.T) {
	input := `0123
1234
8765
9876`
	result, _ := ConvertToPoints(input, ConvertToInt)

	tests := []struct {
		input    Point
		expected int
	}{
		{Point{X: 0, Y: 0}, 0},
		{Point{X: 2, Y: 2}, 6},
		{Point{X: 3, Y: 3}, 6},
	}

	for _, test := range tests {
		if result[test.input] != test.expected {
			t.Errorf("TestConvertToPoints(%v) = %v; want %v", result[test.input], result, test.expected)
		}
	}
}

func TestConvertToPointsString(t *testing.T) {
	input := `ZABC
ABCD
HGFE
IHGF`
	result, _ := ConvertToPoints(input, ConvertToString)

	tests := []struct {
		input    Point
		expected string
	}{
		{Point{X: 0, Y: 0}, "Z"},
		{Point{X: 2, Y: 2}, "F"},
		{Point{X: 3, Y: 3}, "F"},
	}

	for _, test := range tests {
		if result[test.input] != test.expected {
			t.Errorf("TestConvertToPointsString(%v) = %v; want %v", result[test.input], result, test.expected)
		}
	}
}
