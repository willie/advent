package aoc

import (
	"slices"
	"testing"
)

// =============================================================================
// Generic Set[T] Tests
// =============================================================================

func TestNewSet(t *testing.T) {
	s := NewSet(1, 2, 3)
	if len(s) != 3 {
		t.Errorf("NewSet: expected 3 elements, got %d", len(s))
	}
	for _, v := range []int{1, 2, 3} {
		if !s.Contains(v) {
			t.Errorf("NewSet: missing value %d", v)
		}
	}
}

func TestNewSetEmpty(t *testing.T) {
	s := NewSet[int]()
	if len(s) != 0 {
		t.Errorf("NewSet empty: expected 0 elements, got %d", len(s))
	}
}

func TestSetAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3)
	if len(s) != 3 {
		t.Errorf("Add: expected 3 elements, got %d", len(s))
	}

	// Adding duplicates should not increase size
	s.Add(1, 2)
	if len(s) != 3 {
		t.Errorf("Add duplicates: expected 3 elements, got %d", len(s))
	}
}

func TestSetAddSlice(t *testing.T) {
	s := NewSet[string]()
	s.AddSlice([]string{"a", "b", "c"})
	if len(s) != 3 {
		t.Errorf("AddSlice: expected 3 elements, got %d", len(s))
	}
}

func TestSetAddSet(t *testing.T) {
	s1 := NewSet(1, 2)
	s2 := NewSet(3, 4)
	s1.AddSet(s2)

	if len(s1) != 4 {
		t.Errorf("AddSet: expected 4 elements, got %d", len(s1))
	}
	for _, v := range []int{1, 2, 3, 4} {
		if !s1.Contains(v) {
			t.Errorf("AddSet: missing value %d", v)
		}
	}
}

func TestSetRemove(t *testing.T) {
	s := NewSet(1, 2, 3, 4)
	s.Remove(2, 4)

	if len(s) != 2 {
		t.Errorf("Remove: expected 2 elements, got %d", len(s))
	}
	if s.Contains(2) || s.Contains(4) {
		t.Error("Remove: values not removed")
	}
	if !s.Contains(1) || !s.Contains(3) {
		t.Error("Remove: wrong values removed")
	}
}

func TestSetRemoveNonExistent(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Remove(99) // Should not panic
	if len(s) != 3 {
		t.Errorf("Remove non-existent: expected 3 elements, got %d", len(s))
	}
}

func TestSetContains(t *testing.T) {
	s := NewSet("a", "b", "c")
	if !s.Contains("a") {
		t.Error("Contains: should contain 'a'")
	}
	if s.Contains("z") {
		t.Error("Contains: should not contain 'z'")
	}
}

func TestSetContainsAll(t *testing.T) {
	s := NewSet(1, 2, 3, 4, 5)

	if !s.ContainsAll([]int{1, 3, 5}) {
		t.Error("ContainsAll: should contain all of [1,3,5]")
	}
	if s.ContainsAll([]int{1, 6}) {
		t.Error("ContainsAll: should not contain all of [1,6]")
	}
	if !s.ContainsAll([]int{}) {
		t.Error("ContainsAll: empty slice should return true")
	}
}

func TestSetContainsAny(t *testing.T) {
	s := NewSet(1, 2, 3)

	if !s.ContainsAny([]int{3, 4, 5}) {
		t.Error("ContainsAny: should contain at least one of [3,4,5]")
	}
	if s.ContainsAny([]int{6, 7, 8}) {
		t.Error("ContainsAny: should not contain any of [6,7,8]")
	}
	if s.ContainsAny([]int{}) {
		t.Error("ContainsAny: empty slice should return false")
	}
}

func TestSetSubtract(t *testing.T) {
	s1 := NewSet(1, 2, 3, 4)
	s2 := NewSet(3, 4, 5)
	diff := s1.Subtract(s2)

	if len(diff) != 2 {
		t.Errorf("Subtract: expected 2 elements, got %d", len(diff))
	}
	if !diff.Contains(1) || !diff.Contains(2) {
		t.Error("Subtract: wrong difference")
	}
	if diff.Contains(3) || diff.Contains(4) {
		t.Error("Subtract: should not contain subtracted values")
	}
}

func TestSetIntersect(t *testing.T) {
	s1 := NewSet(1, 2, 3, 4)
	s2 := NewSet(3, 4, 5, 6)
	inter := s1.Intersect(s2)

	if len(inter) != 2 {
		t.Errorf("Intersect: expected 2 elements, got %d", len(inter))
	}
	if !inter.Contains(3) || !inter.Contains(4) {
		t.Error("Intersect: wrong intersection")
	}
}

func TestSetIntersectEmpty(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(4, 5, 6)
	inter := s1.Intersect(s2)

	if len(inter) != 0 {
		t.Errorf("Intersect disjoint: expected 0 elements, got %d", len(inter))
	}
}

func TestSetUnion(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(3, 4, 5)
	union := s1.Union(s2)

	if len(union) != 5 {
		t.Errorf("Union: expected 5 elements, got %d", len(union))
	}
	for _, v := range []int{1, 2, 3, 4, 5} {
		if !union.Contains(v) {
			t.Errorf("Union: missing value %d", v)
		}
	}
}

func TestSetValues(t *testing.T) {
	s := NewSet(3, 1, 2)
	vals := s.Values()

	if len(vals) != 3 {
		t.Errorf("Values: expected 3 values, got %d", len(vals))
	}

	// Check all values are present (order not guaranteed)
	slices.Sort(vals)
	expected := []int{1, 2, 3}
	for i, v := range vals {
		if v != expected[i] {
			t.Errorf("Values: expected %v, got %v", expected, vals)
			break
		}
	}
}

// =============================================================================
// Edge Cases
// =============================================================================

func TestSetWithStructs(t *testing.T) {
	type Point struct{ X, Y int }

	s := NewSet(Point{1, 2}, Point{3, 4})
	if !s.Contains(Point{1, 2}) {
		t.Error("Set with structs: should contain Point{1,2}")
	}
	if s.Contains(Point{5, 6}) {
		t.Error("Set with structs: should not contain Point{5,6}")
	}
}

func TestSetChaining(t *testing.T) {
	s := NewSet[int]().Add(1).Add(2).Add(3).Remove(2)
	if len(s) != 2 {
		t.Errorf("Chaining: expected 2 elements, got %d", len(s))
	}
	if !s.Contains(1) || !s.Contains(3) {
		t.Error("Chaining: wrong values")
	}
}
