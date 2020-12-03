package aoc

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
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

// Print the grid
func (grid Grid) Print() {
	for _, row := range grid {
		for _, r := range row {
			fmt.Print(r)
		}
		fmt.Println()
	}
}

// At returns the string at (x,y). Conveinence wrapper since grid is y-coordinate first.
func (grid Grid) At(x, y int) string { return grid[y][x] }

// Set the string at (x,y). Conveinence wrapper since grid is y-coordinate first.
func (grid Grid) Set(x, y int, s string) { grid[y][x] = s }

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
	for x := 0; x < grid.Width(); x++ {
		for y := 0; y < grid.Height(); y++ {
			c, ok := mapping[grid.At(x, y)]
			if !ok {
				continue
			}

			xo := x * scale
			yo := y * scale

			r := image.Rect(xo, yo, xo+scale, yo+scale)
			draw.Draw(img, r, &image.Uniform{c}, image.ZP, draw.Src)
		}
	}
}
