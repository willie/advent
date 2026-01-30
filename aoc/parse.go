package aoc

import (
	"regexp"
	"strconv"
)

var intRegex = regexp.MustCompile(`-?\d+`)

// ParseInts extracts all integers from a string.
// Handles negative numbers. Very common AoC pattern.
func ParseInts(s string) []int {
	matches := intRegex.FindAllString(s, -1)
	result := make([]int, len(matches))
	for i, m := range matches {
		result[i] = AtoI(m)
	}
	return result
}

// ParseInt64s extracts all integers from a string as int64.
func ParseInt64s(s string) []int64 {
	matches := intRegex.FindAllString(s, -1)
	result := make([]int64, len(matches))
	for i, m := range matches {
		n, _ := strconv.ParseInt(m, 10, 64)
		result[i] = n
	}
	return result
}

// MustInt parses a string to int, panics on error.
// Alias for AtoI for clarity.
func MustInt(s string) int {
	return AtoI(s)
}

// MustInt64 parses a string to int64, panics on error.
func MustInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

// SplitInts splits a string by a delimiter and converts each part to int.
func SplitInts(s string, sep string) []int {
	parts := regexp.MustCompile(regexp.QuoteMeta(sep)).Split(s, -1)
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		if p != "" {
			result = append(result, AtoI(p))
		}
	}
	return result
}

// Words splits a string on whitespace, returning non-empty parts.
func Words(s string) []string {
	return regexp.MustCompile(`\s+`).Split(s, -1)
}
