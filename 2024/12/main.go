package main

import (
	"fmt"

	"github.com/RobertYoung/advent-of-code/util"
)

type Region struct {
	id int
}

type Plot struct {
	region Region
	value  string
}

func convertPlotsToRegions(plots map[util.Point]Plot) map[Region]map[util.Point]Plot {
	regions := map[Region]map[util.Point]Plot{}

	for point, plot := range plots {
		if _, exists := regions[plot.region]; !exists {
			regions[plot.region] = map[util.Point]Plot{}
		}
		regions[plot.region][point] = plot
	}

	return regions
}

func findPlots(plots map[util.Point]Plot, points map[util.Point]string, point util.Point, region Region) {
	plots[point] = Plot{
		region: region,
		value:  points[point],
	}

	up := util.Point{X: point.X + util.DirectionMap["up"].X, Y: point.Y + util.DirectionMap["up"].Y}
	down := util.Point{X: point.X + util.DirectionMap["down"].X, Y: point.Y + util.DirectionMap["down"].Y}
	left := util.Point{X: point.X + util.DirectionMap["left"].X, Y: point.Y + util.DirectionMap["left"].Y}
	right := util.Point{X: point.X + util.DirectionMap["right"].X, Y: point.Y + util.DirectionMap["right"].Y}

	if _, exists := plots[up]; !exists && points[up] == points[point] {
		findPlots(plots, points, up, region)
	}

	if _, exists := plots[down]; !exists && points[down] == points[point] {
		findPlots(plots, points, down, region)
	}

	if _, exists := plots[left]; !exists && points[left] == points[point] {
		findPlots(plots, points, left, region)
	}

	if _, exists := plots[right]; !exists && points[right] == points[point] {
		findPlots(plots, points, right, region)
	}
}

func FindPlotPoints(points map[util.Point]string) (map[util.Point]Plot, error) {
	regionIndex := 0
	regions := map[int]Region{}
	plots := map[util.Point]Plot{}

	for point := range points {
		if _, exists := plots[point]; exists {
			continue
		}

		region := Region{id: regionIndex}
		regions[regionIndex] = region
		findPlots(plots, points, point, region)
		regionIndex++
	}

	return plots, nil
}

func CalculateFenceCost(plots map[util.Point]Plot) (int, error) {
	regions := convertPlotsToRegions(plots)
	sum := 0

	for _, plots := range regions {
		area := len(plots)
		perimeter := area * 4

		for point := range plots {
			for _, direction := range util.DirectionMap {
				j := util.Point{X: point.X + direction.X, Y: point.Y + direction.Y}

				if _, exists := plots[j]; exists {
					perimeter--
				}
			}
		}

		sum += area * perimeter
	}

	return sum, nil
}

func CalculateFenceCostWithDiscount(plots map[util.Point]Plot) (int, error) {
	regions := convertPlotsToRegions(plots)
	sum := 0

	for _, plots := range regions {
		area := len(plots)
		perimeter := 0

		for point := range plots {
			_, up := plots[util.Point{
				X: point.X + util.DirectionMap["up"].X,
				Y: point.Y + util.DirectionMap["up"].Y,
			}]

			_, down := plots[util.Point{
				X: point.X + util.DirectionMap["down"].X,
				Y: point.Y + util.DirectionMap["down"].Y,
			}]

			_, left := plots[util.Point{
				X: point.X + util.DirectionMap["left"].X,
				Y: point.Y + util.DirectionMap["left"].Y,
			}]

			_, right := plots[util.Point{
				X: point.X + util.DirectionMap["right"].X,
				Y: point.Y + util.DirectionMap["right"].Y,
			}]

			_, upLeft := plots[util.Point{
				X: point.X + util.DirectionMap["up"].X + util.DirectionMap["left"].X,
				Y: point.Y + util.DirectionMap["up"].Y + util.DirectionMap["left"].Y,
			}]

			_, upRight := plots[util.Point{
				X: point.X + util.DirectionMap["up"].X + util.DirectionMap["right"].X,
				Y: point.Y + util.DirectionMap["up"].Y + util.DirectionMap["right"].Y,
			}]

			_, downLeft := plots[util.Point{
				X: point.X + util.DirectionMap["down"].X + util.DirectionMap["left"].X,
				Y: point.Y + util.DirectionMap["down"].Y + util.DirectionMap["left"].Y,
			}]

			_, downRight := plots[util.Point{
				X: point.X + util.DirectionMap["down"].X + util.DirectionMap["right"].X,
				Y: point.Y + util.DirectionMap["down"].Y + util.DirectionMap["right"].Y,
			}]

			// top left corner
			if !up && !left {
				perimeter++
			}

			// top right corner
			if !up && !right {
				perimeter++
			}

			// bottom left corner
			if !down && !left {
				perimeter++
			}

			// bottom right corner
			if !down && !right {
				perimeter++
			}

			// inside top left
			if up && left && !upLeft {
				perimeter++
			}

			// inside top right
			if up && right && !upRight {
				perimeter++
			}

			// inside bottom left
			if down && left && !downLeft {
				perimeter++
			}

			// inside bottom right
			if down && right && !downRight {
				perimeter++
			}
		}

		sum += area * perimeter
	}

	return sum, nil
}

func main() {
	input, _ := util.ReadFileAsString("input.txt")
	points, _ := util.ConvertToPoints(input, util.ConvertToString)
	plots, _ := FindPlotPoints(points)
	part1, _ := CalculateFenceCost(plots)

	fmt.Println("Part 1:", part1)

	part2, _ := CalculateFenceCostWithDiscount(plots)

	fmt.Println("Part 2:", part2)
}
