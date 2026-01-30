package aoc

import (
	"strings"
	"testing"
)

// =============================================================================
// Map Tests
// =============================================================================

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := Map(func(x int) int { return x * 2 }, input)

	expected := []int{2, 4, 6, 8, 10}
	if len(result) != len(expected) {
		t.Errorf("Map: expected %d elements, got %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Map: expected %v, got %v", expected, result)
			break
		}
	}
}

func TestMapEmpty(t *testing.T) {
	result := Map(func(x int) int { return x * 2 }, []int{})

	if len(result) != 0 {
		t.Errorf("Map empty: expected 0 elements, got %d", len(result))
	}
}

func TestMapTypeConversion(t *testing.T) {
	input := []int{1, 2, 3}
	result := Map(func(x int) string { return strings.Repeat("*", x) }, input)

	expected := []string{"*", "**", "***"}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Map type conversion: expected %v, got %v", expected, result)
			break
		}
	}
}

func TestMapWithStrings(t *testing.T) {
	input := []string{"hello", "world"}
	result := Map(strings.ToUpper, input)

	expected := []string{"HELLO", "WORLD"}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Map strings: expected %v, got %v", expected, result)
			break
		}
	}
}

func TestMapPreservesOrder(t *testing.T) {
	input := []int{5, 3, 8, 1, 9}
	result := Map(func(x int) int { return x }, input)

	for i, v := range result {
		if v != input[i] {
			t.Error("Map should preserve order")
			break
		}
	}
}

// =============================================================================
// Filter Tests
// =============================================================================

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := Filter(func(x int) bool { return x%2 == 0 }, input)

	expected := []int{2, 4, 6, 8, 10}
	if len(result) != len(expected) {
		t.Errorf("Filter: expected %d elements, got %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Filter: expected %v, got %v", expected, result)
			break
		}
	}
}

func TestFilterEmpty(t *testing.T) {
	result := Filter(func(x int) bool { return true }, []int{})

	if len(result) != 0 {
		t.Errorf("Filter empty: expected 0 elements, got %d", len(result))
	}
}

func TestFilterNone(t *testing.T) {
	input := []int{1, 2, 3}
	result := Filter(func(x int) bool { return false }, input)

	if len(result) != 0 {
		t.Errorf("Filter none: expected 0 elements, got %d", len(result))
	}
}

func TestFilterAll(t *testing.T) {
	input := []int{1, 2, 3}
	result := Filter(func(x int) bool { return true }, input)

	if len(result) != 3 {
		t.Errorf("Filter all: expected 3 elements, got %d", len(result))
	}
}

func TestFilterStrings(t *testing.T) {
	input := []string{"apple", "banana", "apricot", "cherry"}
	result := Filter(func(s string) bool { return strings.HasPrefix(s, "a") }, input)

	expected := []string{"apple", "apricot"}
	if len(result) != len(expected) {
		t.Errorf("Filter strings: expected %d elements, got %d", len(expected), len(result))
	}
}

func TestFilterPreservesOrder(t *testing.T) {
	input := []int{1, 3, 5, 7, 9}
	result := Filter(func(x int) bool { return x > 2 }, input)

	expected := []int{3, 5, 7, 9}
	for i, v := range result {
		if v != expected[i] {
			t.Error("Filter should preserve order")
			break
		}
	}
}

// =============================================================================
// FilterMap Tests
// =============================================================================

func TestFilterMap(t *testing.T) {
	input := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	result := FilterMap(input, func(k string, v int) bool {
		return v%2 == 0
	})

	if len(result) != 2 {
		t.Errorf("FilterMap: expected 2 elements, got %d", len(result))
	}
	if result["b"] != 2 || result["d"] != 4 {
		t.Errorf("FilterMap: expected {b:2, d:4}, got %v", result)
	}
}

func TestFilterMapEmpty(t *testing.T) {
	input := map[string]int{}
	result := FilterMap(input, func(k string, v int) bool { return true })

	if len(result) != 0 {
		t.Errorf("FilterMap empty: expected 0 elements, got %d", len(result))
	}
}

func TestFilterMapNone(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}
	result := FilterMap(input, func(k string, v int) bool { return false })

	if len(result) != 0 {
		t.Errorf("FilterMap none: expected 0 elements, got %d", len(result))
	}
}

func TestFilterMapByKey(t *testing.T) {
	input := map[string]int{
		"apple":  1,
		"banana": 2,
		"apricot": 3,
	}

	result := FilterMap(input, func(k string, v int) bool {
		return strings.HasPrefix(k, "a")
	})

	if len(result) != 2 {
		t.Errorf("FilterMap by key: expected 2 elements, got %d", len(result))
	}
}

func TestFilterMapByBoth(t *testing.T) {
	input := map[string]int{
		"a": 10,
		"b": 5,
		"c": 20,
	}

	result := FilterMap(input, func(k string, v int) bool {
		return k != "b" && v > 5
	})

	if len(result) != 2 {
		t.Errorf("FilterMap by both: expected 2 elements, got %d", len(result))
	}
}

// =============================================================================
// Map + Filter Composition
// =============================================================================

func TestMapThenFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	// Double all, then filter evens
	doubled := Map(func(x int) int { return x * 2 }, input)
	result := Filter(func(x int) bool { return x > 5 }, doubled)

	expected := []int{6, 8, 10}
	if len(result) != len(expected) {
		t.Errorf("Map then Filter: expected %d elements, got %d", len(expected), len(result))
	}
}

func TestFilterThenMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	// Filter evens, then double
	evens := Filter(func(x int) bool { return x%2 == 0 }, input)
	result := Map(func(x int) int { return x * 2 }, evens)

	expected := []int{4, 8}
	if len(result) != len(expected) {
		t.Errorf("Filter then Map: expected %d elements, got %d", len(expected), len(result))
	}
}

// =============================================================================
// Edge Cases
// =============================================================================

func TestMapWithNil(t *testing.T) {
	var input []int
	result := Map(func(x int) int { return x * 2 }, input)

	if result == nil {
		t.Error("Map with nil should return empty slice, not nil")
	}
	if len(result) != 0 {
		t.Errorf("Map with nil: expected 0 elements, got %d", len(result))
	}
}

func TestFilterWithNil(t *testing.T) {
	var input []int
	result := Filter(func(x int) bool { return true }, input)

	// Filter returns nil for nil input (unlike Map)
	if len(result) != 0 {
		t.Errorf("Filter with nil: expected 0 elements, got %d", len(result))
	}
}

func TestMapLargeSlice(t *testing.T) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}

	result := Map(func(x int) int { return x + 1 }, input)

	if len(result) != 10000 {
		t.Errorf("Map large: expected 10000 elements, got %d", len(result))
	}
	if result[9999] != 10000 {
		t.Errorf("Map large: last element should be 10000, got %d", result[9999])
	}
}
