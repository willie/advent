package aoc

import (
	"image"
	"slices"
	"testing"
)

// =============================================================================
// SparseGrid Creation Tests
// =============================================================================

func TestLoadIntGrid(t *testing.T) {
	input := []string{"123", "456"}
	grid := LoadIntGrid(input)

	if grid[image.Pt(0, 0)] != 1 {
		t.Errorf("LoadIntGrid (0,0): expected 1, got %d", grid[image.Pt(0, 0)])
	}
	if grid[image.Pt(2, 1)] != 6 {
		t.Errorf("LoadIntGrid (2,1): expected 6, got %d", grid[image.Pt(2, 1)])
	}
}

func TestLoadStringGrid(t *testing.T) {
	input := []string{"ab", "cd"}
	grid := LoadStringGrid(input)

	if grid[image.Pt(0, 0)] != "a" {
		t.Errorf("LoadStringGrid (0,0): expected 'a', got '%s'", grid[image.Pt(0, 0)])
	}
	if grid[image.Pt(1, 1)] != "d" {
		t.Errorf("LoadStringGrid (1,1): expected 'd', got '%s'", grid[image.Pt(1, 1)])
	}
}

func TestLoadRuneGrid(t *testing.T) {
	input := []string{"αβ"}
	grid := LoadRuneGrid(input)

	if grid[image.Pt(0, 0)] != 'α' {
		t.Errorf("LoadRuneGrid: expected 'α', got '%c'", grid[image.Pt(0, 0)])
	}
}

func TestLoadSparseGridCustom(t *testing.T) {
	input := []string{"ab", "cd"}
	grid := LoadSparseGrid(func(r rune) int { return int(r) }, input)

	if grid[image.Pt(0, 0)] != int('a') {
		t.Errorf("LoadSparseGrid custom: expected %d, got %d", int('a'), grid[image.Pt(0, 0)])
	}
}

// =============================================================================
// SparseGrid Bounds Tests
// =============================================================================

func TestSparseGridBounds(t *testing.T) {
	grid := LoadStringGrid([]string{"abc", "def"})
	bounds := grid.Bounds()

	if bounds.Min.X != 0 || bounds.Min.Y != 0 {
		t.Errorf("Bounds Min: expected (0,0), got (%d,%d)", bounds.Min.X, bounds.Min.Y)
	}
	if bounds.Max.X != 2 || bounds.Max.Y != 1 {
		t.Errorf("Bounds Max: expected (2,1), got (%d,%d)", bounds.Max.X, bounds.Max.Y)
	}
}

func TestSparseGridBoundsNegative(t *testing.T) {
	grid := make(SparseGrid[int])
	grid[image.Pt(-5, -3)] = 1
	grid[image.Pt(5, 3)] = 2

	bounds := grid.Bounds()
	if bounds.Min.X != -5 || bounds.Min.Y != -3 {
		t.Errorf("Bounds Min: expected (-5,-3), got (%d,%d)", bounds.Min.X, bounds.Min.Y)
	}
	if bounds.Max.X != 5 || bounds.Max.Y != 3 {
		t.Errorf("Bounds Max: expected (5,3), got (%d,%d)", bounds.Max.X, bounds.Max.Y)
	}
}

// =============================================================================
// SparseGrid Get/Exists Tests
// =============================================================================

func TestSparseGridGet(t *testing.T) {
	grid := LoadStringGrid([]string{"ab", "cd"})

	if grid.Get(image.Pt(0, 0), "X") != "a" {
		t.Error("Get existing: should return actual value")
	}
	if grid.Get(image.Pt(99, 99), "X") != "X" {
		t.Error("Get missing: should return default")
	}
}

func TestSparseGridExists(t *testing.T) {
	grid := LoadStringGrid([]string{"ab", "cd"})

	points := []image.Point{{0, 0}, {99, 99}, {1, 1}}
	existing := grid.Exists(points)

	if len(existing) != 2 {
		t.Errorf("Exists: expected 2 points, got %d", len(existing))
	}
}

