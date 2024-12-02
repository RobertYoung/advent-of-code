package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AllValuesSame(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	firstValue := arr[0]
	for _, value := range arr {
		if value != firstValue {
			return false
		}
	}
	return true
}

func CheckIfSafe(arr []int) bool {
	var diffs []int
	var isSafe bool = false

	for i, value := range arr {

		if i == len(arr)-1 {
			if AllValuesSame(diffs) {
				isSafe = true
			}
			break
		}

		var diff = arr[i+1] - value

		if diff >= 1 && diff <= 3 {
			diffs = append(diffs, 1)
		} else if diff <= -1 && diff >= -3 {
			diffs = append(diffs, -1)
		} else {
			return false
		}
	}

	return isSafe
}

func RemoveAtIndex(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

func IsSafeReportWithTolerance(values []int) bool {
	var isSafe bool = false

	if CheckIfSafe(values) {
		isSafe = true
	} else {
		for i, _ := range values {
			copyValues := make([]int, len(values))
			copy(copyValues, values)

			copyValues = RemoveAtIndex(copyValues, i)

			if CheckIfSafe(copyValues) {
				isSafe = true
				break
			}
		}
	}

	return isSafe
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	var safeReports []bool
	var safeReportsWithTolerance []bool

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		var values []int

		for _, value := range strings.Fields(line) {
			intValue, _ := strconv.Atoi(value)
			values = append(values, intValue)
		}

		if CheckIfSafe(values) {
			safeReports = append(safeReports, true)
		}

		if IsSafeReportWithTolerance(values) {
			safeReportsWithTolerance = append(safeReportsWithTolerance, true)
		}
	}

	fmt.Println("Safe reports:", len(safeReports))
	fmt.Println("Safe reports with tolerance:", len(safeReportsWithTolerance))
}
