package aoc

import (
	"fmt"
	"image"
	"maps"
	"math"
	"slices"
)

// SparseGrid is a sparse 2D grid using a map for storage.
// Good for large/infinite grids where not all cells are filled.
// Supports negative coordinates and generic cell types.
type SparseGrid[T comparable] map[image.Point]T

func LoadIntGrid(in []string) (g SparseGrid[int]) {
	return LoadSparseGrid(func(r rune) int { return int(r - '0') }, in)
}

func LoadStringGrid(in []string) (g SparseGrid[string]) {
	return LoadSparseGrid(func(r rune) string { return string(r) }, in)
}

func LoadRuneGrid(in []string) (g SparseGrid[rune]) {
	return LoadSparseGrid(func(r rune) rune { return r }, in)
}

func LoadSparseGrid[T comparable](f func(rune) T, in []string) (g SparseGrid[T]) {
	g = make(SparseGrid[T])

	for y, row := range in {
		for x, value := range row {
			g[image.Pt(x, y)] = f(value)
		}
	}

	return
}

// SlopeIterate from origin, stepping by delta each iteration
func (grid SparseGrid[T]) SlopeIterate(origin image.Point, delta image.Point, f func(pt image.Point, v T) bool) {
	bounds := grid.Bounds()

	for {
		origin = origin.Add(delta)

		if !origin.In(bounds) {
			return
		}

		// since SparseGrid is sparse, only callback if it exists
		if value, ok := grid[origin]; !ok {
			continue
		} else if !f(origin, value) {
			return
		}
	}
}

func (grid SparseGrid[T]) Bounds() (bounds image.Rectangle) {
	return Bounds(slices.Collect(maps.Keys(grid)))
}

// Print the grid
func (grid SparseGrid[T]) Print(empty T) {
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
func (grid SparseGrid[T]) PrintYFlipped(empty T) {
	bounds := grid.Bounds()
	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		fmt.Printf("%3d:  ", y)

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

func (grid SparseGrid[T]) Get(in image.Point, empty T) (value T) {
	value = empty
	if v, ok := grid[in]; ok {
		value = v
	}
	return
}

func (grid SparseGrid[T]) Exists(in []image.Point) (pts []image.Point) {
	for _, i := range in {
		if _, ok := grid[i]; ok {
			pts = append(pts, i)
		}
	}
	return
}

func (grid SparseGrid[T]) FourWayAdjacent(in image.Point) (pts []image.Point) {
	return grid.Exists(Map(in.Add, []image.Point{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}))
}

func (grid SparseGrid[T]) EightWayAdjacent(in image.Point) (pts []image.Point) {
	return grid.Exists(Map(in.Add, []image.Point{
		{-1, 1}, {0, 1}, {1, 1},
		{-1, 0} /*{0,0}*/, {1, 0},
		{-1, -1}, {0, -1}, {1, -1}}))
}

func Contains[T comparable](grid SparseGrid[T], value T) (pts []image.Point) {
	for pt, v := range grid {
		if v == value {
			pts = append(pts, pt)
		}
	}
	return
}

func (grid SparseGrid[T]) Contains(value T) (pts []image.Point) {
	return Contains(grid, value)
}

func Bounds(points []image.Point) (bounds image.Rectangle) {
	bounds.Max.X, bounds.Max.Y = math.MinInt, math.MinInt
	bounds.Min.X, bounds.Min.Y = math.MaxInt, math.MaxInt

	for _, point := range points {
		if point.X < bounds.Min.X {
			bounds.Min.X = point.X
		}

		if point.X > bounds.Max.X {
			bounds.Max.X = point.X
		}

		if point.Y < bounds.Min.Y {
			bounds.Min.Y = point.Y
		}

		if point.Y > bounds.Max.Y {
			bounds.Max.Y = point.Y
		}
	}
	return
}

func (grid SparseGrid[T]) IterateLine(start, end image.Point, f func(pt image.Point, v T) bool) bool {
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

// func rotate45ccw(i image.Point) image.Point { return image.Point{i.X - i.Y, i.X + i.Y} }
// func rotate45cw(i image.Point) image.Point  { return image.Point{i.X - i.Y, i.X + i.Y} }

func Rotate90cw(i image.Point) image.Point   { return image.Point{i.Y, -i.X} }
func Rotate90ccw(i image.Point) image.Point  { return image.Point{-i.Y, i.X} }
func Rotate180(i image.Point) image.Point    { return image.Point{-i.X, -i.Y} }
func Rotate270ccw(i image.Point) image.Point { return Rotate90cw(i) }
func Rotate270cw(i image.Point) image.Point  { return Rotate90ccw(i) }
