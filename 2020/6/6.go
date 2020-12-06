package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in string) (count int) {
	groups := []aoc.StringSet{}
	in = strings.TrimSpace(in)
	in = strings.ReplaceAll(in, "\n", "|")
	in = strings.ReplaceAll(in, "||", "\n")
	in = strings.ReplaceAll(in, "|", "")
	for _, p := range strings.Split(in, "\n") {
		g := aoc.NewStringSet()

		for _, c := range p {
			letter := string(c)
			g.Add(letter)
			// fmt.Print(letter)
		}

		groups = append(groups, g)
		// fmt.Println()
	}

	for _, g := range groups {
		count += len(g)
	}
	return
}

func part2(in string) (count int) {
	groups := []aoc.StringSet{}
	in = strings.TrimSpace(in)
	for _, p := range strings.Split(in, "\n\n") {
		fmt.Println(".", p, ".")

		g := aoc.NewStringSet()

		for _, c := range p {
			letter := string(c)
			g.Add(letter)
			// fmt.Print(letter)
		}

		groups = append(groups, g)
		fmt.Println()
	}

	for _, g := range groups {
		count += len(g)
	}
	return
}

const day = "https://adventofcode.com/2020/day/6"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.String("test")), 11)
	aoc.Run("part1", part1(aoc.String(day)))

	println("-------")

	aoc.Test("test2", part2(aoc.String("test")), 6)
	// aoc.Test("test2", part2(aoc.String("valid")), 4)
	// aoc.Run("part2", part2(aoc.String(day)))
}
