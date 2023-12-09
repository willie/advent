package main

import (
	"slices"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (total int) {
	for _, line := range in {
		total += oasis(line)
	}

	return
}

func part2(in []string) (total int) {
	for _, line := range in {
		total += oasis2(line)
	}
	return
}

func generateSequence(in []int) (out []int) {
	for i := 1; i < len(in); i++ {
		out = append(out, in[i]-in[i-1])
	}
	return
}

func oasis(in string) (total int) {
	history := [][]int{}
	history = append(history, aoc.StringInts(strings.Fields(in)))

	for next := history[0]; aoc.Sum(next...) != 0; {
		next = generateSequence(next)
		history = append(history, next)

		// fmt.Println(next)
	}

	// get all the last values of history in reverse order
	lasts := []int{}
	for _, h := range history {
		lasts = append(lasts, h[len(h)-1])
	}
	slices.Reverse(lasts)

	for _, l := range lasts {
		total += l
	}

	return
}

func oasis2(in string) (total int) {
	history := [][]int{}
	history = append(history, aoc.StringInts(strings.Fields(in)))

	for next := history[0]; aoc.Sum(next...) != 0; {
		next = generateSequence(next)
		history = append(history, next)
	}

	firsts := []int{}
	for _, h := range history {
		firsts = append(firsts, h[0])
	}
	slices.Reverse(firsts)

	prev := 0
	for _, l := range firsts {
		prev = l - prev
	}

	return prev
}

const day = "https://adventofcode.com/2023/day/9"

func main() {
	println(day)

	aoc.Test("test1", oasis(test1), 18)
	aoc.Test("test2", oasis(test2), 28)
	aoc.Test("test3", oasis(test3), 68)
	aoc.Test("test4", part1(aoc.Strings("test")), 114)

	aoc.Test("test1", oasis2(test1), -3)
	aoc.Test("test2", oasis2(test2), 0)
	aoc.Test("test3", oasis2(test3), 5)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}

const (
	test1 = "0 3 6 9 12 15"
	test2 = "1 3 6 10 15 21"
	test3 = "10 13 16 21 30 45"
)
