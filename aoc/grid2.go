package aoc

import (
	"image"

	"golang.org/x/exp/maps"
)

type Grid2[T any] map[image.Point]T

func LoadIntGrid(in []string) (g Grid2[int]) {
	return LoadGrid2(func(r rune) int { return int(r - '0') }, in)
}

func LoadStringGrid(in []string) (g Grid2[string]) {
	return LoadGrid2(func(r rune) string { return string(r) }, in)
}

func LoadGrid2[T any](f func(rune) T, in []string) (g Grid2[T]) {
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
func (grid Grid2[T]) Print(empty string) {
	bounds := grid.Bounds()
	for y := bounds.Max.Y; y >= bounds.Min.Y; y-- {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			if value, ok := grid[image.Pt(x, y)]; ok {
				print(value)
			} else {
				print(empty)
			}
		}
		println()
	}
}

func Bounds(points []image.Point) (bounds image.Rectangle) {
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
