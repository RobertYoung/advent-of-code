package main

import (
	"fmt"

	"github.com/RobertYoung/advent-of-code/util"
)

func isInbound(x int, y int, maxX int, maxY int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}

func setAntinodes(antinodes map[[2]int]bool, dx int, dy int, x int, y int, maxX int, maxY int) map[[2]int]bool {
	antinodes[[2]int{x, y}] = true

	if isInbound(x-dx, y-dy, maxX, maxY) {
		antinodes[[2]int{x - dx, y - dy}] = true
		setAntinodes(antinodes, dx, dy, x-dx, y-dy, maxX, maxY)
	}

	return antinodes
}

func GetAntinodes(antennas map[rune][][2]int, maxX int, maxY int) (map[[2]int]bool, map[[2]int]bool) {
	antinodesPart1 := make(map[[2]int]bool)
	antinodesPart2 := make(map[[2]int]bool)

	for _, positions := range antennas {
		n := len(positions)

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				p1, p2 := positions[i], positions[j]
				dx, dy := p2[0]-p1[0], p2[1]-p1[1]

				if isInbound(p1[0]-dx, p1[1]-dy, maxX, maxY) {
					antinodesPart1[[2]int{p1[0] - dx, p1[1] - dy}] = true
				}

				if isInbound(p2[0]+dx, p2[1]+dy, maxX, maxY) {
					antinodesPart1[[2]int{p2[0] + dx, p2[1] + dy}] = true
				}

				setAntinodes(antinodesPart2, dx, dy, p1[0], p1[1], maxX, maxY)
				setAntinodes(antinodesPart2, -dx, -dy, p2[0], p2[1], maxX, maxY)
			}
		}
	}

	return antinodesPart1, antinodesPart2
}

func ConvertToAntennas(grid []string) map[rune][][2]int {
	antennas := make(map[rune][][2]int)

	for y, row := range grid {
		for x, char := range row {
			if char != '.' {
				antennas[char] = append(antennas[char], [2]int{x, y})
			}
		}
	}

	return antennas
}

func main() {
	grid, _ := util.ReadFileAsArray("input.txt")
	antennas := ConvertToAntennas(grid)
	part1, part2 := GetAntinodes(antennas, len(grid[0]), len(grid))

	fmt.Println("Part 1", len(part1))
	fmt.Println("Part 2", len(part2))
}
