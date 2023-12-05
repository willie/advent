package aoc

type Range struct {
	Start int
	End   int
}

// Contains returns true if r1 FULLY contains r2
func (r1 Range) Contains(r2 Range) bool {
	return r1.Start <= r2.Start && r1.End >= r2.End
}

// Overlaps returns true if r1 overlaps ANY with r2
func (r1 Range) Overlaps(r2 Range) bool {
	return r1.Start <= r2.End && r2.Start <= r1.End
}

// Equal returns true if r1 == r2
func (r1 Range) Equal(r2 Range) bool {
	return r1.Start == r2.Start && r1.Start == r2.End
}

// a little extra
func (r Range) Intersection(r2 Range) Range {
	if !r.Overlaps(r2) {
		return Range{0, 0}
	}

	return Range{Max(r.Start, r2.Start), Min(r.End, r2.End)}
}
