package aoc

import (
	"image"
	"math"
	"testing"
)

// =============================================================================
// Sum Tests
// =============================================================================

func TestSum(t *testing.T) {
	if Sum(1, 2, 3, 4, 5) != 15 {
		t.Error("Sum: 1+2+3+4+5 should be 15")
	}
	if Sum[int]() != 0 {
		t.Error("Sum empty: should be 0")
	}
	if Sum(-5, 10, -3) != 2 {
		t.Error("Sum negative: -5+10-3 should be 2")
	}
}

func TestSumInt64(t *testing.T) {
	var a, b int64 = 1<<40, 1<<40
	result := Sum(a, b)
	if result != 1<<41 {
		t.Errorf("Sum int64: expected %d, got %d", int64(1<<41), result)
	}
}

// =============================================================================
// Product Tests
// =============================================================================

func TestProduct(t *testing.T) {
	if Product(2, 3, 4) != 24 {
		t.Error("Product: 2*3*4 should be 24")
	}
	if Product[int]() != 1 {
		t.Error("Product empty: should be 1")
	}
	if Product(5) != 5 {
		t.Error("Product single: should be 5")
	}
	if Product(-2, 3) != -6 {
		t.Error("Product negative: -2*3 should be -6")
	}
}

// =============================================================================
// GCD Tests
// =============================================================================

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{12, 8, 4},
		{100, 35, 5},
		{17, 13, 1}, // Coprime
		{0, 5, 5},
		{5, 0, 5},
		{48, 18, 6},
	}

	for _, tc := range tests {
		result := GCD(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("GCD(%d, %d): expected %d, got %d", tc.a, tc.b, tc.expected, result)
		}
	}
}

// =============================================================================
// LCM Tests
// =============================================================================

func TestLCM(t *testing.T) {
	if LCM(4, 6) != 12 {
		t.Errorf("LCM(4,6): expected 12, got %d", LCM(4, 6))
	}
	if LCM(3, 5) != 15 {
		t.Errorf("LCM(3,5): expected 15, got %d", LCM(3, 5))
	}
	if LCM(2, 3, 4) != 12 {
		t.Errorf("LCM(2,3,4): expected 12, got %d", LCM(2, 3, 4))
	}
	if LCM(6) != 6 {
		t.Errorf("LCM(6): expected 6, got %d", LCM(6))
	}
}

func TestLCMEmpty(t *testing.T) {
	result := LCM[int]()
	if result != 0 {
		t.Errorf("LCM empty: expected 0, got %d", result)
	}
}

// =============================================================================
// Abs Tests
// =============================================================================

func TestAbs(t *testing.T) {
	if Abs(5) != 5 {
		t.Error("Abs positive: should stay positive")
	}
	if Abs(-5) != 5 {
		t.Error("Abs negative: should become positive")
	}
	if Abs(0) != 0 {
		t.Error("Abs zero: should be zero")
	}
}

func TestAbsInt64(t *testing.T) {
	var n int64 = -1 << 40
	if Abs(n) != 1<<40 {
		t.Error("Abs int64: should work with large numbers")
	}
}

// =============================================================================
// ManhattanDistance Tests
// =============================================================================

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		x, y, x1, y1, expected int
	}{
		{0, 0, 3, 4, 7},
		{1, 1, 1, 1, 0},
		{-2, -3, 2, 3, 10},
		{0, 0, 0, 5, 5},
		{0, 0, 5, 0, 5},
	}

	for _, tc := range tests {
		result := ManhattanDistance(tc.x, tc.y, tc.x1, tc.y1)
		if result != tc.expected {
			t.Errorf("ManhattanDistance(%d,%d,%d,%d): expected %d, got %d",
				tc.x, tc.y, tc.x1, tc.y1, tc.expected, result)
		}
	}
}

// =============================================================================
// Permutations Tests
// =============================================================================

