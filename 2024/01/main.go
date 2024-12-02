package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	rightCount := make(map[int]int)
	for _, val := range right {
		rightCount[val]++
	}

	var distances []int
	var similarities []int

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		if distance < 0 {
			distance = -distance
		}
		distances = append(distances, distance)

		similarity := rightCount[left[i]]
		similarities = append(similarities, left[i]*similarity)
	}

	totalDistance := 0
	for _, distance := range distances {
		totalDistance += distance
	}

	totalSimilarity := 0
	for _, similarity := range similarities {
		totalSimilarity += similarity
	}

	fmt.Println("Total distance:", totalDistance)
	fmt.Println("Total similarity:", totalSimilarity)
}
