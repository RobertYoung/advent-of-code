package main

import (
	"testing"

	"github.com/RobertYoung/advent-of-code/util"
)

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
		points, _ := util.ConvertToPoints(test.input)
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
		points, _ := util.ConvertToPoints(test.input)
		result, _ := FindNumberOfDistinctTrails(points)

		if result != test.expected {
			t.Errorf("TestFindNumberOfTrails(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