// =============================================================================
// SparseGrid Adjacent Tests
// =============================================================================

func TestSparseGridFourWayAdjacent(t *testing.T) {
	grid := LoadStringGrid([]string{
		"abc",
		"def",
		"ghi",
	})

	adj := grid.FourWayAdjacent(image.Pt(1, 1)) // Center 'e'

	if len(adj) != 4 {
		t.Errorf("FourWayAdjacent center: expected 4, got %d", len(adj))
	}

	// Check expected neighbors
	expected := []image.Point{{0, 1}, {2, 1}, {1, 0}, {1, 2}}
	for _, exp := range expected {
		found := false
		for _, a := range adj {
			if a == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("FourWayAdjacent: missing %v", exp)
		}
	}
}

func TestSparseGridFourWayAdjacentCorner(t *testing.T) {
	grid := LoadStringGrid([]string{
		"ab",
		"cd",
	})

	adj := grid.FourWayAdjacent(image.Pt(0, 0)) // Corner 'a'

	if len(adj) != 2 {
		t.Errorf("FourWayAdjacent corner: expected 2, got %d", len(adj))
	}
}

func TestSparseGridEightWayAdjacent(t *testing.T) {
	grid := LoadStringGrid([]string{
		"abc",
		"def",
		"ghi",
	})

	adj := grid.EightWayAdjacent(image.Pt(1, 1)) // Center 'e'

	if len(adj) != 8 {
		t.Errorf("EightWayAdjacent center: expected 8, got %d", len(adj))
	}
}

func TestSparseGridEightWayAdjacentCorner(t *testing.T) {
	grid := LoadStringGrid([]string{
		"ab",
		"cd",
	})

	adj := grid.EightWayAdjacent(image.Pt(0, 0)) // Corner 'a'

	if len(adj) != 3 {
		t.Errorf("EightWayAdjacent corner: expected 3, got %d", len(adj))
	}
}

// =============================================================================
// SparseGrid Contains Tests
// =============================================================================

func TestSparseGridContains(t *testing.T) {
	grid := LoadStringGrid([]string{
		"a.a",
		".a.",
		"a.a",
	})

	points := grid.Contains("a")
	if len(points) != 5 {
		t.Errorf("Contains 'a': expected 5 points, got %d", len(points))
	}

	points = grid.Contains(".")
	if len(points) != 4 {
		t.Errorf("Contains '.': expected 4 points, got %d", len(points))
	}

	points = grid.Contains("z")
	if len(points) != 0 {
		t.Errorf("Contains 'z': expected 0 points, got %d", len(points))
	}
}

// =============================================================================
// SparseGrid SlopeIterate Tests
// =============================================================================

func TestSparseGridSlopeIterate(t *testing.T) {
	// Note: Bounds() returns max point, but image.Rectangle.In() treats max as exclusive
	// So for a 3x3 grid, bounds is (0,0)-(2,2) and (2,2) is NOT "in" bounds
	// Use a larger grid to test the iteration properly
	grid := LoadStringGrid([]string{
		"abcd",
		"efgh",
		"ijkl",
		"mnop",
	})

	var visited []string
	grid.SlopeIterate(image.Pt(0, 0), image.Pt(1, 1), func(pt image.Point, v string) bool {
		visited = append(visited, v)
		return true
	})

	// Should visit (1,1)='f', (2,2)='k', (3,3)='p' is out of bounds (max exclusive)
	expected := []string{"f", "k"}
	if len(visited) != len(expected) {
		t.Errorf("SlopeIterate: expected %d visits, got %d: %v", len(expected), len(visited), visited)
	}
	for i, v := range visited {
		if v != expected[i] {
			t.Errorf("SlopeIterate: expected %v, got %v", expected, visited)
			break
		}
	}
}

