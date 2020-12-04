package aoc

import "sort"

var exists = struct{}{}

type stringMap map[string]struct{}

// StringSet is a set of strings
type StringSet struct {
	stringMap
}

// NewStringSet returns a new StringSet
func NewStringSet(values ...string) *StringSet {
	s := &StringSet{}
	s.stringMap = make(map[string]struct{})
	s.AddMany(values)
	return s
}

// Add a value to the set
func (s *StringSet) Add(value string) *StringSet {
	s.stringMap[value] = exists
	return s
}

// AddMany values to the set
func (s *StringSet) AddMany(values []string) *StringSet {
	for _, value := range values {
		s.stringMap[value] = exists
	}
	return s
}

// Remove value from set
func (s *StringSet) Remove(value string) *StringSet {
	delete(s.stringMap, value)
	return s
}

// Contains returns if a value is in the set
func (s *StringSet) Contains(value string) bool {
	_, c := s.stringMap[value]
	return c
}

// Values returns the values in set
func (s *StringSet) Values() (values []string) {
	for k := range s.stringMap {
		values = append(values, k)
	}
	return
}

// Subtract returns the differences
func (s *StringSet) Subtract(x *StringSet) (difference []string) {
	for k := range s.stringMap {
		if !x.Contains(k) {
			difference = append(difference, k)
		}
	}
	return
}

type intMap map[int]struct{}

// IntSet is a set of ints
type IntSet struct {
	intMap
}

// NewIntSet returns a new IntSet
func NewIntSet() *IntSet {
	s := &IntSet{}
	s.intMap = make(map[int]struct{})
	return s
}

// Add a value to the set
func (s *IntSet) Add(value int) {
	s.intMap[value] = exists
}

// AddMany values to the set
func (s *IntSet) AddMany(values []int) *IntSet {
	for _, value := range values {
		s.intMap[value] = exists
	}
	return s
}

// Remove value from set
func (s *IntSet) Remove(value int) {
	delete(s.intMap, value)
}

// Contains returns if a value is in the set
func (s *IntSet) Contains(value int) bool {
	_, c := s.intMap[value]
	return c
}

// Values returns the values in set
func (s *IntSet) Values() (values []int) {
	for k := range s.intMap {
		values = append(values, k)
	}
	sort.Ints(values)
	return
}
