package aoc

import (
	"sort"

	"golang.org/x/exp/maps"
)

type Set[T comparable] map[T]struct{}

// NewSet[T] returns a new Set[T]
func NewSet[T comparable](values ...T) Set[T] {
	s := Set[T]{}
	s.AddSlice(values)
	return s
}

// Add values to the set
func (s Set[T]) Add(values ...T) Set[T] { return s.AddSlice(values) }

// AddSlice to the set
func (s Set[T]) AddSlice(values []T) Set[T] {
	for _, value := range values {
		s[value] = struct{}{}
	}
	return s
}

// AddSet to the set
func (s Set[T]) AddSet(set Set[T]) Set[T] {
	for key := range set {
		s[key] = struct{}{}
	}
	return s
}

// Remove values from set
func (s Set[T]) Remove(values ...T) Set[T] {
	for _, value := range values {
		delete(s, value)
	}
	return s
}

// Contains returns if a value is in the set
func (s Set[T]) Contains(value T) bool {
	_, c := s[value]
	return c
}

// ContainsAll returns if all values are in the set
func (s Set[T]) ContainsAll(values []T) bool {
	for _, v := range values {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

// ContainsAny returns if any of the values are in the set
func (s Set[T]) ContainsAny(values []T) bool {
	for _, v := range values {
		if s.Contains(v) {
			return true
		}
	}
	return false
}

// Subtract returns the differences
func (s Set[T]) Subtract(x Set[T]) (difference Set[T]) {
	difference = Set[T]{}
	for k := range s {
		if !x.Contains(k) {
			difference.Add(k)
		}
	}
	return
}

// Values returns the values in set
func (s Set[T]) Values() (values []T) { return maps.Keys(s) }

// Intersect returns the differences
func (s Set[T]) Intersect(x Set[T]) (intersection Set[T]) {
	intersection = Set[T]{}
	for k := range s {
		if x.Contains(k) {
			intersection.Add(k)
		}
	}
	return
}

// Union returns the combination of two sets
func (s Set[T]) Union(x Set[T]) (union Set[T]) {
	union = Set[T]{}
	union.AddSet(s)
	union.AddSet(x)
	return
}

// StringSet is a set of strings, here for historical reasons
type StringSet Set[string]

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
		s[value] = struct{}{}
	}
	return s
}

// AddSet to the set
func (s StringSet) AddSet(set StringSet) StringSet {
	for key := range set {
		s[key] = struct{}{}
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

// ContainsAll returns if all values are in the set
func (s StringSet) ContainsAll(values []string) bool {
	for _, v := range values {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

// ContainsAny returns if any of the values are in the set
func (s StringSet) ContainsAny(values []string) bool {
	for _, v := range values {
		if s.Contains(v) {
			return true
		}
	}
	return false
}

// Values returns the values in set
func (s StringSet) Values() (values []string) {
	for k := range s {
		values = append(values, k)
	}
	sort.Strings(values)
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

// Intersection returns the union
func (s StringSet) Intersection(x StringSet) (intersection StringSet) {
	intersection = StringSet{}
	for k := range s {
		if x.Contains(k) {
			intersection.Add(k)
		}
	}
	return
}

// IntSet is a set of ints, here for historical reasons.
type IntSet Set[int]

// type IntSet map[int]struct{}

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
		s[value] = struct{}{}
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