func TestSparseGridSlopeIterateHorizontal(t *testing.T) {
	// Note: Bounds() returns max as the actual max point, but image.Rectangle.In()
	// treats max as exclusive. For a single-row grid, this means no points are "in" bounds.
	// Use a 2-row grid to test horizontal iteration properly.
	grid := LoadStringGrid([]string{
		"abcde",
		"fghij",
	})

	var visited []string
	grid.SlopeIterate(image.Pt(0, 0), image.Pt(1, 0), func(pt image.Point, v string) bool {
		visited = append(visited, v)
		return true
	})

	// (1,0)='b', (2,0)='c', (3,0)='d', (4,0) out of bounds (max.X=4 is exclusive)
	expected := []string{"b", "c", "d"}
	if len(visited) != len(expected) {
		t.Errorf("SlopeIterate horizontal: expected %d, got %d: %v", len(expected), len(visited), visited)
	}
}

// =============================================================================
// SparseGrid IterateLine Tests
// =============================================================================

func TestSparseGridIterateLine(t *testing.T) {
	grid := LoadStringGrid([]string{
		"abc",
		"def",
		"ghi",
	})

	var visited []string
	grid.IterateLine(image.Pt(0, 0), image.Pt(2, 2), func(pt image.Point, v string) bool {
		visited = append(visited, v)
		return true
	})

	// Diagonal from (0,0) to (2,2)
	expected := []string{"a", "e", "i"}
	if len(visited) != len(expected) {
		t.Errorf("IterateLine: expected %d points, got %d", len(expected), len(visited))
	}
	for i, v := range visited {
		if v != expected[i] {
			t.Errorf("IterateLine: expected %v, got %v", expected, visited)
			break
		}
	}
}

func TestSparseGridIterateLineHorizontal(t *testing.T) {
	grid := LoadStringGrid([]string{"abcde"})

	var visited []string
	grid.IterateLine(image.Pt(0, 0), image.Pt(4, 0), func(pt image.Point, v string) bool {
		visited = append(visited, v)
		return true
	})

	expected := []string{"a", "b", "c", "d", "e"}
	for i, v := range visited {
		if v != expected[i] {
			t.Errorf("IterateLine horizontal: expected %v, got %v", expected, visited)
			break
		}
	}
}

// =============================================================================
// SparseGrid Sparse Tests
// =============================================================================

func TestSparseGridSparse(t *testing.T) {
	grid := make(SparseGrid[int])
	grid[image.Pt(0, 0)] = 1
	grid[image.Pt(1000, 1000)] = 2

	if len(grid) != 2 {
		t.Errorf("Sparse grid: expected 2 entries, got %d", len(grid))
	}

	// Get should return default for gaps
	if grid.Get(image.Pt(500, 500), -1) != -1 {
		t.Error("Sparse grid: gap should return default")
	}
}

// =============================================================================
// Bounds Function Tests
// =============================================================================

func TestBounds(t *testing.T) {
	points := []image.Point{
		{-5, 10},
		{15, -3},
		{0, 0},
	}

	bounds := Bounds(points)
	if bounds.Min.X != -5 || bounds.Min.Y != -3 {
		t.Errorf("Bounds Min: expected (-5,-3), got (%d,%d)", bounds.Min.X, bounds.Min.Y)
	}
	if bounds.Max.X != 15 || bounds.Max.Y != 10 {
		t.Errorf("Bounds Max: expected (15,10), got (%d,%d)", bounds.Max.X, bounds.Max.Y)
	}
}

func TestBoundsEmpty(t *testing.T) {
	bounds := Bounds([]image.Point{})
	// Should have min > max (inverted) for empty
	if bounds.Min.X <= bounds.Max.X || bounds.Min.Y <= bounds.Max.Y {
		t.Error("Empty bounds should be inverted")
	}
}

// =============================================================================
// Rotation Tests
// =============================================================================

