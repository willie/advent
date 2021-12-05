package main

import "github.com/willie/advent/aoc"

func part1(in []string) (result int) {
	return
}

func part2(in []string) (result int) {
	return
}

const day = "https://adventofcode.com/2021/day/5"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 4512)
	aoc.Test("test2", part2(aoc.Strings("test")), 1924)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
