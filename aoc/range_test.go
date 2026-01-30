package aoc

import "testing"

// =============================================================================
// Range Tests
// =============================================================================

func TestRangeContains(t *testing.T) {
	tests := []struct {
		name     string
		r1, r2   Range
		expected bool
	}{
		{"fully contains", Range{1, 10}, Range{3, 7}, true},
		{"exact match", Range{1, 10}, Range{1, 10}, true},
		{"partial overlap left", Range{5, 10}, Range{1, 7}, false},
		{"partial overlap right", Range{1, 5}, Range{3, 10}, false},
		{"no overlap", Range{1, 5}, Range{6, 10}, false},
		{"contains single point", Range{1, 10}, Range{5, 5}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.r1.Contains(tc.r2)
			if result != tc.expected {
				t.Errorf("Range{%d,%d}.Contains(Range{%d,%d}) = %v, want %v",
					tc.r1.Start, tc.r1.End, tc.r2.Start, tc.r2.End, result, tc.expected)
			}
		})
	}
}

func TestRangeOverlaps(t *testing.T) {
	tests := []struct {
		name     string
		r1, r2   Range
		expected bool
	}{
		{"full overlap", Range{1, 10}, Range{3, 7}, true},
		{"partial left", Range{1, 5}, Range{3, 10}, true},
		{"partial right", Range{5, 10}, Range{1, 7}, true},
		{"touching edges", Range{1, 5}, Range{5, 10}, true},
		{"no overlap", Range{1, 5}, Range{6, 10}, false},
		{"same range", Range{1, 10}, Range{1, 10}, true},
		{"adjacent no overlap", Range{1, 4}, Range{6, 10}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.r1.Overlaps(tc.r2)
			if result != tc.expected {
				t.Errorf("Range{%d,%d}.Overlaps(Range{%d,%d}) = %v, want %v",
					tc.r1.Start, tc.r1.End, tc.r2.Start, tc.r2.End, result, tc.expected)
			}
		})
	}
}

func TestRangeEqual(t *testing.T) {
	tests := []struct {
		name     string
		r1, r2   Range
		expected bool
	}{
		{"equal", Range{1, 10}, Range{1, 10}, true},
		{"different start", Range{1, 10}, Range{2, 10}, false},
		{"different end", Range{1, 10}, Range{1, 9}, false},
		{"completely different", Range{1, 5}, Range{6, 10}, false},
		{"zero ranges", Range{0, 0}, Range{0, 0}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.r1.Equal(tc.r2)
			if result != tc.expected {
				t.Errorf("Range{%d,%d}.Equal(Range{%d,%d}) = %v, want %v",
					tc.r1.Start, tc.r1.End, tc.r2.Start, tc.r2.End, result, tc.expected)
			}
		})
	}
}

func TestRangeIntersection(t *testing.T) {
	tests := []struct {
		name     string
		r1, r2   Range
		expected Range
	}{
		{"full overlap", Range{1, 10}, Range{3, 7}, Range{3, 7}},
		{"partial left", Range{1, 5}, Range{3, 10}, Range{3, 5}},
		{"partial right", Range{5, 10}, Range{1, 7}, Range{5, 7}},
		{"same range", Range{1, 10}, Range{1, 10}, Range{1, 10}},
		{"no overlap", Range{1, 4}, Range{6, 10}, Range{0, 0}},
		{"touching edge", Range{1, 5}, Range{5, 10}, Range{5, 5}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.r1.Intersection(tc.r2)
			if result != tc.expected {
				t.Errorf("Range{%d,%d}.Intersection(Range{%d,%d}) = %v, want %v",
					tc.r1.Start, tc.r1.End, tc.r2.Start, tc.r2.End, result, tc.expected)
			}
		})
	}
}

// =============================================================================
// Rangeset Tests
// =============================================================================

func TestRangesetAdd(t *testing.T) {
	var rs Rangeset[int64]

	rs.Add(10, 20)
	if len(rs) != 1 {
		t.Errorf("Add first: expected 1 range, got %d", len(rs))
	}

	rs.Add(30, 40)
	if len(rs) != 2 {
		t.Errorf("Add second: expected 2 ranges, got %d", len(rs))
	}

	// Add overlapping range
	rs.Add(15, 35)
	if len(rs) != 1 {
		t.Errorf("Add overlapping: expected 1 merged range, got %d", len(rs))
	}
	if rs[0].start != 10 || rs[0].end != 40 {
		t.Errorf("Add overlapping: expected [10,40), got [%d,%d)", rs[0].start, rs[0].end)
	}
}

func TestRangesetAddEmpty(t *testing.T) {
	var rs Rangeset[int64]
	rs.Add(10, 10) // Empty range
	if len(rs) != 0 {
		t.Errorf("Add empty range: expected 0 ranges, got %d", len(rs))
	}
}

func TestRangesetAddAdjacent(t *testing.T) {
	var rs Rangeset[int64]
	rs.Add(10, 20)
	rs.Add(20, 30) // Adjacent, should merge
	if len(rs) != 1 {
		t.Errorf("Add adjacent: expected 1 range, got %d", len(rs))
	}
	if rs[0].start != 10 || rs[0].end != 30 {
		t.Errorf("Add adjacent: expected [10,30), got [%d,%d)", rs[0].start, rs[0].end)
	}
}

