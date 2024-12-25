package main

import (
	"container/heap"
	"fmt"
	"math"
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
	EMPTY MapKey = "."
	WALL  MapKey = "#"
	START MapKey = "S"
	END   MapKey = "E"
)

type Game struct {
	points map[util.Point]MapKey
	bound  util.Point
	start  util.Point
	end    util.Point
}

type State struct {
	point     util.Point
	direction util.Direction
	score     int
	previous  *State
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(points map[util.Point]MapKey, start, end util.Point, startDirection DirectionKey) (int, int) {
	dist := map[util.Point]map[util.Direction]int{}

	for point := range points {
		dist[point] = map[util.Direction]int{}

		for _, direction := range DirectionMap {
			dist[point][direction] = math.MaxInt32
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for dirKey, direction := range DirectionMap {
		if dirKey != startDirection {
			heap.Push(pq, &State{start, direction, 1000, nil})
		} else {
			heap.Push(pq, &State{start, direction, 0, nil})
		}
		dist[start][direction] = 0
	}

	bestScore := math.MaxInt32
	bestTiles := map[util.Point]bool{}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)

		if current.point == end {
			if current.score < bestScore {
				bestScore = current.score
				bestTiles = map[util.Point]bool{}
			}

			if current.score == bestScore {
				tile := current

				for tile != nil {
					bestTiles[tile.point] = true
					tile = tile.previous
				}
			}

			continue
		}

		for _, direction := range DirectionMap {
			newScore := current.score + 1000

			if newScore <= dist[current.point][direction] {
				dist[current.point][direction] = newScore
				heap.Push(pq, &State{current.point, direction, newScore, current})
			}
		}

		nextPoint := util.Point{
			X: current.point.X + current.direction.X,
			Y: current.point.Y + current.direction.Y,
		}

		if points[nextPoint] != WALL {
			newScore := current.score + 1

			if newScore <= dist[nextPoint][current.direction] {
				dist[nextPoint][current.direction] = newScore
				heap.Push(pq, &State{nextPoint, current.direction, newScore, current})
			}
		}
	}

	if bestScore == math.MaxInt32 {
		bestScore = -1
	}

	return bestScore, len(bestTiles)
}

func (game *Game) findShortestRoute() (int, int) {
	return dijkstra(game.points, game.start, game.end, RIGHT)
}

func CreateGame(input string) Game {
	game := Game{
		points: make(map[util.Point]MapKey),
	}

	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.Split(line, "") {
			point := util.Point{X: x, Y: y}

			if MapKey(char) == START {
				game.start = point
			} else if MapKey(char) == END {
				game.end = point
			}

			game.points[point] = MapKey(char)
			game.bound = point
		}
	}

	return game
}

func main() {
	input, _ := util.ReadFileAsString("input.txt")
	part1Game := CreateGame(input)
	part1, part2 := part1Game.findShortestRoute()

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
