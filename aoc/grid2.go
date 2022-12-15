package aoc

import (
	"fmt"
	"image"
	"math"

	"golang.org/x/exp/maps"
)

type Grid2[T comparable] map[image.Point]T

func LoadIntGrid(in []string) (g Grid2[int]) {
	return LoadGrid2(func(r rune) int { return int(r - '0') }, in)
}

func LoadStringGrid(in []string) (g Grid2[string]) {
	return LoadGrid2(func(r rune) string { return string(r) }, in)
}

func LoadRuneGrid(in []string) (g Grid2[rune]) {
	return LoadGrid2(func(r rune) rune { return r }, in)
}

func LoadGrid2[T comparable](f func(rune) T, in []string) (g Grid2[T]) {
	g = make(Grid2[T])

	for y, row := range in {
		for x, value := range row {
			g[image.Pt(x, y)] = f(value)
		}
	}

	return
}

// SlopeIterate from origin
func (grid Grid2[T]) SlopeIterate(origin image.Point, delta image.Point, f func(pt image.Point, v T) bool) {
	bounds := grid.Bounds()

	for {
		current := origin.Add(delta)

		if !current.In(bounds) {
			return
		}

		// since Grid2 is sparse, only callback if it exists
		if value, ok := grid[current]; !ok {
			continue
		} else if !f(current, value) {
			return
		}
	}
}

func (grid Grid2[T]) Bounds() (bounds image.Rectangle) {
	return Bounds(maps.Keys(grid))
}

// Print the grid
func (grid Grid2[T]) Print(empty T) {
	bounds := grid.Bounds()
	for y := bounds.Max.Y; y >= bounds.Min.Y; y-- {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			if value, ok := grid[image.Pt(x, y)]; ok {
				fmt.Print(value)
			} else {
				fmt.Print(empty)
			}
		}
		println()
	}
}

// Print the grid
func (grid Grid2[T]) PrintYFlipped(empty T) {
	bounds := grid.Bounds()
	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			if value, ok := grid[image.Pt(x, y)]; ok {
				fmt.Print(value)
			} else {
				fmt.Print(empty)
			}
		}
		println()
	}
}

func (grid Grid2[T]) Get(in image.Point, empty T) (value T) {
	value = empty
	if v, ok := grid[in]; ok {
		value = v
	}
	return
}

func (grid Grid2[T]) Exists(in []image.Point) (pts []image.Point) {
	for _, i := range in {
		if _, ok := grid[i]; ok {
			pts = append(pts, i)
		}
	}
	return
}

func (grid Grid2[T]) FourWayAdjacent(in image.Point) (pts []image.Point) {
	return grid.Exists(Map(in.Add, []image.Point{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}))
}

func (grid Grid2[T]) EightWayAdjacent(in image.Point) (pts []image.Point) {
	return grid.Exists(Map(in.Add, []image.Point{
		{-1, 1}, {0, 1}, {1, 1},
		{-1, 0} /*{0,0}*/, {1, 0},
		{-1, -1}, {0, -1}, {1, -1}}))
}

func Contains[T comparable](grid Grid2[T], value T) (pts []image.Point) {
	for pt, v := range grid {
		if v == value {
			pts = append(pts, pt)
		}
	}
	return
}

func (grid Grid2[T]) Contains(value T) (pts []image.Point) {
	return Contains(grid, value)
}

func Bounds(points []image.Point) (bounds image.Rectangle) {
	bounds.Max.X, bounds.Max.Y = math.MinInt, math.MinInt
	bounds.Min.X, bounds.Min.Y = math.MaxInt, math.MaxInt

	for _, point := range points {
		if point.X < bounds.Min.X {
			bounds.Min.X = point.X
		} else if point.X > bounds.Max.X {
			bounds.Max.X = point.X
		}

		if point.Y < bounds.Min.Y {
			bounds.Min.Y = point.Y
		} else if point.Y > bounds.Max.Y {
			bounds.Max.Y = point.Y
		}
	}
	return
}

func (grid Grid2[T]) IterateLine(start, end image.Point, f func(pt image.Point, v T) bool) bool {
	d := end.Sub(start)

	steps := Abs(d.Y)
	if Abs(d.X) > Abs(d.Y) {
		steps = Abs(d.X)
	}

	delta := d.Div(steps)
	cursor := start

	for v := 0; v < steps+1; v++ {
		if !f(cursor, grid[cursor]) {
			return false
		}

		cursor = cursor.Add(delta)
	}

	return true
}

func ManhattanDistancePt(p, p1 image.Point) (distance int) {
	return Abs(p.X-p1.X) + Abs(p.Y-p1.Y)
}
