package aoc

import (
	"log"
	"slices"
)

// Ints is []int with convenience methods.
// Note: For new code, consider using:
//   - slices.Min, slices.Max from stdlib
//   - slices.Index, slices.Contains from stdlib
//   - Sum[T], Product[T] from math.go
type Ints []int

// StringInts returns ints from strings
func StringInts(ss []string) (ints Ints) {
	for _, s := range ss {
		ints = append(ints, AtoI(s))
	}
	return
}

// Sum returns total
func (ints Ints) Sum() (sum int) {
	for _, i := range ints {
		sum += i
	}

	return
}

// Product multiplies all the numbers together
func (ints Ints) Product() (p int) {
	p = 1
	for _, i := range ints {
		p = p * i
	}
	return
}

// Min returns smallest value.
// Note: slices.Min from stdlib is preferred for new code.
func (ints Ints) Min() int {
	if len(ints) == 0 {
		log.Fatalln("no values in array")
	}
	return slices.Min(ints)
}

// Max returns largest value.
// Note: slices.Max from stdlib is preferred for new code.
func (ints Ints) Max() int {
	if len(ints) == 0 {
		log.Fatalln("no values in array")
	}
	return slices.Max(ints)
}

// Last returns last value
func (ints Ints) Last() (last int) {
	if len(ints) == 0 {
		log.Fatalln("no values in array")
	}

	return ints[len(ints)-1]
}

// Index returns the first index where value is found, or -1 if not found.
// Note: slices.Index from stdlib is preferred for new code.
func (ints Ints) Index(in int) int {
	return slices.Index(ints, in)
}

// LastIndex where i is
func (ints Ints) LastIndex(in int) (last int) {
	last = -1

	for i, c := range ints {
		if c == in {
			last = i
		}
	}

	return
}

// AllIndex all indexes
func (ints Ints) AllIndex(in int) (idx []int) {
	for i, c := range ints {
		if c == in {
			idx = append(idx, i)
		}
	}

	return
}

// Series returns array of low including high
func Series(low, high int) (series Ints) {
	series = make(Ints, (high-low)+1)

	x := 0
	for i := low; i <= high; i++ {
		series[x] = i
		x++
	}
	return
}
