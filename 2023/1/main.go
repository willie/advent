package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/willie/advent/aoc"
)

func IndexVal(numbers map[string]int, s string) (v int) {
	pos := math.MaxInt
	for num, val := range numbers {
		if i := strings.Index(s, num); (i != -1) && i < pos {
			pos = i
			v = val
		}
	}
	return
}

func LastIndexVal(numbers map[string]int, s string) (v int) {
	pos := math.MinInt
	for num, val := range numbers {
		if i := strings.LastIndex(s, num); (i != -1) && i > pos {
			pos = i
			v = val
		}
	}
	return
}

func part1(in []string) (total int) {
	for _, s := range in {
		f := strings.IndexAny(s, "0123456789")
		l := strings.LastIndexAny(s, "0123456789")

		n := fmt.Sprintf("%c%c", s[f], s[l])
		total += aoc.AtoI(n)
	}

	return
}

func part2(in []string) (total int) {
	for _, s := range in {
		f := IndexVal(numbers(), s)
		l := LastIndexVal(numbers(), s)

		n := fmt.Sprintf("%d%d", f, l)
		total += aoc.AtoI(n)
	}
	return
}

const day = "https://adventofcode.com/2023/day/1"

// create a map of the first 10 numbers by name to their value
func numbers() map[string]int {
	return map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
}

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 142)
	aoc.Test("test2", part2(aoc.Strings("test2")), 281)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
