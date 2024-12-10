package util

import "testing"

func TestConvertToPoints(t *testing.T) {
	input := `0123
1234
8765
9876`
	result, _ := ConvertToPoints(input)

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
