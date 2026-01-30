package aoc

import (
	"image"
	"testing"
)

// =============================================================================
// Grid Creation Tests
// =============================================================================

func TestNewGrid(t *testing.T) {
	input := []string{"abc", "def", "ghi"}
	grid := NewGrid(input)

	if grid.Height() != 3 {
		t.Errorf("Height: expected 3, got %d", grid.Height())
	}
	if grid.Width() != 3 {
		t.Errorf("Width: expected 3, got %d", grid.Width())
	}
	if grid.At(0, 0) != "a" {
		t.Errorf("At(0,0): expected 'a', got '%s'", grid.At(0, 0))
	}
	if grid.At(2, 2) != "i" {
		t.Errorf("At(2,2): expected 'i', got '%s'", grid.At(2, 2))
	}
}

func TestNewGridUnicode(t *testing.T) {
	input := []string{"日本語"}
	grid := NewGrid(input)

	if grid.Width() != 3 {
		t.Errorf("Unicode Width: expected 3, got %d", grid.Width())
	}
	if grid.At(0, 0) != "日" {
		t.Errorf("Unicode At(0,0): expected '日', got '%s'", grid.At(0, 0))
	}
}

func TestNewBlankGrid(t *testing.T) {
	grid := NewBlankGrid(5, 3, ".")

	if grid.Width() != 5 {
		t.Errorf("Width: expected 5, got %d", grid.Width())
	}
	if grid.Height() != 3 {
		t.Errorf("Height: expected 3, got %d", grid.Height())
	}
	if grid.At(2, 1) != "." {
		t.Errorf("At(2,1): expected '.', got '%s'", grid.At(2, 1))
	}
}

// =============================================================================
// Grid Access Tests
// =============================================================================

func TestGridAtSet(t *testing.T) {
	grid := NewBlankGrid(3, 3, ".")
	grid.Set(1, 1, "X")

	if grid.At(1, 1) != "X" {
		t.Errorf("Set/At: expected 'X', got '%s'", grid.At(1, 1))
	}
}

func TestGridGet(t *testing.T) {
	grid := NewGrid([]string{"abc", "def"})

	// In bounds
	if grid.Get(1, 0, "?") != "b" {
		t.Errorf("Get in bounds: expected 'b', got '%s'", grid.Get(1, 0, "?"))
	}

	// Out of bounds
	if grid.Get(-1, 0, "?") != "?" {
		t.Errorf("Get negative x: expected '?', got '%s'", grid.Get(-1, 0, "?"))
	}
	if grid.Get(0, -1, "?") != "?" {
		t.Errorf("Get negative y: expected '?', got '%s'", grid.Get(0, -1, "?"))
	}
	if grid.Get(10, 0, "?") != "?" {
		t.Errorf("Get x too large: expected '?', got '%s'", grid.Get(10, 0, "?"))
	}
	if grid.Get(0, 10, "?") != "?" {
		t.Errorf("Get y too large: expected '?', got '%s'", grid.Get(0, 10, "?"))
	}
}

func TestGridBounds(t *testing.T) {
	grid := NewBlankGrid(5, 3, ".")
	bounds := grid.Bounds()

	expected := image.Rect(0, 0, 5, 3)
	if bounds != expected {
		t.Errorf("Bounds: expected %v, got %v", expected, bounds)
	}
}

// =============================================================================
// Grid Search Tests
// =============================================================================

func TestGridFindFirst(t *testing.T) {
	grid := NewGrid([]string{
		"..X..",
		".....",
		"..O..",
	})

	x, y := grid.FindFirst("X")
	if x != 2 || y != 0 {
		t.Errorf("FindFirst X: expected (2,0), got (%d,%d)", x, y)
	}

	x, y = grid.FindFirst("O")
	if x != 2 || y != 2 {
		t.Errorf("FindFirst O: expected (2,2), got (%d,%d)", x, y)
	}

	x, y = grid.FindFirst("Z")
	if x != -1 || y != -1 {
		t.Errorf("FindFirst not found: expected (-1,-1), got (%d,%d)", x, y)
	}
}

