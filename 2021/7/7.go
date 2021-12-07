package main

import (
	"strings"

	"github.com/willie/advent/aoc"
)

func splitInts(in string) (ints []int) {
	for _, s := range strings.Split(strings.TrimSpace(in), ",") {
		ints = append(ints, aoc.AtoI(s))
	}
	return
}

func part1(in string) (result int) {
	crabs := splitInts(in)
	max := aoc.Max(crabs...)
	moves := make([]int, max)

	for p := 0; p < max; p++ {
		count := 0

		for _, c := range crabs {
			count += aoc.Abs(c - p)
		}

		moves[p] = count
	}

	return aoc.Min(moves...)
}

func fuel(distance int) (cost int) {
	for distance != 0 {
		cost += distance
		distance--
	}
	return
}

func part2(in string) (result int) {
	crabs := splitInts(in)
	max := aoc.Max(crabs...)
	moves := make([]int, max)

	for p := 0; p < max; p++ {
		count := 0

		for _, c := range crabs {
			count += fuel(aoc.Abs(c - p))
		}

		moves[p] = count
	}

	return aoc.Min(moves...)
}

const day = "https://adventofcode.com/2021/day/7"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.String("test")), 37)
	aoc.Test("test2", part2(aoc.String("test")), 168)

	println("-------")

	aoc.Run("part1", part1(aoc.String(day)))
	aoc.Run("part2", part2(aoc.String(day)))
}
