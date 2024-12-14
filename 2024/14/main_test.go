package main

import (
	"testing"

	"github.com/RobertYoung/advent-of-code/util"
)

func TestCreateGame(t *testing.T) {
	input := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	tests := []struct {
		id         int
		point      util.Point
		expectedVX int
		expectedVY int
	}{
		{0, util.Point{X: 0, Y: 4}, 3, -3},
		{11, util.Point{X: 9, Y: 5}, -3, -3},
	}

	game, _ := CreateGame(input, 10, 10)

	for _, test := range tests {
		if game.robots[test.id].point != test.point {
			t.Errorf("TestCreateGame().point = %v; want %v", game.robots[test.id].point, test.point)
		}
		if game.robots[test.id].vx != test.expectedVX {
			t.Errorf("TestCreateGame().vx = %v; want %v", game.robots[test.id].vx, test.expectedVX)
		}
		if game.robots[test.id].vy != test.expectedVY {
			t.Errorf("TestCreateGame().vy = %v; want %v", game.robots[test.id].vy, test.expectedVY)
		}
	}
}

func TestMoveRobot1Time(t *testing.T) {
	game := Game{mapX: 10, mapY: 7, robots: map[int]Robot{}}
	tests := []struct {
		robot     Robot
		expectedX int
		expectedY int
	}{
		{robot: Robot{game: &game, point: util.Point{X: 0, Y: 4}, vx: 3, vy: -3}, expectedX: 3, expectedY: 1},
		{robot: Robot{game: &game, point: util.Point{X: 9, Y: 5}, vx: -3, vy: -3}, expectedX: 6, expectedY: 2},
		{robot: Robot{game: &game, point: util.Point{X: 0, Y: 0}, vx: -3, vy: -3}, expectedX: 7, expectedY: 4},
		{robot: Robot{game: &game, point: util.Point{X: 9, Y: 6}, vx: 3, vy: 3}, expectedX: 2, expectedY: 2},
	}

	for _, test := range tests {
		test.robot.move(1)

		if test.robot.point.X != test.expectedX {
			t.Errorf("TestMoveRobot().x = %v; want %v", test.robot.point.X, test.expectedX)
		}
		if test.robot.point.Y != test.expectedY {
			t.Errorf("TestMoveRobot().y = %v; want %v", test.robot.point.Y, test.expectedY)
		}
	}
}

func TestMoveRobot(t *testing.T) {
	game := Game{mapX: 11, mapY: 7, robots: map[int]Robot{}}
	robot := Robot{game: &game, point: util.Point{X: 2, Y: 4}, vx: 2, vy: -3}

	tests := map[int]util.Point{
		1: {X: 4, Y: 1},
		2: {X: 6, Y: 5},
		3: {X: 8, Y: 2},
		4: {X: 10, Y: 6},
		5: {X: 1, Y: 3},
	}

	for _, test := range tests {
		robot.move(1)

		if robot.point.X != test.X {
			t.Errorf("TestMoveRobot().x = %v; want %v", robot.point.X, test.X)
		}
		if robot.point.Y != test.Y {
			t.Errorf("TestMoveRobot().y = %v; want %v", robot.point.Y, test.Y)
		}
	}
}

func TestCalculateQuadrant(t *testing.T) {
	game := Game{mapX: 11, mapY: 7, robots: map[int]Robot{}}
	tests := []struct {
		robot    Robot
		quadrant Quadrant
	}{
		{Robot{game: &game, point: util.Point{X: 2, Y: 2}}, TOP_LEFT},
		{Robot{game: &game, point: util.Point{X: 7, Y: 1}}, TOP_RIGHT},
		{Robot{game: &game, point: util.Point{X: 2, Y: 6}}, BOTTOM_LEFT},
		{Robot{game: &game, point: util.Point{X: 8, Y: 5}}, BOTTOM_RIGHT},
		{Robot{game: &game, point: util.Point{X: 5, Y: 6}}, UNKNOWN},
		{Robot{game: &game, point: util.Point{X: 3, Y: 3}}, UNKNOWN},
	}

	for _, test := range tests {
		test.robot.setQuadrant()

		if test.robot.quadrant != test.quadrant {
			t.Errorf("TestCalculateQuadrant().quadrant = %v; want %v", test.robot.quadrant, test.quadrant)
		}
	}
}

func TestCalculateGameSafetyFactor(t *testing.T) {
	input := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	game, _ := CreateGame(input, 11, 7)
	game.moveRobots(100)
	result := game.calculateSafetyFactor()

	if result != 12 {
		t.Errorf("TestCalculateGameSafetyFactor() = %v; want %v", result, 12)
	}
}
