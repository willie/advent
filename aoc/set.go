package aoc

import (
	"sort"
)

var exists = struct{}{}

// StringSet is a set of strings
type StringSet map[string]struct{}

// NewStringSet returns a new StringSet
func NewStringSet(values ...string) StringSet {
	s := StringSet{}
	s.AddMany(values)
	return s
}

// Add values to the set
func (s StringSet) Add(values ...string) StringSet { return s.AddMany(values) }

// AddMany values to the set
func (s StringSet) AddMany(values []string) StringSet {
	for _, value := range values {
		s[value] = exists
	}
	return s
}

// AddSet to the set
func (s StringSet) AddSet(set StringSet) StringSet {
	for key := range set {
		s[key] = exists
	}
	return s
}

// Remove values from set
func (s StringSet) Remove(values ...string) StringSet {
	for _, value := range values {
		delete(s, value)
	}
	return s
}

// Contains returns if a value is in the set
func (s StringSet) Contains(value string) bool {
	_, c := s[value]
	return c
}

// Values returns the values in set
func (s StringSet) Values() (values []string) {
	for k := range s {
		values = append(values, k)
	}
	return
}

// Subtract returns the differences
func (s StringSet) Subtract(x StringSet) (difference StringSet) {
	difference = StringSet{}

	for k := range s {
		if !x.Contains(k) {
			difference.Add(k)
		}
	}
	return
}

// IntSet is a set of ints
type IntSet map[int]struct{}

// NewIntSet returns a new IntSet
func NewIntSet(values ...int) IntSet {
	s := IntSet{}
	s.AddMany(values)
	return s
}

// Add a value to the set
func (s IntSet) Add(values ...int) IntSet { return s.AddMany(values) }

// AddMany values to the set
func (s IntSet) AddMany(values []int) IntSet {
	for _, value := range values {
		s[value] = exists
	}
	return s
}

// Remove value from set
func (s IntSet) Remove(values ...int) IntSet {
	for _, value := range values {
		delete(s, value)
	}
	return s
}

// Contains returns if a value is in the set
func (s IntSet) Contains(value int) bool {
	_, c := s[value]
	return c
}

// Values returns the values in set
func (s IntSet) Values() (values []int) {
	for k := range s {
		values = append(values, k)
	}
	sort.Ints(values)
	return
}

// Subtract returns the differences
func (s IntSet) Subtract(x IntSet) (difference IntSet) {
	difference = IntSet{}

	for k := range s {
		if !x.Contains(k) {
			difference.Add(k)
		}
	}
	return
}
