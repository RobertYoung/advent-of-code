package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Equation struct {
	total  int
	values []int
}

func (e Equation) IsValidPart1() bool {
	if len(e.values) == 0 {
		return false
	}

	return checkCombination(e.values, e.total, e.values[0], 1, false)
}

func (e Equation) IsValidPart2() bool {
	if len(e.values) == 0 {
		return false
	}

	return checkCombination(e.values, e.total, e.values[0], 1, true)
}

func checkCombination(values []int, target int, current int, index int, concatenation bool) bool {
	if index == len(values) {
		return current == target
	}

	if checkCombination(values, target, current+values[index], index+1, concatenation) {
		return true
	}

	if checkCombination(values, target, current*values[index], index+1, concatenation) {
		return true
	}

	if concatenation {
		concatInt, _ := strconv.Atoi(strconv.Itoa(current) + strconv.Itoa(values[index]))
		concatValues := []int{concatInt}
		concatValues = append(concatValues, values[index+1:]...)

		if checkCombination(concatValues, target, concatValues[0], 1, concatenation) {
			return true
		}
	}

	return false
}

func ConvertEquation(line string) (Equation, error) {
	parse := regexp.MustCompile(`\d+`).FindAllString(line, -1)
	total, totalErr := strconv.Atoi(parse[0])

	if totalErr != nil {
		return Equation{}, fmt.Errorf("error converting total for line: %s", line)
	}

	values := []int{}

	for _, value := range parse[1:] {
		val, err := strconv.Atoi(value)
		if err != nil {
			return Equation{}, fmt.Errorf("error converting value for line: %s", line)
		}
		values = append(values, val)
	}

	return Equation{
		total:  total,
		values: values,
	}, nil
}

func main() {
	bytes, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	file := string(bytes)
	lines := strings.Split(file, "\n")
	part1 := 0
	part2 := 0

	for _, line := range lines {
		equation, _ := ConvertEquation(line)

		if equation.IsValidPart1() {
			part1 += equation.total
		}
		if equation.IsValidPart2() {
			part2 += equation.total
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
