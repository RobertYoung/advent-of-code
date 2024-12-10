package main

import (
	"fmt"

	"github.com/RobertYoung/advent-of-code/util"
)

func findTrails(trails map[util.Point]int, points map[util.Point]int, curr util.Point) {
	trailIndex := points[curr]
	nextTrailIndex := trailIndex + 1

	if trailIndex == 9 {
		trails[curr] = trailIndex
		return
	}

	upPoint := util.Point{X: curr.X + util.DirectionMap["up"].X, Y: curr.Y + util.DirectionMap["up"].Y}
	downPoint := util.Point{X: curr.X + util.DirectionMap["down"].X, Y: curr.Y + util.DirectionMap["down"].Y}
	leftPoint := util.Point{X: curr.X + util.DirectionMap["left"].X, Y: curr.Y + util.DirectionMap["left"].Y}
	rightPoint := util.Point{X: curr.X + util.DirectionMap["right"].X, Y: curr.Y + util.DirectionMap["right"].Y}

	if points[upPoint] == nextTrailIndex {
		findTrails(trails, points, upPoint)
	}

	if points[downPoint] == nextTrailIndex {
		findTrails(trails, points, downPoint)
	}

	if points[leftPoint] == nextTrailIndex {
		findTrails(trails, points, leftPoint)
	}

	if points[rightPoint] == nextTrailIndex {
		findTrails(trails, points, rightPoint)
	}
}

func findDistinctTrails(trails *int, points map[util.Point]int, curr util.Point) {
	trailIndex := points[curr]
	nextTrailIndex := trailIndex + 1

	if trailIndex == 9 {
		*trails++
		return
	}

	upPoint := util.Point{X: curr.X + util.DirectionMap["up"].X, Y: curr.Y + util.DirectionMap["up"].Y}
	downPoint := util.Point{X: curr.X + util.DirectionMap["down"].X, Y: curr.Y + util.DirectionMap["down"].Y}
	leftPoint := util.Point{X: curr.X + util.DirectionMap["left"].X, Y: curr.Y + util.DirectionMap["left"].Y}
	rightPoint := util.Point{X: curr.X + util.DirectionMap["right"].X, Y: curr.Y + util.DirectionMap["right"].Y}

	if points[upPoint] == nextTrailIndex {
		findDistinctTrails(trails, points, upPoint)
	}

	if points[downPoint] == nextTrailIndex {
		findDistinctTrails(trails, points, downPoint)
	}

	if points[leftPoint] == nextTrailIndex {
		findDistinctTrails(trails, points, leftPoint)
	}

	if points[rightPoint] == nextTrailIndex {
		findDistinctTrails(trails, points, rightPoint)
	}
}

func FindNumberOfTrails(points map[util.Point]int) (int, error) {
	count := 0

	for point, value := range points {
		trails := map[util.Point]int{}

		if value == 0 {
			findTrails(trails, points, point)
		}

		count += len(trails)
	}

	return count, nil
}

func FindNumberOfDistinctTrails(points map[util.Point]int) (int, error) {
	count := 0

	for point, value := range points {
		trails := 0

		if value == 0 {
			findDistinctTrails(&trails, points, point)
		}

		count += trails
	}

	return count, nil
}

func main() {
	input, _ := util.ReadFileAsString("input.txt")
	points, _ := util.ConvertToPoints(input)
	part1, _ := FindNumberOfTrails(points)
	part2, _ := FindNumberOfDistinctTrails(points)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