func TestGridCount(t *testing.T) {
	grid := NewGrid([]string{
		"##..#",
		".#.#.",
		"..#..",
	})

	if grid.Count("#") != 6 {
		t.Errorf("Count #: expected 6, got %d", grid.Count("#"))
	}
	if grid.Count(".") != 9 {
		t.Errorf("Count .: expected 9, got %d", grid.Count("."))
	}
	if grid.Count("X") != 0 {
		t.Errorf("Count X: expected 0, got %d", grid.Count("X"))
	}
}

// =============================================================================
// Grid Iteration Tests
// =============================================================================

func TestGridIterate(t *testing.T) {
	grid := NewGrid([]string{"ab", "cd"})

	var visited []string
	grid.Iterate(func(x, y int, s string) bool {
		visited = append(visited, s)
		return true
	})

	if len(visited) != 4 {
		t.Errorf("Iterate: expected 4 visits, got %d", len(visited))
	}

	// Should visit in row-major order
	expected := []string{"a", "b", "c", "d"}
	for i, v := range visited {
		if v != expected[i] {
			t.Errorf("Iterate order: expected %v, got %v", expected, visited)
			break
		}
	}
}

func TestGridIterateEarlyExit(t *testing.T) {
	grid := NewGrid([]string{"abc", "def", "ghi"})

	count := 0
	result := grid.Iterate(func(x, y int, s string) bool {
		count++
		return s != "e" // Stop when we hit 'e'
	})

	if result {
		t.Error("Iterate should return false on early exit")
	}
	if count != 5 { // a, b, c, d, e
		t.Errorf("Iterate early exit: expected 5 visits, got %d", count)
	}
}

func TestGridSlopeIterate(t *testing.T) {
	grid := NewGrid([]string{
		"12345",
		"67890",
		"abcde",
	})

	var visited []string
	grid.SlopeIterate(0, 0, 1, 1, func(x, y int, s string) bool {
		visited = append(visited, s)
		return true
	})

	// Starting from (0,0), moving (+1,+1): (1,1)='7', (2,2)='c'
	expected := []string{"7", "c"}
	if len(visited) != len(expected) {
		t.Errorf("SlopeIterate: expected %d visits, got %d", len(expected), len(visited))
	}
	for i, v := range visited {
		if v != expected[i] {
			t.Errorf("SlopeIterate: expected %v, got %v", expected, visited)
			break
		}
	}
}

func TestGridSlopeIterateOutOfBounds(t *testing.T) {
	grid := NewGrid([]string{"abc", "def"})

	var count int
	grid.SlopeIterate(2, 0, 1, 0, func(x, y int, s string) bool {
		count++
		return true
	})

	// Starting at (2,0), moving (+1,0) goes out of bounds immediately
	if count != 0 {
		t.Errorf("SlopeIterate out of bounds: expected 0 visits, got %d", count)
	}
}

// =============================================================================
// Grid Copy/Compare Tests
// =============================================================================

func TestGridCopy(t *testing.T) {
	grid := NewGrid([]string{"abc", "def"})
	copy := grid.Copy()

	// Modify original
	grid.Set(0, 0, "X")

	// Copy should be unchanged
	if copy.At(0, 0) != "a" {
		t.Error("Copy should be independent of original")
	}
}

func TestGridEqual(t *testing.T) {
	grid1 := NewGrid([]string{"abc", "def"})
	grid2 := NewGrid([]string{"abc", "def"})
	grid3 := NewGrid([]string{"abc", "xyz"})
	grid4 := NewGrid([]string{"ab", "cd"})

	if !grid1.Equal(grid2) {
		t.Error("Identical grids should be equal")
	}
	if grid1.Equal(grid3) {
		t.Error("Different content should not be equal")
	}
	if grid1.Equal(grid4) {
		t.Error("Different size should not be equal")
	}
}

// =============================================================================
// Grid Row/Column Tests
// =============================================================================

