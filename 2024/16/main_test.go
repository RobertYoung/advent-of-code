package main

import (
	"testing"

	"github.com/RobertYoung/advent-of-code/util"
)

func TestCreateGame(t *testing.T) {
	input := `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

	tests := []struct {
		point    util.Point
		expected MapKey
	}{
		{util.Point{X: 0, Y: 0}, WALL},
		{util.Point{X: 1, Y: 1}, EMPTY},
		{util.Point{X: 1, Y: 13}, START},
		{util.Point{X: 13, Y: 1}, END},
	}

	result := CreateGame(input)

	for _, test := range tests {
		if result.points[test.point] != test.expected {
			t.Errorf("TestCreateGame() = %v; want %v", result.points[test.point], test.expected)
		}
	}

	score, tiles := result.findShortestRoute()

	if score != 7036 {
		t.Errorf("TestCreateGame().score = %v; want %v", score, 7036)
	}

	if tiles != 45 {
		t.Errorf("TestCreateGame().tiles = %v; want %v", tiles, 45)
	}
}

func TestCreateGame2(t *testing.T) {
	input := `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

	tests := []struct {
		point    util.Point
		expected MapKey
	}{
		{util.Point{X: 0, Y: 0}, WALL},
		{util.Point{X: 1, Y: 1}, EMPTY},
		{util.Point{X: 1, Y: 15}, START},
		{util.Point{X: 15, Y: 1}, END},
	}

	result := CreateGame(input)

	for _, test := range tests {
		if result.points[test.point] != test.expected {
			t.Errorf("TestCreateGame2() = %v; want %v", result.points[test.point], test.expected)
		}
	}

	score, tiles := result.findShortestRoute()

	if score != 11048 {
		t.Errorf("TestCreateGame2().score = %v; want %v", score, 11048)
	}

	if tiles != 64 {
		t.Errorf("TestCreateGame2().tiles = %v; want %v", tiles, 64)
	}
}
