package aoc

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strings"
)

// Row is a series of strings along the x-axis
type Row []string

// Grid is a series of rows along the y-axis
type Grid []Row

// NewGrid returns a Grid from []strings, breaking each character into a single rune as string
func NewGrid(in []string) (grid Grid) {
	grid = Grid{}
	for _, i := range in {
		var row Row
		for _, c := range i {
			row = append(row, string(c))
		}
		grid = append(grid, row)
	}
	return
}

// NewBlankGrid returns a Grid
func NewBlankGrid(width, height int, init string) (grid Grid) {
	grid = Grid{}
	for y := 0; y < height; y++ {
		var row Row
		for x := 0; x < width; x++ {
			row = append(row, init)
		}
		grid = append(grid, row)
	}
	return
}

// Print the grid
func (grid Grid) Print() {
	for _, row := range grid {
		for _, r := range row {
			fmt.Print(r)
		}
		fmt.Println()
	}
}

// PrintSep the grid with separator
func (grid Grid) PrintSep(sep string) {
	for _, row := range grid {
		// for _, r := range row {
		fmt.Print(strings.Join(row, sep))
		// }
		fmt.Println()
	}
}

// Iterate the grid, return false if f returns false
func (grid Grid) Iterate(f func(x, y int, s string) bool) bool {
	for y, row := range grid {
		for x, s := range row {
			if !f(x, y, s) {
				return false
			}
		}
	}
	return true
}

// SlopeIterate from origin
func (grid Grid) SlopeIterate(x, y int, dx, dy int, f func(x, y int, s string) bool) {
	bounds := grid.Bounds()

	for {
		x += dx
		y += dy

		// out of bounds?
		if !image.Pt(x, y).In(bounds) {
			return
		}

		// call f
		if !f(x, y, grid.At(x, y)) {
			return
		}
	}
}

// Copy and return a new grid
func (grid Grid) Copy() (c Grid) {
	c = NewBlankGrid(grid.Width(), grid.Height(), "")

	grid.Iterate(func(x, y int, s string) bool {
		c.Set(x, y, s)
		return true
	})

	return
}

// Equal returns if the grids are the same
func (grid Grid) Equal(in Grid) (equal bool) {
	if (grid.Width() != in.Width()) || (grid.Height() != in.Height()) {
		return
	}

	return grid.Iterate(func(x, y int, s string) bool {
		return s == in.At(x, y)
	})
}

// At returns the string at (x,y). Conveinence wrapper since grid is y-coordinate first.
func (grid Grid) At(x, y int) string { return grid[y][x] }

// Set the string at (x,y). Conveinence wrapper since grid is y-coordinate first.
func (grid Grid) Set(x, y int, s string) { grid[y][x] = s }

// Get the string at (x,y). Returns default if out of bounds
func (grid Grid) Get(x, y int, s string) string {
	if x < 0 || x >= grid.Width() || y < 0 || y >= grid.Height() {
		return s
	}

	return grid.At(x, y)
}

// FindFirst string in grid, return (x,y)
func (grid Grid) FindFirst(find string) (x, y int) {
	x = -1
	y = -1

	grid.Iterate(func(gx, gy int, s string) bool {
		if s == find {
			x = gx
			y = gy
			return false
		}
		return true
	})

	return
}

// Count occurences of s in the grid
func (grid Grid) Count(t string) (total int) {
	grid.Iterate(func(x, y int, s string) bool {
		if s == t {
			total++
		}
		return true
	})
	return
}

// Height of the grid
func (grid Grid) Height() int { return len(grid) }

// Width of the grid
func (grid Grid) Width() int { return len(grid[0]) }

// Bounds of the grid (0, 0, width, height)
func (grid Grid) Bounds() image.Rectangle { return image.Rect(0, 0, grid.Width(), grid.Height()) }

// NewRGBAImage returns an RGBAImage sized to the grid * scale
func (grid Grid) NewRGBAImage(scale int) draw.Image {
	return image.NewRGBA(image.Rect(0, 0, grid.Width()*scale, grid.Height()*scale))
}

// DrawImage convert strings to color
func (grid Grid) DrawImage(img draw.Image, scale int, mapping map[string]color.Color) {
	grid.Iterate(func(x, y int, s string) bool {
		c, ok := mapping[s]
		if !ok {
			return true
		}

		xo := x * scale
		yo := y * scale

		r := image.Rect(xo, yo, xo+scale, yo+scale)
		draw.Draw(img, r, &image.Uniform{c}, image.Point{}, draw.Src)

		return true
	})
}

// Column of values
func (grid Grid) Column(col int) (column Row) {
	grid.Iterate(func(gx, gy int, s string) bool {
		if col == gx {
			column = append(column, s)
		}
		return true
	})
	return
}

// Columns of values
func (grid Grid) Columns() (columns []Row) {
	for x := 0; x < grid.Width(); x++ {
		var column Row
		grid.Iterate(func(gx, gy int, s string) bool {
			if x == gx {
				column = append(column, s)
			}
			return true
		})

		columns = append(columns, column)
	}
	return
}

// Row of values
func (grid Grid) Row(r int) (row Row) {
	grid.Iterate(func(gx, gy int, s string) bool {
		if r == gy {
			row = append(row, s)
		}
		return true
	})
	return
}

// Rows of values
func (grid Grid) Rows() (rows []Row) {
	return grid
}
