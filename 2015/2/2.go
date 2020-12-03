package main

import (
	"fmt"

	"github.com/willie/advent/aoc"
)

func parse(in string) (l, w, h int) {
	fmt.Sscanf(in, "%dx%dx%d", &l, &w, &h)
	return
}

func wrap(in string) (paper int) {
	l, w, h := parse(in)
	sides := []int{l * w, w * h, h * l}

	return 2*aoc.Sum(sides...) + aoc.Min(sides...)
}

func ribbon(in string) (ribbon int) {
	l, w, h := parse(in)
	perimeter := []int{2 * (l + w), 2 * (w + h), 2 * (h + l)}

	return aoc.Min(perimeter...) + (l * w * h)
}

func part1(in []string) (total int) {
	for _, i := range in {
		total += wrap(i)
	}

	return
}

func part2(in []string) (total int) {
	for _, i := range in {
		total += ribbon(i)
	}

	return
}

const day = "https://adventofcode.com/2015/day/2"

func main() {
	println(day)

	test1 := []struct {
		s        string
		expected int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, t := range test1 {
		aoc.Test("test1", wrap(t.s), t.expected)
	}

	aoc.Run("part1", part1(aoc.Strings(day)))

	println("-------")

	test2 := []struct {
		s        string
		expected int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	for _, t := range test2 {
		aoc.Test("test2", ribbon(t.s), t.expected)
	}

	aoc.Run("part2", part2(aoc.Strings(day)))
}
