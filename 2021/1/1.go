package main

import (
	"github.com/willie/advent/aoc"
)

func part1(ints []int) (increase int) {
	prev := ints[0]

	for _, i := range ints {
		if i > prev {
			increase++
		}

		prev = i
	}
	return
}

func part2(ints []int) (increase int) {
	prev := ints[0] + ints[1] + ints[2]

	for start := 0; start+2 < len(ints); start++ {
		sum := ints[start] + ints[start+1] + ints[start+2]

		if sum > prev {
			increase++
		}

		prev = sum
	}
	return
}

const day = "https://adventofcode.com/2021/day/1"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.LoadInts("test")), 7)
	aoc.Test("test2", part2(aoc.LoadInts("test")), 5)

	println("-------")

	aoc.Run("part1", part1(aoc.LoadInts(day)))
	aoc.Run("part2", part2(aoc.LoadInts(day)))
}
