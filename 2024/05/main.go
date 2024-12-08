package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

type PageRule struct {
	x int
	y int
}

func findInvalidPages(ruleset []PageRule, num int) []int {
	var invalidPages []int
	for _, rule := range ruleset {
		if rule.y == num {
			invalidPages = append(invalidPages, rule.x)
		}
	}
	return invalidPages
}

func ReorderPages(rules []PageRule, pages []int) []int {
	var invalidPages []int
	for pageIndex, pageNum := range pages {
		fmt.Println("check page num:", pageNum)

		if slices.Contains(invalidPages, pageNum) {
			pages[pageIndex-1], pages[pageIndex] = pages[pageIndex], pages[pageIndex-1]
			return ReorderPages(rules, pages)
		} else {
			invalidPages = append(invalidPages, findInvalidPages(rules, pageNum)...)
		}
	}
	return pages
}

func IsValidPages(rules []PageRule, pages []int) bool {
	valid := true

	for _, rule := range rules {
		indexX := -1
		indexY := -1

		for pageIndex, page := range pages {
			if page == rule.x {
				indexX = pageIndex
			}
			if page == rule.y {
				indexY = pageIndex
			}

			if indexX >= 0 && indexY >= 0 {
				if indexX >= indexY {
					valid = false
					break
				}
			}
		}

		if !valid {
			break
		}
	}

	return valid
}

func main() {
	file, _ := util.ReadFileAsString("input.txt")
	sectionsStr := strings.Split(file, "\n\n")
	rulesStr := strings.Split(sectionsStr[0], "\n")
	pagesStr := strings.Split(sectionsStr[1], "\n")

	var pageRules []PageRule

	for _, rule := range rulesStr {
		parts := strings.Split(rule, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		pageRule := PageRule{x: x, y: y}
		pageRules = append(pageRules, pageRule)
	}

	part1Count := 0
	part2Count := 0

	for _, page := range pagesStr {
		parts := strings.Split(page, ",")
		var pages []int

		for _, part := range parts {
			page, _ := strconv.Atoi(part)
			pages = append(pages, page)
		}

		if IsValidPages(pageRules, pages) {
			part1Count += pages[(len(pages)+1)/2-1]
		} else {
			reorder := ReorderPages(pageRules, pages)
			part2Count += reorder[(len(reorder)+1)/2-1]
		}
	}

	fmt.Println("Result Part 1:", part1Count)
	fmt.Println("Result Part 2:", part2Count)
}
