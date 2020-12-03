package main

import (
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in string) (floor int) {
	return strings.Count(in, "(") - strings.Count(in, ")")
}

func part2(in string) (pos int) {
	floor := 0
	for i, c := range in {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor < 0 {
			return i + 1
		}
	}
	return
}

const day = "https://adventofcode.com/2015/day/1"

func main() {
	println(day)

	test1 := []struct {
		s        string
		expected int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, t := range test1 {
		aoc.Test("test1", part1(t.s), t.expected)
	}

	aoc.Run("part1", part1(aoc.String(day)))

	println("-------")

	test2 := []struct {
		s        string
		expected int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, t := range test2 {
		aoc.Test("test2", part2(t.s), t.expected)
	}

	aoc.Run("part2", part2(aoc.String(day)))
}
