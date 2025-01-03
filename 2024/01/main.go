package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculateSimilarity(left []int, right []int) int {
	rightCount := make(map[int]int)

	for _, val := range right {
		rightCount[val]++
	}

	var similarities []int

	for i := 0; i < len(left); i++ {
		similarity := rightCount[left[i]]
		similarities = append(similarities, left[i]*similarity)
	}

	total := 0

	for _, similarity := range similarities {
		total += similarity
	}

	return total
}

func CalculateDistance(left []int, right []int) int {
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var distances []int

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]

		if distance < 0 {
			distance = -distance
		}

		distances = append(distances, distance)
	}

	total := 0

	for _, distance := range distances {
		total += distance
	}

	return total
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		values := strings.Fields(line)
		if len(values) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		leftVal, err1 := strconv.Atoi(values[0])
		rightVal, err2 := strconv.Atoi(values[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting values:", values)
			continue
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	totalDistance := CalculateDistance(left, right)
	totalSimilarity := CalculateSimilarity(left, right)

	fmt.Println("Total distance:", totalDistance)
	fmt.Println("Total similarity:", totalSimilarity)
}