func TestGridRow(t *testing.T) {
	grid := NewGrid([]string{"abc", "def", "ghi"})

	row := grid.Row(1)
	expected := Row{"d", "e", "f"}

	if len(row) != len(expected) {
		t.Errorf("Row: expected length %d, got %d", len(expected), len(row))
	}
	for i, v := range row {
		if v != expected[i] {
			t.Errorf("Row: expected %v, got %v", expected, row)
			break
		}
	}
}

func TestGridColumn(t *testing.T) {
	grid := NewGrid([]string{"abc", "def", "ghi"})

	col := grid.Column(1)
	expected := Row{"b", "e", "h"}

	if len(col) != len(expected) {
		t.Errorf("Column: expected length %d, got %d", len(expected), len(col))
	}
	for i, v := range col {
		if v != expected[i] {
			t.Errorf("Column: expected %v, got %v", expected, col)
			break
		}
	}
}

func TestGridRows(t *testing.T) {
	grid := NewGrid([]string{"ab", "cd"})
	rows := grid.Rows()

	if len(rows) != 2 {
		t.Errorf("Rows: expected 2 rows, got %d", len(rows))
	}
}

func TestGridColumns(t *testing.T) {
	grid := NewGrid([]string{"ab", "cd"})
	cols := grid.Columns()

	if len(cols) != 2 {
		t.Errorf("Columns: expected 2 columns, got %d", len(cols))
	}

	// First column should be a, c
	if cols[0][0] != "a" || cols[0][1] != "c" {
		t.Errorf("Columns[0]: expected [a,c], got %v", cols[0])
	}
}

// =============================================================================
// Grid Line Iteration Tests
// =============================================================================

func TestGridIterateLine(t *testing.T) {
	grid := NewGrid([]string{
		"12345",
		"67890",
		"abcde",
	})

	var visited []string
	grid.IterateLine(0, 0, 4, 2, func(x, y int, s string) bool {
		visited = append(visited, s)
		return true
	})

	// Diagonal from (0,0) to (4,2)
	// Steps: max(4,2) = 4, so increments are (1, 0.5) -> effectively (1,0), (1,1), (1,0), (1,1)
	// This is integer division based, so: (0,0), (1,0), (2,1), (3,1), (4,2)
	if len(visited) != 5 {
		t.Errorf("IterateLine: expected 5 points, got %d: %v", len(visited), visited)
	}
}

func TestGridIterateLineHorizontal(t *testing.T) {
	grid := NewGrid([]string{"abcde"})

	var visited []string
	grid.IterateLine(0, 0, 4, 0, func(x, y int, s string) bool {
		visited = append(visited, s)
		return true
	})

	expected := []string{"a", "b", "c", "d", "e"}
	if len(visited) != len(expected) {
		t.Errorf("IterateLine horizontal: expected %d, got %d", len(expected), len(visited))
	}
}

func TestGridIterateLineVertical(t *testing.T) {
	grid := NewGrid([]string{"a", "b", "c", "d", "e"})

	var visited []string
	grid.IterateLine(0, 0, 0, 4, func(x, y int, s string) bool {
		visited = append(visited, s)
		return true
	})

	expected := []string{"a", "b", "c", "d", "e"}
	if len(visited) != len(expected) {
		t.Errorf("IterateLine vertical: expected %d, got %d", len(expected), len(visited))
	}
}

// =============================================================================
// MaxXY Tests
// =============================================================================

func TestMaxXY(t *testing.T) {
	points := []image.Point{
		{1, 2},
		{5, 3},
		{2, 8},
		{0, 0},
	}

	x, y := MaxXY(points)
	if x != 5 {
		t.Errorf("MaxXY x: expected 5, got %d", x)
	}
	if y != 8 {
		t.Errorf("MaxXY y: expected 8, got %d", y)
	}
}

func TestMaxXYEmpty(t *testing.T) {
	x, y := MaxXY([]image.Point{})
	if x != 0 || y != 0 {
		t.Errorf("MaxXY empty: expected (0,0), got (%d,%d)", x, y)
	}
}
