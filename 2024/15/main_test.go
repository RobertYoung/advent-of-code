package main

import (
	"testing"

	"github.com/RobertYoung/advent-of-code/util"
)

func TestCreateGamePart1(t *testing.T) {
	input := `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

	testsWarehouse := []struct {
		point    util.Point
		expected MapKey
	}{
		{util.Point{X: 0, Y: 0}, WALL},
		{util.Point{X: 1, Y: 1}, EMPTY},
		{util.Point{X: 2, Y: 2}, ROBOT},
		{util.Point{X: 4, Y: 3}, BOX},
	}

	testsMoves := []struct {
		index    int
		expected DirectionKey
	}{
		{0, LEFT},
		{1, UP},
		{2, UP},
		{3, RIGHT},
		{6, DOWN},
	}

	game := CreateGamePart1(input)

	for _, test := range testsWarehouse {
		if game.warehouse[test.point] != test.expected {
			t.Errorf("TestCreateGamePart1().warehouse = %v; want %v", game.warehouse[test.point], test.expected)
		}
	}

	for _, test := range testsMoves {
		if game.moves[test.index] != test.expected {
			t.Errorf("TestCreateGamePart1().moves = %v; want %v", game.moves[test.index], test.expected)
		}
	}
}

func TestCreateGamePart2(t *testing.T) {
	input := `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

	testsWarehouse := []struct {
		point    util.Point
		expected MapKey
	}{
		{util.Point{X: 0, Y: 0}, WALL},
		{util.Point{X: 1, Y: 1}, WALL},
		{util.Point{X: 10, Y: 3}, ROBOT},
		{util.Point{X: 13, Y: 6}, WALL},
		{util.Point{X: 2, Y: 1}, EMPTY},
		{util.Point{X: 11, Y: 5}, EMPTY},
	}

	testsMoves := []struct {
		index    int
		expected DirectionKey
	}{
		{0, LEFT},
		{1, DOWN},
		{2, DOWN},
		{3, LEFT},
		{6, UP},
	}

	game := CreateGamePart2(input)
	game.print()

	for _, test := range testsWarehouse {
		if game.warehouse[test.point] != test.expected {
			t.Errorf("TestCreateGamePart2().warehouse = %v; want %v", game.warehouse[test.point], test.expected)
		}
	}

	for _, test := range testsMoves {
		if game.moves[test.index] != test.expected {
			t.Errorf("TestCreateGamePart2().moves = %v; want %v", game.moves[test.index], test.expected)
		}
	}
}

func TestGameMovePart1(t *testing.T) {
	input := `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

	test := []struct {
		point    util.Point
		expected MapKey
	}{
		{util.Point{X: 0, Y: 0}, WALL},
		{util.Point{X: 6, Y: 1}, BOX},
		{util.Point{X: 4, Y: 4}, ROBOT},
	}

	game := CreateGamePart1(input)
	game.movePart1()

	for _, test := range test {
		if game.warehouse[test.point] != test.expected {
			t.Errorf("TestGameMovePart1() = %v; want %v", game.warehouse[test.point], test.expected)
		}
	}
}

func TestGameMovePart2(t *testing.T) {
	input := `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

	test := []struct {
		point    util.Point
		expected MapKey
	}{
		{util.Point{X: 0, Y: 0}, WALL},
		{util.Point{X: 5, Y: 1}, LEFT_BOX},
		{util.Point{X: 6, Y: 1}, RIGHT_BOX},
		{util.Point{X: 5, Y: 2}, ROBOT},
	}

	game := CreateGamePart2(input)
	game.movePart2()

	for _, test := range test {
		if game.warehouse[test.point] != test.expected {
			t.Errorf("TestGameMovePart2() = %v; want %v", game.warehouse[test.point], test.expected)
		}
	}
}

func TestGameCalculateGPS1(t *testing.T) {
	input := `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

	expected := 2028
	game := CreateGamePart1(input)
	game.movePart1()
	result := game.calculateGPS()

	if result != expected {
		t.Errorf("TestGameCalculateGPS1() = %v; want %v", result, expected)
	}
}

func TestGameCalculateGPS2(t *testing.T) {
	input := `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

	expected := 10092
	game := CreateGamePart1(input)
	game.movePart1()
	result := game.calculateGPS()

	if result != expected {
		t.Errorf("TestGameCalculateGPS() = %v; want %v", result, expected)
	}
}

func TestGameCalculatePart2GPS(t *testing.T) {
	input := `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

	expected := 9021
	game := CreateGamePart2(input)
	game.movePart2()

	result := game.calculateGPS()

	if result != expected {
		t.Errorf("TestGameCalculatePart2GPS() = %v; want %v", result, expected)
	}
}
