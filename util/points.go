package util

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func ConvertToPoints[T any](input string, convert func(string) (T, error)) (map[Point]T, error) {
	result := map[Point]T{}

	for y, row := range strings.Split(input, "\n") {
		for x, char := range strings.Split(row, "") {
			value, err := convert(char)
			if err != nil {
				return nil, err
			}
			result[Point{X: x, Y: y}] = value
		}
	}

	return result, nil
}

func ConvertToInt(char string) (int, error) {
	return strconv.Atoi(char)
}

func ConvertToString(char string) (string, error) {
	return char, nil
}
