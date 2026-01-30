package aoc

import "testing"

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
