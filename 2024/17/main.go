package main

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

type Register struct {
	a int64
	b int64
	c int64
}

type Game struct {
	register Register
	program  []int64
}

func (game *Game) getComboOperandValue(value int64) int64 {
	switch value {
	case 0, 1, 2, 3:
		return value
	case 4:
		return game.register.a
	case 5:
		return game.register.b
	case 6:
		return game.register.c
	case 7:
		return -1
	}

	return -1
}

func (game *Game) findCopy() int64 {
	result := int64(0)
	index := int64(1)

	for result == 0 {
		game.register.a = index
		game.register.b = 0
		game.register.c = 0

		out := game.execute()

		if len(out) == len(game.program) {

			if reflect.DeepEqual(out, game.program) {
				result = index
				continue
			}

			for i := len(out) - 1; i >= 0; i-- {
				if out[i] != game.program[i] {
					index += int64(math.Pow(8, float64(i)))
					break
				}
			}
			continue
		}

		index <<= 1
	}

	return result
}

func (game *Game) execute() []int64 {
	result := []int64{}
	instructionPointer := int64(0)

	for instructionPointer < int64(len(game.program)) {
		opcode := game.program[instructionPointer]
		operand := game.program[instructionPointer+1]

		increaseInstructionPointer := true

		switch opcode {
		case 0:
			game.register.a = game.register.a / int64(math.Pow(2, float64(game.getComboOperandValue(operand))))
		case 1:
			game.register.b = game.register.b ^ operand
		case 2:
			game.register.b = game.getComboOperandValue(operand) % 8
		case 3:
			if game.register.a != 0 {
				instructionPointer = operand
				increaseInstructionPointer = false
				continue
			}
		case 4:
			game.register.b = game.register.b ^ game.register.c
		case 5:
			result = append(result, game.getComboOperandValue(operand)%8)
		case 6:
			game.register.b = game.register.a / int64(math.Pow(2, float64(game.getComboOperandValue(operand))))
		case 7:
			game.register.c = game.register.a / int64(math.Pow(2, float64(game.getComboOperandValue(operand))))
		}

		if increaseInstructionPointer {
			instructionPointer += 2
		}
	}

	return result
}

func CreateGame(input string) Game {
	game := Game{
		register: Register{},
		program:  []int64{},
	}

	for _, line := range strings.Split(input, "\n") {
		if _, err := fmt.Sscanf(line, "Register A: %d", &game.register.a); err == nil {
			continue
		}
		if _, err := fmt.Sscanf(line, "Register B: %d", &game.register.b); err == nil {
			continue
		}
		if _, err := fmt.Sscanf(line, "Register C: %d", &game.register.c); err == nil {
			continue
		}

		matches := regexp.MustCompile(`^Program: (.*)$`).FindStringSubmatch(line)

		if len(matches) > 1 {
			programStr := matches[1]
			for _, numStr := range strings.Split(programStr, ",") {
				num, err := strconv.Atoi(strings.TrimSpace(numStr))
				if err == nil {
					game.program = append(game.program, int64(num))
				}
			}
		}
	}

	return game
}

func main() {
	input, _ := util.ReadFileAsString("input.txt")
	game := CreateGame(input)
	result := game.execute()

	part1 := []string{}

	for _, num := range result {
		part1 = append(part1, strconv.FormatInt(num, 10))
	}

	fmt.Println("Part 1:", strings.Join(part1, ","))

	part2 := game.findCopy()

	fmt.Println("Part 2:", part2)
}
