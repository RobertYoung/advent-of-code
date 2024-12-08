package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x        int
	y        int
	value    string
	antinode bool
}

type Direction struct {
	y    int
	x    int
	name string
}

var directionMap = map[string]Direction{
	"up-left":    {-1, -1, "up-left"},
	"up-right":   {-1, 1, "up-right"},
	"down-left":  {1, -1, "down-left"},
	"down-right": {1, 1, "down-right"},
}

func getPosition(positions [][]Position, y int, x int) (Position, bool) {
	if y >= 0 && y < len(positions) {
		if x >= 0 && x < len(positions[y]) {
			return positions[y][x], true
		}
	}
	return Position{}, false
}

func setAntinodes(positions [][]Position, start Position, diffY int, diffX int) bool {
	positions[start.y][start.x].antinode = true

	next, found := getPosition(positions, start.y+diffY, start.x+diffX)

	if found {
		return setAntinodes(positions, next, diffY, diffX)
	}

	return false
}

func checkFrequency(positions [][]Position, a Position, b Position, inline bool) bool {
	if a.value != "." && a.value == b.value {
		diffX := b.x - a.x
		diffY := b.y - a.y

		if inline {
			// backwards
			positions[a.y][a.x].antinode = true
			if position, found := getPosition(positions, a.y-diffY, a.x-diffX); found {
				setAntinodes(positions, position, diffY, diffX)
			}

			// forwards
			positions[b.y][b.x].antinode = true
			if position, found := getPosition(positions, b.y+diffY, b.x+diffX); found {
				setAntinodes(positions, position, diffY, diffX)
			}
		}

		startPosition, startFound := getPosition(positions, a.y-diffY, a.x-diffX)

		if startFound {
			if !checkFrequency(positions, startPosition, a, inline) {
				positions[startPosition.y][startPosition.x].antinode = true
			}
		}

		endPosition, endFound := getPosition(positions, b.y+diffY, b.x+diffX)

		if endFound {
			if !checkFrequency(positions, b, endPosition, inline) {
				positions[endPosition.y][endPosition.x].antinode = true
			}
		}

		return true
	}

	return false
}

func findNextAntenna(positions [][]Position, currPosition Position, direction Direction, inline bool) ([]Position, bool) {
	totalAntinodes := []Position{}
	y := currPosition.y + direction.y

	for y >= 0 && y < len(positions) {
		x := currPosition.x + direction.x

		for x >= 0 && x < len(positions[y]) {
			checkFrequency(positions, currPosition, positions[y][x], inline)

			if x == (x + direction.x) {
				break
			}

			x += direction.x
		}

		if y == (y + direction.y) {
			break
		}

		y += direction.y
	}

	return totalAntinodes, false
}

func FindInlineAntenna(positions [][]Position, y int, x int, inline bool) Position {
	currPosition, _ := getPosition(positions, y, x)

	findNextAntenna(positions, currPosition, directionMap["up-left"], inline)
	findNextAntenna(positions, currPosition, directionMap["up-right"], inline)
	findNextAntenna(positions, currPosition, directionMap["down-left"], inline)
	findNextAntenna(positions, currPosition, directionMap["down-right"], inline)

	return Position{}
}

func FindUniqueAntinodes(positions [][]Position) (int, error) {
	for y := range positions {
		for x := range positions[y] {
			FindInlineAntenna(positions, y, x, false)
		}
	}

	antinodeCount := 0

	for y := range positions {
		for _, position := range positions[y] {
			if position.antinode {
				antinodeCount++
			}
		}
	}

	return antinodeCount, nil
}

func FindUniqueAntinodesInline(positions [][]Position) (int, error) {
	for y := range positions {
		for x := range positions[y] {
			FindInlineAntenna(positions, y, x, true)
		}
	}

	antinodeCount := 0

	for y := range positions {
		for _, position := range positions[y] {
			if position.antinode {
				antinodeCount++
			}
		}
	}

	return antinodeCount, nil
}

func ConvertToPositions(input string) [][]Position {
	lines := strings.Split(input, "\n")
	positions := [][]Position{}

	for y, line := range lines {
		positions = append(positions, []Position{})

		for x, character := range strings.Split(line, "") {
			positions[y] = append(positions[y], Position{x: x, y: y, value: character, antinode: false})
		}
	}

	return positions
}

func main() {
	bytes, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	file := string(bytes)
	part1, _ := FindUniqueAntinodes(ConvertToPositions(file))
	part2, _ := FindUniqueAntinodesInline(ConvertToPositions(file))

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
