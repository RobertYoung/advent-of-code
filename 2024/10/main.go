package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

type Point struct {
	x int
	y int
}

type Direction struct {
	y    int
	x    int
	name string
}

var directionMap = map[string]Direction{
	"up":    {-1, 0, "up"},
	"down":  {1, 0, "down"},
	"left":  {0, -1, "left"},
	"right": {0, 1, "right"},
}

func ConvertToPoints(input string) (map[Point]int, error) {
	result := map[Point]int{}

	for y, row := range strings.Split(input, "\n") {
		for x, char := range strings.Split(row, "") {
			value, err := strconv.Atoi(char)
			if err == nil {
				result[Point{x: x, y: y}] = value
			} else {
				result[Point{x: x, y: y}] = -1
			}
		}
	}

	return result, nil
}

func findTrails(trails map[Point]int, points map[Point]int, curr Point) {
	trailIndex := points[curr]
	nextTrailIndex := trailIndex + 1

	if trailIndex == 9 {
		trails[curr] = trailIndex
		return
	}

	upPoint := Point{x: curr.x + directionMap["up"].x, y: curr.y + directionMap["up"].y}
	downPoint := Point{x: curr.x + directionMap["down"].x, y: curr.y + directionMap["down"].y}
	leftPoint := Point{x: curr.x + directionMap["left"].x, y: curr.y + directionMap["left"].y}
	rightPoint := Point{x: curr.x + directionMap["right"].x, y: curr.y + directionMap["right"].y}

	if points[upPoint] == nextTrailIndex {
		findTrails(trails, points, upPoint)
	}

	if points[downPoint] == nextTrailIndex {
		findTrails(trails, points, downPoint)
	}

	if points[leftPoint] == nextTrailIndex {
		findTrails(trails, points, leftPoint)
	}

	if points[rightPoint] == nextTrailIndex {
		findTrails(trails, points, rightPoint)
	}
}

func findDistinctTrails(trails *int, points map[Point]int, curr Point) {
	trailIndex := points[curr]
	nextTrailIndex := trailIndex + 1

	if trailIndex == 9 {
		*trails++
		return
	}

	upPoint := Point{x: curr.x + directionMap["up"].x, y: curr.y + directionMap["up"].y}
	downPoint := Point{x: curr.x + directionMap["down"].x, y: curr.y + directionMap["down"].y}
	leftPoint := Point{x: curr.x + directionMap["left"].x, y: curr.y + directionMap["left"].y}
	rightPoint := Point{x: curr.x + directionMap["right"].x, y: curr.y + directionMap["right"].y}

	if points[upPoint] == nextTrailIndex {
		findDistinctTrails(trails, points, upPoint)
	}

	if points[downPoint] == nextTrailIndex {
		findDistinctTrails(trails, points, downPoint)
	}

	if points[leftPoint] == nextTrailIndex {
		findDistinctTrails(trails, points, leftPoint)
	}

	if points[rightPoint] == nextTrailIndex {
		findDistinctTrails(trails, points, rightPoint)
	}
}

func FindNumberOfTrails(points map[Point]int) (int, error) {
	count := 0

	for point, value := range points {
		trails := map[Point]int{}

		if value == 0 {
			findTrails(trails, points, point)
		}

		count += len(trails)
	}

	return count, nil
}

func FindNumberOfDistinctTrails(points map[Point]int) (int, error) {
	count := 0

	for point, value := range points {
		trails := 0

		if value == 0 {
			findDistinctTrails(&trails, points, point)
		}

		count += trails
	}

	return count, nil
}

func main() {
	input, _ := util.ReadFileAsString("input.txt")
	points, _ := ConvertToPoints(input)
	part1, _ := FindNumberOfTrails(points)
	part2, _ := FindNumberOfDistinctTrails(points)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
