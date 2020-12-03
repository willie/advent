package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (valid int) {
	for _, i := range in {
		var low, high int
		var letter, password string

		// Go Sscanf doesn't handle colons for string
		i = strings.ReplaceAll(i, ":", "")
		fmt.Sscanf(i, "%d-%d %s %s", &low, &high, &letter, &password)

		count := strings.Count(password, letter)
		if low <= count && count <= high {
			valid++
		}
	}
	return
}

func part2(in []string) (valid int) {
	for _, i := range in {
		var low, high int
		var letter, password string

		// Go Sscanf doesn't handle colons for string
		i = strings.ReplaceAll(i, ":", "")
		fmt.Sscanf(i, "%d-%d %s %s", &low, &high, &letter, &password)

		l := password[low-1:low] == letter
		h := password[high-1:high] == letter

		switch {
		case l && h:
		case l, h:
			valid++
		}
	}
	return
}

const day = "https://adventofcode.com/2020/day/2"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 2)
	aoc.Run("part1", part1(aoc.Strings(day)))

	println("-------")

	aoc.Test("test2", part2(aoc.Strings("test")), 1)
	aoc.Run("part2", part2(aoc.Strings(day)))
}
