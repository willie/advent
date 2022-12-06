package main

import "github.com/willie/advent/aoc"

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
func (r Range) Intersection(r2 Range) Range {
	if !r.Overlaps(r2) {
		return Range{0, 0}
	}

	return Range{aoc.Max(r.start, r2.start), aoc.Min(r.end, r2.end)}
}
