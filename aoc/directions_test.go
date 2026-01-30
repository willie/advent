package aoc

import (
	"image"
	"testing"
)

func TestDirectionConstants(t *testing.T) {
	// Verify cardinal directions
	if Up != image.Pt(0, -1) {
		t.Error("Up should be (0, -1)")
	}
	if Down != image.Pt(0, 1) {
		t.Error("Down should be (0, 1)")
	}
	if Left != image.Pt(-1, 0) {
		t.Error("Left should be (-1, 0)")
	}
	if Right != image.Pt(1, 0) {
		t.Error("Right should be (1, 0)")
	}

	// Verify aliases
	if North != Up || South != Down || West != Left || East != Right {
		t.Error("Cardinal aliases wrong")
	}
}

func TestDiagonalConstants(t *testing.T) {
	if UpLeft != image.Pt(-1, -1) {
		t.Error("UpLeft wrong")
	}
	if DownRight != image.Pt(1, 1) {
		t.Error("DownRight wrong")
	}
}

func TestDirectionSlices(t *testing.T) {
	if len(FourWay) != 4 {
		t.Errorf("FourWay: expected 4, got %d", len(FourWay))
	}
	if len(EightWay) != 8 {
		t.Errorf("EightWay: expected 8, got %d", len(EightWay))
	}
	if len(Diagonals) != 4 {
		t.Errorf("Diagonals: expected 4, got %d", len(Diagonals))
	}
}

func TestDirFromChar(t *testing.T) {
	tests := []struct {
		char     rune
		expected image.Point
	}{
		{'U', Up},
		{'D', Down},
		{'L', Left},
		{'R', Right},
		{'N', Up},
		{'S', Down},
		{'E', Right},
		{'W', Left},
		{'^', Up},
		{'v', Down},
		{'<', Left},
		{'>', Right},
		{'X', image.Point{}}, // Unknown
	}

	for _, tc := range tests {
		result := DirFromChar(tc.char)
		if result != tc.expected {
			t.Errorf("DirFromChar(%c): expected %v, got %v", tc.char, tc.expected, result)
		}
	}
}

func TestDirFromString(t *testing.T) {
	tests := []struct {
		str      string
		expected image.Point
	}{
		{"up", Up},
		{"UP", Up},
		{"Up", Up},
		{"down", Down},
		{"left", Left},
		{"right", Right},
		{"north", Up},
		{"SOUTH", Down},
		{"East", Right},
		{"west", Left},
		{"U", Up},
		{"", image.Point{}},
		{"unknown", image.Point{}},
	}

	for _, tc := range tests {
		result := DirFromString(tc.str)
		if result != tc.expected {
			t.Errorf("DirFromString(%q): expected %v, got %v", tc.str, tc.expected, result)
		}
	}
}

func TestTurnLeft(t *testing.T) {
	// Turning left from each direction
	if TurnLeft(Up) != Left {
		t.Error("TurnLeft(Up) should be Left")
	}
	if TurnLeft(Left) != Down {
		t.Error("TurnLeft(Left) should be Down")
	}
	if TurnLeft(Down) != Right {
		t.Error("TurnLeft(Down) should be Right")
	}
	if TurnLeft(Right) != Up {
		t.Error("TurnLeft(Right) should be Up")
	}
}

func TestTurnRight(t *testing.T) {
	if TurnRight(Up) != Right {
		t.Error("TurnRight(Up) should be Right")
	}
	if TurnRight(Right) != Down {
		t.Error("TurnRight(Right) should be Down")
	}
	if TurnRight(Down) != Left {
		t.Error("TurnRight(Down) should be Left")
	}
	if TurnRight(Left) != Up {
		t.Error("TurnRight(Left) should be Up")
	}
}

func TestTurnAround(t *testing.T) {
	if TurnAround(Up) != Down {
		t.Error("TurnAround(Up) should be Down")
	}
	if TurnAround(Left) != Right {
		t.Error("TurnAround(Left) should be Right")
	}
}

func TestTurnConsistency(t *testing.T) {
	// Four left turns = back to start
	dir := Up
	for i := 0; i < 4; i++ {
		dir = TurnLeft(dir)
	}
	if dir != Up {
		t.Error("Four left turns should return to start")
	}

	// Left + Right = no change
	if TurnRight(TurnLeft(Up)) != Up {
		t.Error("Left then Right should cancel out")
	}
}

func TestNeighbors4(t *testing.T) {
	center := image.Pt(5, 5)
	neighbors := Neighbors4(center)

	if len(neighbors) != 4 {
		t.Errorf("Neighbors4: expected 4, got %d", len(neighbors))
	}

	// All should be Manhattan distance 1
	for _, n := range neighbors {
		dist := Abs(n.X-center.X) + Abs(n.Y-center.Y)
		if dist != 1 {
			t.Errorf("Neighbors4: neighbor %v should be distance 1", n)
		}
	}
}

func TestNeighbors8(t *testing.T) {
	center := image.Pt(5, 5)
	neighbors := Neighbors8(center)

	if len(neighbors) != 8 {
		t.Errorf("Neighbors8: expected 8, got %d", len(neighbors))
	}

	// All should be Chebyshev distance 1
	for _, n := range neighbors {
		dx := Abs(n.X - center.X)
		dy := Abs(n.Y - center.Y)
		if dx > 1 || dy > 1 || (dx == 0 && dy == 0) {
			t.Errorf("Neighbors8: invalid neighbor %v", n)
		}
	}
}

func TestDirectionsWithGrid(t *testing.T) {
	// Common AoC pattern: move in direction
	pos := image.Pt(0, 0)
	moves := []rune{'R', 'R', 'D', 'D', 'L', 'U'}

	for _, m := range moves {
		pos = pos.Add(DirFromChar(m))
	}

	expected := image.Pt(1, 1)
	if pos != expected {
		t.Errorf("Move sequence: expected %v, got %v", expected, pos)
	}
}