func TestPermutations(t *testing.T) {
	perms := Permutations([]int{1, 2, 3})

	if len(perms) != 6 { // 3! = 6
		t.Errorf("Permutations [1,2,3]: expected 6, got %d", len(perms))
	}

	// Verify all permutations are unique
	seen := make(map[string]bool)
	for _, p := range perms {
		key := ""
		for _, v := range p {
			key += string(rune('0' + v))
		}
		if seen[key] {
			t.Errorf("Permutations: duplicate found %v", p)
		}
		seen[key] = true
	}
}

func TestPermutationsEmpty(t *testing.T) {
	perms := Permutations([]int{})
	if len(perms) != 0 {
		t.Errorf("Permutations empty: expected 0, got %d", len(perms))
	}
}

func TestPermutationsSingle(t *testing.T) {
	perms := Permutations([]int{42})
	if len(perms) != 1 {
		t.Errorf("Permutations single: expected 1, got %d", len(perms))
	}
	if perms[0][0] != 42 {
		t.Errorf("Permutations single: expected [42], got %v", perms[0])
	}
}

// =============================================================================
// AngleDistance Tests
// =============================================================================

func TestAngleDistance(t *testing.T) {
	// AngleDistance was designed for a specific AoC problem (asteroid detection)
	// The angle calculation uses a particular coordinate system
	// We primarily test the distance calculation here

	// Test distance (3-4-5 triangle)
	_, dist := AngleDistance(image.Pt(0, 0), image.Pt(3, 4))
	if math.Abs(dist-5) > 0.001 {
		t.Errorf("AngleDistance: expected distance 5, got %f", dist)
	}

	// Test distance of 0
	_, dist = AngleDistance(image.Pt(5, 5), image.Pt(5, 5))
	if dist != 0 {
		t.Errorf("AngleDistance same point: expected 0, got %f", dist)
	}

	// Test that angle is consistent for same relative position
	angle1, _ := AngleDistance(image.Pt(0, 0), image.Pt(1, 1))
	angle2, _ := AngleDistance(image.Pt(5, 5), image.Pt(6, 6))
	if math.Abs(angle1-angle2) > 0.001 {
		t.Errorf("AngleDistance consistency: angles differ %f vs %f", angle1, angle2)
	}
}

// =============================================================================
// Distance Tests
// =============================================================================

func TestDistance(t *testing.T) {
	tests := []struct {
		p, q     image.Point
		expected float64
	}{
		{image.Pt(0, 0), image.Pt(3, 4), 5.0},
		{image.Pt(1, 1), image.Pt(1, 1), 0.0},
		{image.Pt(0, 0), image.Pt(1, 0), 1.0},
		{image.Pt(0, 0), image.Pt(0, 1), 1.0},
	}

	for _, tc := range tests {
		result := Distance(tc.p, tc.q)
		if math.Abs(result-tc.expected) > 0.001 {
			t.Errorf("Distance(%v, %v): expected %f, got %f",
				tc.p, tc.q, tc.expected, result)
		}
	}
}

// =============================================================================
// ComparePoints / LessThan Tests
// =============================================================================

func TestComparePoints(t *testing.T) {
	tests := []struct {
		a, b     image.Point
		expected int
	}{
		{image.Pt(0, 0), image.Pt(1, 1), -1}, // a < b (by Y, then X)
		{image.Pt(1, 1), image.Pt(0, 0), 1},  // a > b
		{image.Pt(1, 1), image.Pt(1, 1), 0},  // equal
		{image.Pt(0, 1), image.Pt(1, 0), 1},  // a.Y > b.Y
		{image.Pt(1, 0), image.Pt(2, 0), -1}, // same Y, a.X < b.X
	}

	for _, tc := range tests {
		result := ComparePoints(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("ComparePoints(%v, %v): expected %d, got %d",
				tc.a, tc.b, tc.expected, result)
		}
	}
}

func TestLessThan(t *testing.T) {
	if !LessThan(image.Pt(0, 0), image.Pt(1, 1)) {
		t.Error("LessThan: (0,0) should be less than (1,1)")
	}
	if LessThan(image.Pt(1, 1), image.Pt(0, 0)) {
		t.Error("LessThan: (1,1) should not be less than (0,0)")
	}
	if LessThan(image.Pt(1, 1), image.Pt(1, 1)) {
		t.Error("LessThan: equal points should return false")
	}
}
