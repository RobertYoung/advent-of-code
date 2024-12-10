package util

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func ConvertToPoints(input string) (map[Point]int, error) {
	result := map[Point]int{}

	for y, row := range strings.Split(input, "\n") {
		for x, char := range strings.Split(row, "") {
			value, err := strconv.Atoi(char)
			if err == nil {
				result[Point{X: x, Y: y}] = value
			} else {
				result[Point{X: x, Y: y}] = -1
			}
		}
	}

	return result, nil
}
