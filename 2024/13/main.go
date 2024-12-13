package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

type Button struct {
	x    float64
	y    float64
	cost int
}

type Prize struct {
	x float64
	y float64
}

type Machine struct {
	a     Button
	b     Button
	prize Prize
}

func getXAndY(input string) (float64, float64) {
	re := regexp.MustCompile(`X[+=](\d+), Y[+=](\d+)`)
	matches := re.FindStringSubmatch(input)
	var x, y int

	if len(matches) == 3 {
		x, _ = strconv.Atoi(matches[1])
		y, _ = strconv.Atoi(matches[2])
	}
	return float64(x), float64(y)
}

func ConvertInput(input string, addToPrice float64) ([]Machine, error) {
	result := []Machine{}
	next := Machine{}

	for _, line := range strings.Split(input, "\n") {
		x, y := getXAndY(line)

		if strings.HasPrefix(line, "Button A") {
			next = Machine{}
			next.a = Button{x, y, 3}
		} else if strings.HasPrefix(line, "Button B") {
			next.b = Button{x, y, 1}
		} else if strings.HasPrefix(line, "Prize") {
			next.prize = Prize{x + addToPrice, y + addToPrice}
			result = append(result, next)
		}
	}

	return result, nil
}

func (machine *Machine) calculateTokens() (int, error) {
	result := -1
	a := (machine.prize.x*machine.b.y - machine.b.x*machine.prize.y) / (machine.a.x*machine.b.y - machine.b.x*machine.a.y)
	b := (machine.prize.x - machine.a.x*a) / machine.b.x

	if math.Mod(a, 1) == 0 && math.Mod(b, 1) == 0 {
		result = int(a)*machine.a.cost + int(b)*machine.b.cost
	} else {
		return -1, fmt.Errorf("invalid tokens")
	}

	return result, nil
}

func main() {
	input, _ := util.ReadFileAsString("input.txt")
	part1Machines, _ := ConvertInput(input, 0)
	part1 := 0

	for _, machine := range part1Machines {
		if tokens, err := machine.calculateTokens(); err == nil {
			part1 += tokens
		}
	}

	fmt.Println("Part 1:", part1)

	part2Machines, _ := ConvertInput(input, 10000000000000)
	part2 := 0

	for _, machine := range part2Machines {
		if tokens, err := machine.calculateTokens(); err == nil {
			part2 += tokens
		}
	}

	fmt.Println("Part 2:", part2)
}
