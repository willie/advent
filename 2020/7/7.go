package main

import (
	"strings"

	"github.com/willie/advent/aoc"
)

func contains(m map[string][]string, color string) (out []string) {
	if colors, has := m[color]; has {
		out = append(out, colors...)
		for _, c := range colors {
			out = append(out, contains(m, c)...)
		}
	}

	return
}

func part1(in []string) (count int) {
	rules := map[string][]string{}
	contained := map[string][]string{}

	for _, i := range in {
		i := strings.NewReplacer(" bags", "", " bag", "", ".", "").Replace(i) // convert to binary
		parts := strings.Split(i, " contain ")

		color := parts[0]
		inner := strings.Split(parts[1], ",")

		for _, j := range inner {
			if j == "no other" {
				continue
			}

			f := strings.Fields(j)
			c := f[1] + " " + f[2]

			// count := aoc.AtoI(f[0])

			rules[color] = append(rules[color], c)
			contained[c] = append(contained[c], color)
			// fmt.Println(count, color)
		}
	}

	con := aoc.NewStringSet(contains(contained, "shiny gold")...)

	// fmt.Println(con)

	return len(con)
}

func bagger(rules map[string]map[string]int, color string) (bags int) {
	for c, count := range rules[color] {
		bags += count * (1 + bagger(rules, c))
	}

	return
}

func part2(in []string) (count int) {
	rules := map[string]map[string]int{}
	contained := map[string][]string{}

	for _, i := range in {
		i := strings.NewReplacer(" bags", "", " bag", "", ".", "").Replace(i) // convert to binary
		parts := strings.Split(i, " contain ")

		color := parts[0]
		inner := strings.Split(parts[1], ",")

		rules[color] = map[string]int{}

		for _, j := range inner {
			if j == "no other" {
				continue
			}

			f := strings.Fields(j)
			c := f[1] + " " + f[2]

			count := aoc.AtoI(f[0])

			rules[color][c] = count
			contained[c] = append(contained[c], color)
			// fmt.Println(count, color)
		}
	}

	return bagger(rules, "shiny gold")
}

const day = "https://adventofcode.com/2020/day/7"

func main() {
	println(day)

	println("------- part 1")

	aoc.Test("test1", part1(aoc.Strings("test")), 4)
	aoc.Run("part1", part1(aoc.Strings(day)))

	println("------- part 2")

	aoc.Test("test2", part2(aoc.Strings("test")), 32)
	aoc.Test("test2-2", part2(aoc.Strings("test2")), 126)
	aoc.Run("part2", part2(aoc.Strings(day)))
}
