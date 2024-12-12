package main

import (
	"testing"

	"github.com/RobertYoung/advent-of-code/util"
)

func TestFindRegions(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	points, _ := util.ConvertToPoints(input, util.ConvertToString)
	result, _ := FindPlotPoints(points)

	tests := [][]util.Point{
		{util.Point{X: 0, Y: 0}, util.Point{X: 1, Y: 0}, util.Point{X: 2, Y: 0}, util.Point{X: 3, Y: 0}},
		{util.Point{X: 0, Y: 1}, util.Point{X: 1, Y: 1}, util.Point{X: 0, Y: 2}, util.Point{X: 1, Y: 2}},
		{util.Point{X: 2, Y: 1}, util.Point{X: 2, Y: 2}, util.Point{X: 3, Y: 2}, util.Point{X: 3, Y: 3}},
		{util.Point{X: 3, Y: 1}},
		{util.Point{X: 0, Y: 3}, util.Point{X: 1, Y: 3}, util.Point{X: 2, Y: 3}},
	}

	for _, test := range tests {
		for _, point := range test {
			if result[point].region != result[test[0]].region {
				t.Errorf("TestFindRegions(%v) = %v; want %v", result[point].region, result[point].region, result[test[0]].region.id)
			}
		}
	}
}

func TestFindNestedRegions(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	points, _ := util.ConvertToPoints(input, util.ConvertToString)
	result, _ := FindPlotPoints(points)
	regions := map[int]Region{}

	for _, plot := range result {
		if _, exists := regions[plot.region.id]; !exists {
			regions[plot.region.id] = plot.region
		}
	}

	if len(regions) != 5 {
		t.Errorf("TestFindNestedRegions() = %v; want %v", len(regions), 5)
	}
}

func TestCalculateFenceCost(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`AAAA
BBCD
BBCC
EEEC`, 140},
	}

	for _, test := range tests {
		points, _ := util.ConvertToPoints(test.input, util.ConvertToString)
		plots, _ := FindPlotPoints(points)
		cost, _ := CalculateFenceCost(plots)

		if cost != test.expected {
			t.Errorf("CalculateFenceCost() = %v; want %v", cost, test.expected)
		}
	}
}

func TestCalculateFenceCostWithDiscount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`AAAA
BBCD
BBCC
EEEC`, 80},
	}

	for _, test := range tests {
		points, _ := util.ConvertToPoints(test.input, util.ConvertToString)
		plots, _ := FindPlotPoints(points)
		cost, _ := CalculateFenceCostWithDiscount(plots)

		if cost != test.expected {
			t.Errorf("CalculateFenceCostWithDiscount() = %v; want %v", cost, test.expected)
		}
	}
}
