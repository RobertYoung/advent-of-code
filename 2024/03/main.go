package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func FixCorruptedLine(line string) [][]int {
	re := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
	matches := re.FindAllStringSubmatch(line, -1)
	var result [][]int

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		result = append(result, []int{x, y})
	}

	return result
}

func FixCorruptedLineConditional(line string) [][]int {
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don\'t\(\)`)
	characters := strings.Split(line, "")
	section := ""
	enabled := true
	var result = [][]int{}

	for _, character := range characters {
		section += character
		dontMatch := dontRegex.FindStringSubmatch(section)

		if dontMatch != nil {
			enabled = false
			section = ""
			continue
		}

		doMatch := doRegex.FindStringSubmatch(section)

		if doMatch != nil {
			enabled = true
			section = ""
			continue
		}

		mulMatch := mulRegex.FindStringSubmatch(section)

		if mulMatch != nil {
			if !enabled {
				section = ""
				continue
			}

			x, _ := strconv.Atoi(mulMatch[1])
			y, _ := strconv.Atoi(mulMatch[2])
			result = append(result, []int{x, y})
			section = ""
			continue
		}
	}

	return result
}

func main() {
	bytes, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var valuesPart1 [][]int
	var valuesPart2 [][]int

	file := string(bytes)
	valuesPart1 = append(valuesPart1, FixCorruptedLine(file)...)
	valuesPart2 = append(valuesPart2, FixCorruptedLineConditional(file)...)

	resultPart1 := 0
	for _, value := range valuesPart1 {
		resultPart1 += value[0] * value[1]
	}

	resultPart2 := 0
	for _, value := range valuesPart2 {
		resultPart2 += value[0] * value[1]
	}

	fmt.Println("Result Part 1:", resultPart1)
	fmt.Println("Result Part 2:", resultPart2)
}
