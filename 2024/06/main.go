package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

type Position struct {
	x int
	y int
}

type PositionType int32

const (
	EMPTY    PositionType = 0
	GUARD    PositionType = 1
	OBSTACLE PositionType = 2
)

type Direction struct {
	y    int
	x    int
	name string
}

type VisitedPosition struct {
	position   Position
	directions []string
}

var directionMap = map[string]Direction{
	"up":    {-1, 0, "up"},
	"down":  {1, 0, "down"},
	"left":  {0, -1, "left"},
	"right": {0, 1, "right"},
}

var positionMap = map[string]PositionType{
	".": PositionType(EMPTY),
	"^": PositionType(GUARD),
	"#": PositionType(OBSTACLE),
}

func findVisitedPositions(visitedPositions map[string]VisitedPosition, guardMap [][]string, guardIndex Position, direction Direction) (map[string]VisitedPosition, bool) {
	var isDone = false

	currPosition := Position{
		x: guardIndex.x,
		y: guardIndex.y,
	}

	nextPosition := Position{
		x: guardIndex.x,
		y: guardIndex.y,
	}

	for !isDone {
		isDone = false

		currPosition = nextPosition
		nextPosition.x = nextPosition.x + direction.x
		nextPosition.y = nextPosition.y + direction.y

		if nextPosition.x < 0 || nextPosition.x >= len(guardMap[0]) || nextPosition.y < 0 || nextPosition.y >= len(guardMap) {
			isDone = true
			visitedPositions[fmt.Sprintf("%d,%d", currPosition.y, currPosition.x)] = VisitedPosition{
				position:   currPosition,
				directions: append(visitedPositions[fmt.Sprintf("%d,%d", currPosition.y, currPosition.x)].directions, direction.name),
			}
			break
		}

		// detect loop
		if visitedPosition, ok := visitedPositions[fmt.Sprintf("%d,%d", currPosition.y, currPosition.x)]; ok {
			if slices.Contains(visitedPosition.directions, direction.name) {
				isDone = true
				return visitedPositions, true
			}
		}

		if positionMap[guardMap[nextPosition.y][nextPosition.x]] == OBSTACLE {
			isDone = true
			nextDirection := directionMap["up"]

			switch direction.name {
			case "up":
				nextDirection = directionMap["right"]
			case "down":
				nextDirection = directionMap["left"]
			case "left":
				nextDirection = directionMap["up"]
			case "right":
				nextDirection = directionMap["down"]
			}

			return findVisitedPositions(visitedPositions, guardMap, currPosition, nextDirection)
		} else {
			visitedPositions[fmt.Sprintf("%d,%d", currPosition.y, currPosition.x)] = VisitedPosition{
				position:   currPosition,
				directions: append(visitedPositions[fmt.Sprintf("%d,%d", currPosition.y, currPosition.x)].directions, direction.name),
			}
		}
	}

	return visitedPositions, false
}

func findGuardIndex(guardMap [][]string) Position {
	guardIndex := Position{x: -1, y: -1}

	for y, line := range guardMap {
		for x, position := range line {
			if positionMap[position] == GUARD {
				guardIndex = Position{x: x, y: y}
				break
			}
		}

		if guardIndex.x != -1 && guardIndex.y != -1 {
			break
		}
	}

	return guardIndex
}

func CountPositions(guardMap [][]string) int {
	guardIndex := findGuardIndex(guardMap)
	visitedPositions, _ := findVisitedPositions(map[string]VisitedPosition{}, guardMap, guardIndex, directionMap["up"])

	visitedPositionsLength := len(visitedPositions)

	return visitedPositionsLength
}

func CountLoopObstructions(guardMap [][]string) int {
	guardIndex := findGuardIndex(guardMap)
	loopsDetected := map[string]bool{}
	visitedPositions, _ := findVisitedPositions(map[string]VisitedPosition{}, guardMap, guardIndex, directionMap["up"])

	for _, visitedPosition := range visitedPositions {
		for _, direction := range visitedPosition.directions {
			nextY := visitedPosition.position.y + directionMap[direction].y
			nextX := visitedPosition.position.x + directionMap[direction].x

			if nextX < 0 || nextX >= len(guardMap[0]) || nextY < 0 || nextY >= len(guardMap) {
				continue
			}

			prevValue := guardMap[nextY][nextX]

			if prevValue == "#" {
				continue
			}

			guardMap[nextY][nextX] = "#"
			_, loopDetected := findVisitedPositions(map[string]VisitedPosition{}, guardMap, guardIndex, directionMap["up"])

			if loopDetected {
				loopsDetected[fmt.Sprintf("%d,%d", nextY, nextX)] = true
			}

			guardMap[nextY][nextX] = prevValue
		}
	}

	return len(loopsDetected)
}

func main() {
	lines, _ := util.ReadFileAsArray("input.txt")
	positions := [][]string{}

	for _, line := range lines {
		positions = append(positions, strings.Split(line, ""))
	}

	part1Count := CountPositions(positions)
	part2Count := CountLoopObstructions(positions)

	fmt.Println("Result Part 1:", part1Count)
	fmt.Println("Result Part 2:", part2Count)
}
