package main

import (
	"fmt"
	"image"
	"math"

	"github.com/tidwall/geometry"
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

			if aoc.ManhattanDistancePt(origin, dest) == distance {
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

	// grid.PrintYFlipped(".")

	// notBeacon := 0
	// bounds := grid.Bounds()
	// for x := bounds.Min.X; x <= bounds.Max.X; x++ {
	// 	if grid.Get(image.Pt(x, row), ' ') == '#' {
	// 		notBeacon++
	// 	}
	// }

	// fmt.Println(notBeacon)
}

func geometryPoint(p image.Point) geometry.Point {
	return geometry.Point{X: float64(p.X), Y: float64(p.Y)}
}

func PointsOnLine(start, end image.Point) (line []image.Point) {
	d := end.Sub(start)

	steps := aoc.Abs(d.Y)
	if aoc.Abs(d.X) > aoc.Abs(d.Y) {
		steps = aoc.Abs(d.X)
	}

	delta := d.Div(steps)
	cursor := start

	for v := 0; v <= steps; v++ {
		line = append(line, cursor)
		cursor = cursor.Add(delta)
	}

	// fmt.Println(start, end, line)

	return
}

func part1geometry(name string, row float64) {
	features := map[geometry.Point]string{}
	area := []*geometry.Poly{}

	testPts := aoc.Set[geometry.Point]{}

	for _, s := range aoc.Strings(name) {
		var sensor, beacon image.Point
		fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)

		features[geometryPoint(sensor)], features[geometryPoint(beacon)] = "S", "B"
		distance := aoc.ManhattanDistancePt(sensor, beacon)

		// left, up, right, down
		deltas := []image.Point{{-distance, 0}, {0, distance}, {distance, 0}, {0, -distance}}
		points := aoc.Map(func(d image.Point) geometry.Point {
			p := sensor.Add(d)
			return geometryPoint(p)
		}, deltas)

		coverage := geometry.NewPoly([]geometry.Point{points[0], points[1], points[2], points[3], points[0]},
			[][]geometry.Point{}, nil)

		area = append(area, coverage)

		// part 2
		distance++
		deltas = []image.Point{{-distance, 0}, {0, distance}, {distance, 0}, {0, -distance}, {-distance, 0}}
		for d := 0; d < len(deltas)-1; d++ {
			start, end := sensor.Add(deltas[d]), sensor.Add(deltas[d+1])
			testPts.AddMany(aoc.Map(geometryPoint, PointsOnLine(start, end)))
		}
	}

	minX, maxX, minY, maxY := math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64
	for _, p := range area {
		bounds := p.Rect()
		if bounds.Min.X < minX {
			minX = bounds.Min.X
		} else if bounds.Max.X > maxX {
			maxX = bounds.Max.X
		}

		if bounds.Min.Y < minY {
			minY = bounds.Min.Y
		} else if bounds.Max.Y > maxY {
			maxY = bounds.Max.Y
		}
	}

	notBeacon := 0
	for x := minX; x <= maxX; x++ {
		pt := geometry.Point{X: x, Y: row}

		if _, found := features[pt]; found {
			continue
		}

		for _, poly := range area {
			if poly.ContainsPoint(pt) {
				notBeacon++
				break
			}
		}
	}

	fmt.Println(notBeacon)

	{
		minX, maxX, minY, maxY := math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64
		for pt, s := range features {
			if s != "S" {
				continue
			}

			if pt.X < minX {
				minX = pt.X
			}

			if pt.X > maxX {
				maxX = pt.X
			}

			if pt.Y < minY {
				minY = pt.Y
			}
			if pt.Y > maxY {
				maxY = pt.Y
			}
		}

		searchArea := geometry.Rect{Min: geometry.Point{minX, minY}, Max: geometry.Point{maxX, maxY}}
		fmt.Println(searchArea)

		// testPts.Add(geometry.Point{14, 11})
		for _, pt := range testPts.Values() {
			if !searchArea.ContainsPoint(pt) {
				continue
			}

			found := false
			for _, poly := range area {
				if poly.ContainsPoint(pt) {
					found = true
					break
				}
			}

			if !found {
				fmt.Println("found", int(pt.X), int(pt.Y), (4000000*int(pt.X))+int(pt.Y))
				break
			}
		}
	}
}

func main() {
	// part1("test.txt", 10)
	// part1("input.txt", 2000000)

	part1draw("test.txt")
	part1geometry("test.txt", 10)
	part1geometry("input.txt", 2000000)

	// println("------")

	// part2("test.txt")
	// part2("input.txt")

}
