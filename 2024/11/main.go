package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

// var mapCalcs = map[int]int{}

func SplitStringToNumbers(input string) ([]int, error) {
	numbers := strings.Split(input, " ")
	result := []int{}

	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)

		if err != nil {
			return nil, err
		}

		result = append(result, num)
	}

	return result, nil
}

func CountStones(stones []int, times int) (int, error) {
	result := map[int]int{}

	for _, num := range stones {
		result[num]++
	}

	for i := 0; i < times; i++ {
		updated := map[int]int{}

		for num, count := range result {
			// rule 1
			if num == 0 {
				updated[1] += count
				continue
			}

			// rule 2
			str := strconv.Itoa(num)
			strLength := len(str)

			if strLength%2 == 0 {
				num1, _ := strconv.Atoi(str[:strLength/2])
				num2, _ := strconv.Atoi(str[strLength/2:])

				updated[num1] += count
				updated[num2] += count
				continue
			}

			// rule 3
			updated[num*2024] += count
		}

		result = updated
	}

	sum := 0
	for _, count := range result {
		sum += count
	}

	return sum, nil
}

func main() {
	input, _ := util.ReadFileAsString("input.txt")
	numbers, _ := SplitStringToNumbers(input)
	part1, _ := CountStones(numbers, 25)

	fmt.Println("Part 1:", part1)

	part2, _ := CountStones(numbers, 75)

	fmt.Println("Part 2:", part2)
}
