// copied from the go internal standard library

package aoc

// A Rangeset is a set of int64s, stored as an ordered list of non-overlapping,
// non-empty ranges.
//
// Rangesets are efficient for small numbers of ranges,
// which is expected to be the common case.
type Rangeset[T ~int64] []I64range[T]

type I64range[T ~int64] struct {
	start, end T // [start, end)
}

// Size returns the Size of the range.
func (r I64range[T]) Size() T {
	return r.end - r.start
}

// Contains reports whether v is in the range.
func (r I64range[T]) Contains(v T) bool {
	return r.start <= v && v < r.end
}

// Add adds [start, end) to the set, combining it with existing ranges if necessary.
func (s *Rangeset[T]) Add(start, end T) {
	if start == end {
		return
	}
	for i := range *s {
		r := &(*s)[i]
		if r.start > end {
			// The new range comes before range i.
			s.InsertRange(i, start, end)
			return
		}
		if start > r.end {
			// The new range comes after range i.
			continue
		}
		// The new range is adjacent to or overlapping range i.
		if start < r.start {
			r.start = start
		}
		if end <= r.end {
			return
		}
		// Possibly coalesce subsquent ranges into range i.
		r.end = end
		j := i + 1
		for ; j < len(*s) && r.end >= (*s)[j].start; j++ {
			if e := (*s)[j].end; e > r.end {
				// Range j ends after the new range.
				r.end = e
			}
		}
		s.RemoveRanges(i+1, j)
		return
	}
	*s = append(*s, I64range[T]{start, end})
}

// Sub removes [start, end) from the set.
func (s *Rangeset[T]) Sub(start, end T) {
	removefrom, removeto := -1, -1
	for i := range *s {
		r := &(*s)[i]
		if end < r.start {
			break
		}
		if r.end < start {
			continue
		}
		switch {
		case start <= r.start && end >= r.end:
			// Remove the entire range.
			if removefrom == -1 {
				removefrom = i
			}
			removeto = i + 1
		case start <= r.start:
			// Remove a prefix.
			r.start = end
		case end >= r.end:
			// Remove a suffix.
			r.end = start
		default:
			// Remove the middle, leaving two new ranges.
			rend := r.end
			r.end = start
			s.InsertRange(i+1, end, rend)
			return
		}
	}
	if removefrom != -1 {
		s.RemoveRanges(removefrom, removeto)
	}
}

// Contains reports whether s Contains v.
func (s Rangeset[T]) Contains(v T) bool {
	for _, r := range s {
		if v >= r.end {
			continue
		}
		if r.start <= v {
			return true
		}
		return false
	}
	return false
}

// RangeContaining returns the range containing v, or the range [0,0) if v is not in s.
func (s Rangeset[T]) RangeContaining(v T) I64range[T] {
	for _, r := range s {
		if v >= r.end {
			continue
		}
		if r.start <= v {
			return r
		}
		break
	}
	return I64range[T]{0, 0}
}

// min returns the minimum value in the set, or 0 if empty.
func (s Rangeset[T]) Min() T {
	if len(s) == 0 {
		return 0
	}
	return s[0].start
}

// max returns the maximum value in the set, or 0 if empty.
func (s Rangeset[T]) Max() T {
	if len(s) == 0 {
		return 0
	}
	return s[len(s)-1].end - 1
}

// end returns the end of the last range in the set, or 0 if empty.
func (s Rangeset[T]) End() T {
	if len(s) == 0 {
		return 0
	}
	return s[len(s)-1].end
}

// NumRanges returns the number of ranges in the rangeset.
func (s Rangeset[T]) NumRanges() int {
	return len(s)
}

// IsRange reports if the rangeset covers exactly the range [start, end).
func (s Rangeset[T]) IsRange(start, end T) bool {
	switch len(s) {
	case 0:
		return start == 0 && end == 0
	case 1:
		return s[0].start == start && s[0].end == end
	}
	return false
}

// RemoveRanges removes ranges [i,j).
func (s *Rangeset[T]) RemoveRanges(i, j int) {
	if i == j {
		return
	}
	copy((*s)[i:], (*s)[j:])
	*s = (*s)[:len(*s)-(j-i)]
}

// insert adds a new range at index i.
func (s *Rangeset[T]) InsertRange(i int, start, end T) {
	*s = append(*s, I64range[T]{})
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = I64range[T]{start, end}
}