func TestRangesetSub(t *testing.T) {
	var rs Rangeset[int64]
	rs.Add(10, 50)

	// Remove middle
	rs.Sub(20, 30)
	if len(rs) != 2 {
		t.Errorf("Sub middle: expected 2 ranges, got %d", len(rs))
	}

	// Should have [10,20) and [30,50)
	if rs[0].start != 10 || rs[0].end != 20 {
		t.Errorf("Sub middle: first range expected [10,20), got [%d,%d)", rs[0].start, rs[0].end)
	}
	if rs[1].start != 30 || rs[1].end != 50 {
		t.Errorf("Sub middle: second range expected [30,50), got [%d,%d)", rs[1].start, rs[1].end)
	}
}

func TestRangesetSubPrefix(t *testing.T) {
	var rs Rangeset[int64]
	rs.Add(10, 50)
	rs.Sub(10, 20)

	if len(rs) != 1 {
		t.Errorf("Sub prefix: expected 1 range, got %d", len(rs))
	}
	if rs[0].start != 20 || rs[0].end != 50 {
		t.Errorf("Sub prefix: expected [20,50), got [%d,%d)", rs[0].start, rs[0].end)
	}
}

func TestRangesetSubSuffix(t *testing.T) {
	var rs Rangeset[int64]
	rs.Add(10, 50)
	rs.Sub(40, 50)

	if len(rs) != 1 {
		t.Errorf("Sub suffix: expected 1 range, got %d", len(rs))
	}
	if rs[0].start != 10 || rs[0].end != 40 {
		t.Errorf("Sub suffix: expected [10,40), got [%d,%d)", rs[0].start, rs[0].end)
	}
}

func TestRangesetSubAll(t *testing.T) {
	var rs Rangeset[int64]
	rs.Add(10, 50)
	rs.Sub(5, 55)

	if len(rs) != 0 {
		t.Errorf("Sub all: expected 0 ranges, got %d", len(rs))
	}
}

func TestRangesetContains(t *testing.T) {
	var rs Rangeset[int64]
	rs.Add(10, 20)
	rs.Add(30, 40)

	tests := []struct {
		v        int64
		expected bool
	}{
		{5, false},
		{10, true},
		{15, true},
		{19, true},
		{20, false}, // Half-open [10,20)
		{25, false},
		{30, true},
		{35, true},
		{40, false},
		{50, false},
	}

	for _, tc := range tests {
		if rs.Contains(tc.v) != tc.expected {
			t.Errorf("Contains(%d) = %v, want %v", tc.v, !tc.expected, tc.expected)
		}
	}
}

func TestRangesetMinMax(t *testing.T) {
	var rs Rangeset[int64]

	// Empty rangeset
	if rs.Min() != 0 || rs.Max() != 0 {
		t.Error("Empty rangeset: Min/Max should be 0")
	}

	rs.Add(10, 20)
	rs.Add(30, 40)

	if rs.Min() != 10 {
		t.Errorf("Min: expected 10, got %d", rs.Min())
	}
	if rs.Max() != 39 { // Max is end-1
		t.Errorf("Max: expected 39, got %d", rs.Max())
	}
	if rs.End() != 40 {
		t.Errorf("End: expected 40, got %d", rs.End())
	}
}

func TestRangesetRangeContaining(t *testing.T) {
	var rs Rangeset[int64]
	rs.Add(10, 20)
	rs.Add(30, 40)

	r := rs.RangeContaining(15)
	if r.start != 10 || r.end != 20 {
		t.Errorf("RangeContaining(15): expected [10,20), got [%d,%d)", r.start, r.end)
	}

	r = rs.RangeContaining(25)
	if r.start != 0 || r.end != 0 {
		t.Errorf("RangeContaining(25): expected [0,0), got [%d,%d)", r.start, r.end)
	}
}

func TestRangesetIsRange(t *testing.T) {
	var rs Rangeset[int64]

	if !rs.IsRange(0, 0) {
		t.Error("Empty rangeset should be IsRange(0,0)")
	}

	rs.Add(10, 20)
	if !rs.IsRange(10, 20) {
		t.Error("Single range should match IsRange")
	}
	if rs.IsRange(10, 25) {
		t.Error("Different range should not match IsRange")
	}

	rs.Add(30, 40)
	if rs.IsRange(10, 40) {
		t.Error("Multiple ranges should not match IsRange")
	}
}

func TestI64rangeSize(t *testing.T) {
	r := I64range[int64]{10, 20}
	if r.Size() != 10 {
		t.Errorf("Size: expected 10, got %d", r.Size())
	}

	empty := I64range[int64]{10, 10}
	if empty.Size() != 0 {
		t.Errorf("Empty Size: expected 0, got %d", empty.Size())
	}
}

func TestI64rangeContains(t *testing.T) {
	r := I64range[int64]{10, 20}

	if !r.Contains(10) {
		t.Error("Should contain start")
	}
	if !r.Contains(15) {
		t.Error("Should contain middle")
	}
	if r.Contains(20) {
		t.Error("Should not contain end (half-open)")
	}
	if r.Contains(5) {
		t.Error("Should not contain before start")
	}
}
