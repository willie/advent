package main

import (
	"sort"

	"github.com/willie/advent/aoc"
)

func part1(in []string) int {
	elves := make(map[int]int)
	i := 1

	for _, s := range in {
		if s == "" {
			i++
			continue
		}
		elves[i] = elves[i] + aoc.AtoI(s)
	}

	most := 0
	for _, c := range elves {
		if c > most {
			most = c
		}
	}

	return most
}

func part2(in []string) int {
	elves := make(map[int]int)
	i := 1

	for _, s := range in {
		if s == "" {
			i++
			continue
		}
		elves[i] = elves[i] + aoc.AtoI(s)
	}

	calories := []int{}
	for _, c := range elves {
		calories = append(calories, c)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	return calories[0] + calories[1] + calories[2]
}

const day = "https://adventofcode.com/2022/day/1"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 24000)
	aoc.Test("test2", part2(aoc.Strings("test")), 45000)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
