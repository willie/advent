package main

import (
	"fmt"
	"image"

	"github.com/willie/advent/aoc"
)

func drawSensor(grid aoc.Grid2[string], origin image.Point, distance int, row int) {
	for x := -distance; x <= distance; x++ {
		for y := -distance; y <= distance; y++ {
			dest := origin.Add(image.Pt(x, y))

			if dest.Y != row {
				continue
			}

			if aoc.ManhattanDistancePt(origin, dest) <= distance {
				if _, ok := grid[dest]; !ok {
					grid[dest] = "#"
				}
			}
		}
	}
}

func part1(name string, row int) {
	grid := aoc.Grid2[string]{}

	for _, s := range aoc.Strings(name) {
		var sensor, beacon image.Point
		fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)

		grid[sensor], grid[beacon] = "S", "B"
		distance := aoc.ManhattanDistancePt(sensor, beacon)
		if (sensor.Y-distance <= row) && (row <= sensor.Y+distance) {
			drawSensor(grid, sensor, distance, row)
		}

	}

	// grid.PrintYFlipped(".")

	notBeacon := 0
	bounds := grid.Bounds()
	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
		if grid.Get(image.Pt(x, row), " ") == "#" {
			notBeacon++
		}
	}

	fmt.Println(notBeacon)
}

func drawSensorAll(grid aoc.Grid2[string], origin image.Point, distance int) {
	for x := -distance; x <= distance; x++ {
		for y := -distance; y <= distance; y++ {
			dest := origin.Add(image.Pt(x, y))

			if aoc.ManhattanDistancePt(origin, dest) <= distance {
				if _, ok := grid[dest]; !ok {
					grid[dest] = "#"
				}
			}
		}
	}
}

func part1draw(name string) {
	grid := aoc.Grid2[string]{}

	for _, s := range aoc.Strings(name) {
		var sensor, beacon image.Point
		fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)

		grid[sensor], grid[beacon] = "S", "B"
		distance := aoc.ManhattanDistancePt(sensor, beacon)
		drawSensorAll(grid, sensor, distance)
	}

	grid.PrintYFlipped(".")

	// notBeacon := 0
	// bounds := grid.Bounds()
	// for x := bounds.Min.X; x <= bounds.Max.X; x++ {
	// 	if grid.Get(image.Pt(x, row), ' ') == '#' {
	// 		notBeacon++
	// 	}
	// }

	// fmt.Println(notBeacon)
}

func main() {
	// part1("test.txt", 10)
	// part1("input.txt", 2000000)

	part1draw("test.txt")
	// part1("input.txt", 2000000)

	// println("------")

	// part2("test.txt")
	// part2("input.txt")

}
