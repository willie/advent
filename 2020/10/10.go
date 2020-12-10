package main

import (
	"sort"

	"github.com/willie/advent/aoc"
)

func combined(in aoc.Ints) (first, second int) {
	sort.Ints(in)

	last := 0
	differences := map[int]int{3: 1}

	perms := map[int]int{0: 1}

	for _, current := range in {
		differences[current-last]++

		// consecutive runs cause permutations
		perms[current] = perms[current-1] + perms[current-2] + perms[current-3]

		last = current
	}

	first = differences[1] * differences[3]
	second = perms[aoc.Max(in...)]

	return
}

const day = "https://adventofcode.com/2020/day/10"

func main() {
	println(day)
	aoc.Input(day)

	println("------- combined")

	t1, t2 := combined(aoc.LoadInts("test"))
	aoc.TestX("test", t1, t2, 7*5, 8)

	t1, t2 = combined(aoc.LoadInts("test2"))
	aoc.TestX("test2", t1, t2, 22*10, 19208)

	r1, r2 := combined(aoc.LoadInts(day))
	aoc.RunX("combined", r1, r2)
}