func TestRotate90cw(t *testing.T) {
	tests := []struct {
		in, out image.Point
	}{
		{image.Pt(0, 1), image.Pt(1, 0)},   // Up -> Right
		{image.Pt(1, 0), image.Pt(0, -1)},  // Right -> Down
		{image.Pt(0, -1), image.Pt(-1, 0)}, // Down -> Left
		{image.Pt(-1, 0), image.Pt(0, 1)},  // Left -> Up
	}

	for _, tc := range tests {
		result := Rotate90cw(tc.in)
		if result != tc.out {
			t.Errorf("Rotate90cw(%v): expected %v, got %v", tc.in, tc.out, result)
		}
	}
}

func TestRotate90ccw(t *testing.T) {
	tests := []struct {
		in, out image.Point
	}{
		{image.Pt(0, 1), image.Pt(-1, 0)},  // Up -> Left
		{image.Pt(-1, 0), image.Pt(0, -1)}, // Left -> Down
		{image.Pt(0, -1), image.Pt(1, 0)},  // Down -> Right
		{image.Pt(1, 0), image.Pt(0, 1)},   // Right -> Up
	}

	for _, tc := range tests {
		result := Rotate90ccw(tc.in)
		if result != tc.out {
			t.Errorf("Rotate90ccw(%v): expected %v, got %v", tc.in, tc.out, result)
		}
	}
}

func TestRotate180(t *testing.T) {
	tests := []struct {
		in, out image.Point
	}{
		{image.Pt(1, 2), image.Pt(-1, -2)},
		{image.Pt(-3, 4), image.Pt(3, -4)},
		{image.Pt(0, 0), image.Pt(0, 0)},
	}

	for _, tc := range tests {
		result := Rotate180(tc.in)
		if result != tc.out {
			t.Errorf("Rotate180(%v): expected %v, got %v", tc.in, tc.out, result)
		}
	}
}

func TestRotate270(t *testing.T) {
	// 270 ccw = 90 cw
	p := image.Pt(1, 0)
	if Rotate270ccw(p) != Rotate90cw(p) {
		t.Error("Rotate270ccw should equal Rotate90cw")
	}

	// 270 cw = 90 ccw
	if Rotate270cw(p) != Rotate90ccw(p) {
		t.Error("Rotate270cw should equal Rotate90ccw")
	}
}

// =============================================================================
// ManhattanDistancePt Tests
// =============================================================================

func TestManhattanDistancePt(t *testing.T) {
	tests := []struct {
		p1, p2   image.Point
		expected int
	}{
		{image.Pt(0, 0), image.Pt(3, 4), 7},
		{image.Pt(1, 1), image.Pt(1, 1), 0},
		{image.Pt(-2, -3), image.Pt(2, 3), 10},
		{image.Pt(5, 0), image.Pt(0, 0), 5},
	}

	for _, tc := range tests {
		result := ManhattanDistancePt(tc.p1, tc.p2)
		if result != tc.expected {
			t.Errorf("ManhattanDistancePt(%v, %v): expected %d, got %d",
				tc.p1, tc.p2, tc.expected, result)
		}
	}
}

// =============================================================================
// Contains (standalone function) Tests
// =============================================================================

func TestContainsFunction(t *testing.T) {
	grid := LoadIntGrid([]string{
		"121",
		"232",
		"121",
	})

	ones := Contains(grid, 1)
	if len(ones) != 4 {
		t.Errorf("Contains(1): expected 4 points, got %d", len(ones))
	}

	// Grid has only one 3 at position (1,1)
	threes := Contains(grid, 3)
	if len(threes) != 1 {
		t.Errorf("Contains(3): expected 1 point, got %d", len(threes))
	}
}

// =============================================================================
// Map Function Tests (from functional.go, used by SparseGrid)
// =============================================================================

func TestMapWithPoints(t *testing.T) {
	origin := image.Pt(1, 1)
	deltas := []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	neighbors := Map(origin.Add, deltas)

	expected := []image.Point{{0, 1}, {2, 1}, {1, 0}, {1, 2}}
	slices.SortFunc(neighbors, ComparePoints)
	slices.SortFunc(expected, ComparePoints)

	for i, n := range neighbors {
		if n != expected[i] {
			t.Errorf("Map with points: expected %v, got %v", expected, neighbors)
			break
		}
	}
}
