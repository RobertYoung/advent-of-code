package main

import (
	"fmt"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

type DirectionKey string

const (
	UP    DirectionKey = "^"
	DOWN  DirectionKey = "v"
	LEFT  DirectionKey = "<"
	RIGHT DirectionKey = ">"
)

var DirectionMap = map[DirectionKey]util.Direction{
	UP:    {Y: -1, X: 0, Name: "up"},
	DOWN:  {Y: 1, X: 0, Name: "down"},
	LEFT:  {Y: 0, X: -1, Name: "left"},
	RIGHT: {Y: 0, X: 1, Name: "right"},
}

type MapKey string

const (
	WALL      MapKey = "#"
	BOX       MapKey = "O"
	ROBOT     MapKey = "@"
	EMPTY     MapKey = "."
	LEFT_BOX  MapKey = "["
	RIGHT_BOX MapKey = "]"
)

type Game struct {
	warehouse map[util.Point]MapKey
	moves     []DirectionKey
	bound     util.Point
}

func (game *Game) movePart1() {
	robot := util.Point{}

	for point, key := range game.warehouse {
		if key == ROBOT {
			robot = point
			break
		}
	}

	for _, move := range game.moves {
		direction := DirectionMap[move]
		next := util.Point{X: robot.X + direction.X, Y: robot.Y + direction.Y}

		if game.warehouse[next] == WALL {
			continue
		}

		if game.warehouse[next] == EMPTY {
			game.warehouse[next] = ROBOT
			game.warehouse[robot] = EMPTY
			robot = next
			continue
		}

		if game.warehouse[next] == BOX {
			done := false
			index := 1

			for !done {
				nextBox := util.Point{X: next.X + (direction.X * index), Y: next.Y + (direction.Y * index)}

				if game.warehouse[nextBox] == WALL {
					done = true
				} else if game.warehouse[nextBox] == EMPTY {
					for i := index; i > 0; i-- {
						game.warehouse[util.Point{X: next.X + (direction.X * i), Y: next.Y + (direction.Y * i)}] = BOX
						game.warehouse[util.Point{X: next.X + (direction.X * (i - 1)), Y: next.Y + (direction.Y * (i - 1))}] = EMPTY
					}
					game.warehouse[next] = ROBOT
					game.warehouse[robot] = EMPTY
					robot = next
					done = true
				} else if game.warehouse[nextBox] == BOX {
					index++
				}
			}
			continue
		}
	}
}

func pointsToMove(current util.Point, index int, direction util.Direction, warehouse map[util.Point]MapKey, points map[util.Point]MapKey) map[util.Point]MapKey {
	next := util.Point{X: current.X + direction.X, Y: current.Y + direction.Y}

	points[next] = warehouse[current]

	if index == 0 && points[current] == "" {
		points[current] = EMPTY
	}

	if warehouse[next] == WALL {
		return map[util.Point]MapKey{}
	}

	if warehouse[next] == EMPTY {
		points[next] = warehouse[current]
		return points
	}

	if warehouse[next] == LEFT_BOX && direction.Y != 0 {
		moves := pointsToMove(
			util.Point{X: next.X + 1, Y: next.Y},
			0,
			direction,
			warehouse,
			points,
		)

		if len(moves) == 0 {
			return map[util.Point]MapKey{}
		}
	}

	if warehouse[next] == RIGHT_BOX && direction.Y != 0 {
		moves := pointsToMove(
			util.Point{X: next.X - 1, Y: next.Y},
			0,
			direction,
			warehouse,
			points,
		)

		if len(moves) == 0 {
			return map[util.Point]MapKey{}
		}
	}

	return pointsToMove(
		next,
		index+1,
		direction,
		warehouse,
		points,
	)
}

func (game *Game) movePart2() {
	robot := util.Point{}

	for point, key := range game.warehouse {
		if key == ROBOT {
			robot = point
			break
		}
	}

	for _, move := range game.moves {
		direction := DirectionMap[move]
		next := util.Point{X: robot.X + direction.X, Y: robot.Y + direction.Y}

		if game.warehouse[next] == WALL {
			continue
		}

		if game.warehouse[next] == EMPTY {
			game.warehouse[next] = ROBOT
			game.warehouse[robot] = EMPTY
			robot = next
			continue
		}

		if game.warehouse[next] == RIGHT_BOX || game.warehouse[next] == LEFT_BOX {
			x := pointsToMove(robot, 0, direction, game.warehouse, map[util.Point]MapKey{})

			for point, value := range x {
				game.warehouse[point] = value
			}

			if len(x) > 0 {
				robot = next
			}

			continue
		}
	}
}

func (game *Game) calculateGPS() int {
	result := 0

	for point, key := range game.warehouse {
		if key == BOX || key == LEFT_BOX {
			result += 100*point.Y + point.X
		}
	}

	return result
}

func (game *Game) print() {
	for y := 0; y <= game.bound.Y; y++ {
		for x := 0; x <= game.bound.X; x++ {
			fmt.Print(game.warehouse[util.Point{X: x, Y: y}])
		}
		fmt.Println()
	}
}

func CreateGamePart1(input string) Game {
	split := strings.Split(input, "\n\n")
	warehouseStr := split[0]
	movesStr := split[1]

	game := Game{
		warehouse: make(map[util.Point]MapKey),
		moves:     []DirectionKey{},
	}

	for y, line := range strings.Split(warehouseStr, "\n") {
		for x, char := range strings.Split(line, "") {
			game.warehouse[util.Point{X: x, Y: y}] = MapKey(char)
			game.bound = util.Point{X: x, Y: y}
		}
	}

	for _, char := range strings.Split(movesStr, "") {
		game.moves = append(game.moves, DirectionKey(char))
	}

	return game
}

func CreateGamePart2(input string) Game {
	split := strings.Split(input, "\n\n")
	warehouseStr := split[0]
	movesStr := split[1]

	game := Game{
		warehouse: make(map[util.Point]MapKey),
		moves:     []DirectionKey{},
	}

	x := 0
	y := 0

	for _, line := range strings.Split(warehouseStr, "\n") {
		for _, char := range strings.Split(line, "") {
			if MapKey(char) == ROBOT {
				game.warehouse[util.Point{X: x, Y: y}] = ROBOT
				x++
				game.warehouse[util.Point{X: x, Y: y}] = EMPTY
				x++
			} else if MapKey(char) == BOX {
				game.warehouse[util.Point{X: x, Y: y}] = LEFT_BOX
				x++
				game.warehouse[util.Point{X: x, Y: y}] = RIGHT_BOX
				x++
			} else {
				game.warehouse[util.Point{X: x, Y: y}] = MapKey(char)
				x++
				game.warehouse[util.Point{X: x, Y: y}] = MapKey(char)
				x++
			}

			game.bound = util.Point{X: x, Y: y}
		}
		x = 0
		y++
	}

	for _, char := range strings.Split(movesStr, "") {
		game.moves = append(game.moves, DirectionKey(char))
	}

	return game
}

func main() {
	input, _ := util.ReadFileAsString("input.txt")
	part1Game := CreateGamePart1(input)
	part1Game.movePart1()
	part1 := part1Game.calculateGPS()

	fmt.Println("Part 1:", part1)

	part2Game := CreateGamePart2(input)
	part2Game.movePart2()
	part2 := part2Game.calculateGPS()

	fmt.Println("Part 2:", part2)
}
