package aoc

import "testing"

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
