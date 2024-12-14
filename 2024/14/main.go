package main

import (
	"fmt"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

type Quadrant int32

const (
	UNKNOWN      Quadrant = -1
	TOP_LEFT     Quadrant = 0
	TOP_RIGHT    Quadrant = 1
	BOTTOM_LEFT  Quadrant = 2
	BOTTOM_RIGHT Quadrant = 3
)

type Game struct {
	mapX   int
	mapY   int
	robots map[int]Robot
}

func (game *Game) moveRobots(seconds int) {
	for index := range game.robots {
		robot := game.robots[index]
		(&robot).move(seconds)
		game.robots[index] = robot
	}
}

func (game *Game) calculateSafetyFactor() int {
	quadrants := []int{0, 0, 0, 0}

	for _, robot := range game.robots {
		if robot.quadrant == UNKNOWN {
			continue
		}

		quadrants[robot.quadrant]++
	}

	result := 1

	for _, quadrant := range quadrants {
		result *= quadrant
	}

	return result
}

func (game *Game) createGrid() ([][]int, string) {
	grid := make([][]int, game.mapY)
	for i := range grid {
		grid[i] = make([]int, game.mapX)
	}

	for _, robot := range game.robots {
		grid[robot.point.Y][robot.point.X]++
	}

	print := ""

	for _, y := range grid {
		for _, x := range y {
			if x > 0 {
				print += "#"
			} else {
				print += "."
			}
		}

		print += "\n"
	}

	return grid, print
}

func (game *Game) findChristmasTree() int {
	for i := 0; i < 100000; i++ {
		game.moveRobots(1)
		_, print := game.createGrid()

		if strings.Contains(print, "##################") {
			fmt.Println(print)
			return i + 1
		}
	}

	return -1
}

type Robot struct {
	game     *Game
	point    util.Point
	vx       int
	vy       int
	quadrant Quadrant
}

func (robot *Robot) setQuadrant() Quadrant {
	if robot.point.X < robot.game.mapX/2 && robot.point.Y < robot.game.mapY/2 {
		robot.quadrant = TOP_LEFT
	} else if robot.point.X > robot.game.mapX/2 && robot.point.Y < robot.game.mapY/2 {
		robot.quadrant = TOP_RIGHT
	} else if robot.point.X < robot.game.mapX/2 && robot.point.Y > robot.game.mapY/2 {
		robot.quadrant = BOTTOM_LEFT
	} else if robot.point.X > robot.game.mapX/2 && robot.point.Y > robot.game.mapY/2 {
		robot.quadrant = BOTTOM_RIGHT
	} else {
		robot.quadrant = UNKNOWN
	}

	return robot.quadrant
}

func (robot *Robot) move(seconds int) {
	for i := 0; i < seconds; i++ {
		x := robot.point.X + robot.vx
		y := robot.point.Y + robot.vy

		if x >= robot.game.mapX {
			x = robot.point.X + robot.vx - robot.game.mapX
		} else if x < 0 {
			x = robot.game.mapX + x
		}

		if y >= robot.game.mapY {
			y = robot.point.Y + robot.vy - robot.game.mapY
		} else if y < 0 {
			y = robot.game.mapY + y
		}

		robot.point.X = x
		robot.point.Y = y

		robot.setQuadrant()
	}
}

func CreateGame(input string, mapX int, mapY int) (Game, error) {
	game := Game{
		mapX:   mapX,
		mapY:   mapY,
		robots: map[int]Robot{},
	}

	for index, line := range strings.Split(input, "\n") {
		robot := Robot{
			quadrant: UNKNOWN,
		}
		robot.point = util.Point{}
		robot.game = &game

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.point.X, &robot.point.Y, &robot.vx, &robot.vy)
		game.robots[index] = robot
	}

	return game, nil
}

func main() {
	input, _ := util.ReadFileAsString("input.txt")
	part1Game, _ := CreateGame(input, 101, 103)
	part1Game.moveRobots(100)
	part1 := part1Game.calculateSafetyFactor()

	fmt.Println("Part 1:", part1)

	part2Game, _ := CreateGame(input, 101, 103)
	part2 := part2Game.findChristmasTree()

	fmt.Println("Part 2:", part2)
}
