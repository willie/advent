package aoc

import "testing"

// =============================================================================
// Ints Creation Tests
// =============================================================================

func TestStringInts(t *testing.T) {
	ints := StringInts([]string{"1", "2", "3"})

	if len(ints) != 3 {
		t.Errorf("StringInts: expected 3 elements, got %d", len(ints))
	}
	if ints[0] != 1 || ints[1] != 2 || ints[2] != 3 {
		t.Errorf("StringInts: expected [1,2,3], got %v", ints)
	}
}

func TestStringIntsNegative(t *testing.T) {
	ints := StringInts([]string{"-5", "10", "-3"})

	if ints[0] != -5 || ints[1] != 10 || ints[2] != -3 {
		t.Errorf("StringInts negative: expected [-5,10,-3], got %v", ints)
	}
}

// =============================================================================
// Ints.Sum Tests
// =============================================================================

func TestIntsSum(t *testing.T) {
	ints := Ints{1, 2, 3, 4, 5}
	if ints.Sum() != 15 {
		t.Errorf("Sum: expected 15, got %d", ints.Sum())
	}
}

func TestIntsSumEmpty(t *testing.T) {
	ints := Ints{}
	if ints.Sum() != 0 {
		t.Errorf("Sum empty: expected 0, got %d", ints.Sum())
	}
}

func TestIntsSumNegative(t *testing.T) {
	ints := Ints{-5, 10, -3}
	if ints.Sum() != 2 {
		t.Errorf("Sum negative: expected 2, got %d", ints.Sum())
	}
}

// =============================================================================
// Ints.Product Tests
// =============================================================================

func TestIntsProduct(t *testing.T) {
	ints := Ints{2, 3, 4}
	if ints.Product() != 24 {
		t.Errorf("Product: expected 24, got %d", ints.Product())
	}
}

func TestIntsProductEmpty(t *testing.T) {
	ints := Ints{}
	if ints.Product() != 1 {
		t.Errorf("Product empty: expected 1, got %d", ints.Product())
	}
}

func TestIntsProductWithZero(t *testing.T) {
	ints := Ints{5, 0, 3}
	if ints.Product() != 0 {
		t.Errorf("Product with zero: expected 0, got %d", ints.Product())
	}
}

// =============================================================================
// Ints.Min/Max Tests
// =============================================================================

func TestIntsMin(t *testing.T) {
	ints := Ints{5, 2, 8, 1, 9}
	if ints.Min() != 1 {
		t.Errorf("Min: expected 1, got %d", ints.Min())
	}
}

func TestIntsMinSingle(t *testing.T) {
	ints := Ints{42}
	if ints.Min() != 42 {
		t.Errorf("Min single: expected 42, got %d", ints.Min())
	}
}

func TestIntsMinNegative(t *testing.T) {
	ints := Ints{-5, -2, -8}
	if ints.Min() != -8 {
		t.Errorf("Min negative: expected -8, got %d", ints.Min())
	}
}

func TestIntsMax(t *testing.T) {
	ints := Ints{5, 2, 8, 1, 9}
	if ints.Max() != 9 {
		t.Errorf("Max: expected 9, got %d", ints.Max())
	}
}

func TestIntsMaxSingle(t *testing.T) {
	ints := Ints{42}
	if ints.Max() != 42 {
		t.Errorf("Max single: expected 42, got %d", ints.Max())
	}
}

func TestIntsMaxNegative(t *testing.T) {
	ints := Ints{-5, -2, -8}
	if ints.Max() != -2 {
		t.Errorf("Max negative: expected -2, got %d", ints.Max())
	}
}

// =============================================================================
// Ints.Last Tests
// =============================================================================

func TestIntsLast(t *testing.T) {
	ints := Ints{1, 2, 3, 4, 5}
	if ints.Last() != 5 {
		t.Errorf("Last: expected 5, got %d", ints.Last())
	}
}

