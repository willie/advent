package main

type Range struct {
	start int
	end   int
}

// Contains returns true if r1 FULLY contains r2
func (r1 Range) Contains(r2 Range) bool {
	return r1.start <= r2.start && r1.end >= r2.end
}

// Overlaps returns true if r1 overlaps ANY with r2
func (r1 Range) Overlaps(r2 Range) bool {
	return r1.start <= r2.end && r2.start <= r1.end
}

// Equal returns true if r1 == r2
func (r1 Range) Equal(r2 Range) bool {
	return r1.start == r2.start && r1.start == r2.end
}

// a little extra

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (r Range) Intersection(r2 Range) Range {
	if !r.Overlaps(r2) {
		return Range{0, 0}
	}

	return Range{max(r.start, r2.start), min(r.end, r2.end)}
}
