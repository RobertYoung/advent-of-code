package main

import (
	"testing"
)

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
		{Point{x: 0, y: 0}, 0},
		{Point{x: 2, y: 2}, 6},
		{Point{x: 3, y: 3}, 6},
	}

	for _, test := range tests {
		if result[test.input] != test.expected {
			t.Errorf("TestConvertToPoints(%v) = %v; want %v", result[test.input], result, test.expected)
		}
	}
}

func TestFindNumberOfTrails(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`0123
1234
8765
9876`, 1},
		{`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, 36},
	}

	for _, test := range tests {
		points, _ := ConvertToPoints(test.input)
		result, _ := FindNumberOfTrails(points)

		if result != test.expected {
			t.Errorf("TestFindNumberOfTrails(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestFindNumberOfDistinctTrails(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`012345
123456
234567
345678
4.6789
56789.`, 227},
	}

	for _, test := range tests {
		points, _ := ConvertToPoints(test.input)
		result, _ := FindNumberOfDistinctTrails(points)

		if result != test.expected {
			t.Errorf("TestFindNumberOfTrails(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
