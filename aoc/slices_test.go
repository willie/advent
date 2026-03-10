package aoc

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := Reverse(input)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Reverse: expected %v, got %v", expected, result)
			break
		}
	}

	// Original should be unchanged
	if input[0] != 1 {
		t.Error("Reverse: should not modify original")
	}
}

func TestChunk(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	chunks := Chunk(input, 3)

	if len(chunks) != 3 {
		t.Errorf("Chunk: expected 3 chunks, got %d", len(chunks))
	}
	if len(chunks[0]) != 3 || len(chunks[1]) != 3 || len(chunks[2]) != 1 {
		t.Errorf("Chunk: wrong chunk sizes")
	}
}

func TestChunkExact(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	chunks := Chunk(input, 2)

	if len(chunks) != 3 {
		t.Errorf("Chunk exact: expected 3 chunks, got %d", len(chunks))
	}
}

func TestChunkEdgeCases(t *testing.T) {
	if Chunk([]int{1, 2, 3}, 0) != nil {
		t.Error("Chunk with n=0 should return nil")
	}
	if Chunk([]int{}, 2) != nil {
		t.Error("Chunk empty slice should return nil")
	}
}

func TestWindow(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	windows := Window(input, 3)

	if len(windows) != 3 {
		t.Errorf("Window: expected 3 windows, got %d", len(windows))
	}
	// [1,2,3], [2,3,4], [3,4,5]
	if windows[0][0] != 1 || windows[1][0] != 2 || windows[2][0] != 3 {
		t.Errorf("Window: wrong values")
	}
}

func TestWindowEdgeCases(t *testing.T) {
	if Window([]int{1, 2}, 3) != nil {
		t.Error("Window with n > len should return nil")
	}
	if Window([]int{1, 2, 3}, 0) != nil {
		t.Error("Window with n=0 should return nil")
	}
}

func TestPairs(t *testing.T) {
	input := []int{1, 2, 3, 4}
	pairs := Pairs(input)

	if len(pairs) != 3 {
		t.Errorf("Pairs: expected 3 pairs, got %d", len(pairs))
	}
	if pairs[0] != [2]int{1, 2} || pairs[1] != [2]int{2, 3} || pairs[2] != [2]int{3, 4} {
		t.Errorf("Pairs: wrong values")
	}
}

func TestPairsEdgeCases(t *testing.T) {
	if Pairs([]int{1}) != nil {
		t.Error("Pairs with single element should return nil")
	}
	if Pairs([]int{}) != nil {
		t.Error("Pairs with empty slice should return nil")
	}
}

func TestCount(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	count := Count(input, func(n int) bool { return n%2 == 0 })

	if count != 3 {
		t.Errorf("Count: expected 3, got %d", count)
	}
}

func TestAll(t *testing.T) {
	positive := []int{1, 2, 3, 4, 5}
	mixed := []int{1, -2, 3}

	if !All(positive, func(n int) bool { return n > 0 }) {
		t.Error("All: should be true for all positive")
	}
	if All(mixed, func(n int) bool { return n > 0 }) {
		t.Error("All: should be false for mixed")
	}
	if !All([]int{}, func(n int) bool { return false }) {
		t.Error("All: empty slice should return true")
	}
}

func TestAny(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	if !Any(input, func(n int) bool { return n == 3 }) {
		t.Error("Any: should find 3")
	}
	if Any(input, func(n int) bool { return n == 99 }) {
		t.Error("Any: should not find 99")
	}
	if Any([]int{}, func(n int) bool { return true }) {
		t.Error("Any: empty slice should return false")
	}
}

func TestFind(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	val, found := Find(input, func(n int) bool { return n > 3 })
	if !found || val != 4 {
		t.Errorf("Find: expected 4, got %d (found=%v)", val, found)
	}

	_, found = Find(input, func(n int) bool { return n > 10 })
	if found {
		t.Error("Find: should not find value > 10")
	}
}

func TestFindIndex(t *testing.T) {
	input := []int{10, 20, 30, 40}

	idx := FindIndex(input, func(n int) bool { return n == 30 })
	if idx != 2 {
		t.Errorf("FindIndex: expected 2, got %d", idx)
	}

	idx = FindIndex(input, func(n int) bool { return n == 99 })
	if idx != -1 {
		t.Errorf("FindIndex not found: expected -1, got %d", idx)
	}
}

func TestUnique(t *testing.T) {
	input := []int{1, 2, 2, 3, 1, 4, 3}
	result := Unique(input)

	expected := []int{1, 2, 3, 4}
	if len(result) != len(expected) {
		t.Errorf("Unique: expected %v, got %v", expected, result)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Unique: expected %v, got %v", expected, result)
			break
		}
	}
}

func TestFlatten(t *testing.T) {
	input := [][]int{{1, 2}, {3, 4, 5}, {6}}
	result := Flatten(input)

	expected := []int{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Errorf("Flatten: expected %v, got %v", expected, result)
	}
}

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	sum := Reduce(input, 0, func(acc, val int) int { return acc + val })

	if sum != 15 {
		t.Errorf("Reduce sum: expected 15, got %d", sum)
	}

	product := Reduce(input, 1, func(acc, val int) int { return acc * val })
	if product != 120 {
		t.Errorf("Reduce product: expected 120, got %d", product)
	}
}

// =============================================================================
// Series Tests
// =============================================================================

func TestSeries(t *testing.T) {
	series := Series(1, 5)

	expected := []int{1, 2, 3, 4, 5}
	if len(series) != len(expected) {
		t.Errorf("Series: expected %d elements, got %d", len(expected), len(series))
	}
	for i, v := range series {
		if v != expected[i] {
			t.Errorf("Series: expected %v, got %v", expected, series)
			break
		}
	}
}

func TestSeriesNegative(t *testing.T) {
	series := Series(-2, 2)

	expected := []int{-2, -1, 0, 1, 2}
	if len(series) != len(expected) {
		t.Errorf("Series negative: expected %d elements, got %d", len(expected), len(series))
	}
	for i, v := range series {
		if v != expected[i] {
			t.Errorf("Series negative: expected %v, got %v", expected, series)
			break
		}
	}
}

func TestSeriesSingle(t *testing.T) {
	series := Series(5, 5)

	if len(series) != 1 {
		t.Errorf("Series single: expected 1 element, got %d", len(series))
	}
	if series[0] != 5 {
		t.Errorf("Series single: expected [5], got %v", series)
	}
}

func TestSeriesEmpty(t *testing.T) {
	series := Series(5, 3) // high < low

	if series != nil {
		t.Errorf("Series empty: expected nil, got %v", series)
	}
}

func TestSeriesLarge(t *testing.T) {
	series := Series(0, 999)

	if len(series) != 1000 {
		t.Errorf("Series large: expected 1000 elements, got %d", len(series))
	}
	if series[0] != 0 || series[999] != 999 {
		t.Error("Series large: boundary values wrong")
	}
}

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
		"apple":   1,
		"banana":  2,
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