func TestIntsLastSingle(t *testing.T) {
	ints := Ints{42}
	if ints.Last() != 42 {
		t.Errorf("Last single: expected 42, got %d", ints.Last())
	}
}

// =============================================================================
// Ints.Index Tests
// =============================================================================

func TestIntsIndex(t *testing.T) {
	ints := Ints{10, 20, 30, 20, 40}

	if ints.Index(20) != 1 {
		t.Errorf("Index: expected 1, got %d", ints.Index(20))
	}
	if ints.Index(40) != 4 {
		t.Errorf("Index: expected 4, got %d", ints.Index(40))
	}
}

func TestIntsIndexNotFound(t *testing.T) {
	ints := Ints{1, 2, 3}
	if ints.Index(99) != -1 {
		t.Errorf("Index not found: expected -1, got %d", ints.Index(99))
	}
}

func TestIntsIndexFirst(t *testing.T) {
	ints := Ints{5, 5, 5}
	if ints.Index(5) != 0 {
		t.Errorf("Index first: expected 0 (first occurrence), got %d", ints.Index(5))
	}
}

// =============================================================================
// Ints.LastIndex Tests
// =============================================================================

func TestIntsLastIndex(t *testing.T) {
	ints := Ints{10, 20, 30, 20, 40}

	if ints.LastIndex(20) != 3 {
		t.Errorf("LastIndex: expected 3, got %d", ints.LastIndex(20))
	}
}

func TestIntsLastIndexNotFound(t *testing.T) {
	ints := Ints{1, 2, 3}
	if ints.LastIndex(99) != -1 {
		t.Errorf("LastIndex not found: expected -1, got %d", ints.LastIndex(99))
	}
}

func TestIntsLastIndexSingle(t *testing.T) {
	ints := Ints{5}
	if ints.LastIndex(5) != 0 {
		t.Errorf("LastIndex single: expected 0, got %d", ints.LastIndex(5))
	}
}

// =============================================================================
// Ints.AllIndex Tests
// =============================================================================

func TestIntsAllIndex(t *testing.T) {
	ints := Ints{1, 2, 1, 3, 1, 4}
	indices := ints.AllIndex(1)

	expected := []int{0, 2, 4}
	if len(indices) != len(expected) {
		t.Errorf("AllIndex: expected %d indices, got %d", len(expected), len(indices))
	}
	for i, idx := range indices {
		if idx != expected[i] {
			t.Errorf("AllIndex: expected %v, got %v", expected, indices)
			break
		}
	}
}

func TestIntsAllIndexNotFound(t *testing.T) {
	ints := Ints{1, 2, 3}
	indices := ints.AllIndex(99)

	if len(indices) != 0 {
		t.Errorf("AllIndex not found: expected empty, got %v", indices)
	}
}

func TestIntsAllIndexAll(t *testing.T) {
	ints := Ints{5, 5, 5}
	indices := ints.AllIndex(5)

	if len(indices) != 3 {
		t.Errorf("AllIndex all: expected 3, got %d", len(indices))
	}
}

// =============================================================================
// Series Tests
// =============================================================================

func TestSeries(t *testing.T) {
	series := Series(1, 5)

	expected := Ints{1, 2, 3, 4, 5}
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

	expected := Ints{-2, -1, 0, 1, 2}
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
// Integration with Math Functions
// =============================================================================

func TestIntsWithGenericSum(t *testing.T) {
	ints := Ints{1, 2, 3, 4, 5}

	// Both should give same result
	methodSum := ints.Sum()
	genericSum := Sum(ints...)

	if methodSum != genericSum {
		t.Errorf("Sum consistency: method=%d, generic=%d", methodSum, genericSum)
	}
}

func TestIntsWithGenericMinMax(t *testing.T) {
	ints := Ints{5, 2, 8, 1, 9}

	if ints.Min() != Min(ints...) {
		t.Error("Min consistency: method and generic differ")
	}
	if ints.Max() != Max(ints...) {
		t.Error("Max consistency: method and generic differ")
	}
}
