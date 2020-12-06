package main

import (
	"github.com/willie/advent/aoc"
)

func part1(ints []int, sum int) (product int) {
	for _, i := range ints {
		for _, j := range ints {
			if j == i {
				continue
			}

			if i+j == sum {
				return i * j
			}
		}
	}
	return
}

func part2(ints []int, sum int) (product int) {
	for _, i := range ints {
		for _, j := range ints {
			if j == i {
				continue
			}

			for _, k := range ints {
				if k == j {
					continue
				}

				if i+j+k == sum {
					return i * j * k
				}
			}
		}
	}
	return
}

const day = "https://adventofcode.com/2020/day/1"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.LoadInts("test"), 2020), 514579)
	aoc.Test("test2", part2(aoc.LoadInts("test"), 2020), 241861950)

	println("-------")

	aoc.Run("part1", part1(aoc.LoadInts(day), 2020))
	aoc.Run("part2", part2(aoc.LoadInts(day), 2020))
}
