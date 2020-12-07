package main

import (
	"strings"

	"github.com/willie/advent/aoc"
)

func contains(m map[string][]string, color string) (out aoc.StringSet) {
	out = aoc.NewStringSet()

	if colors, has := m[color]; has {
		out.AddMany(colors)
		for _, c := range colors {
			out.AddSet(contains(m, c))
		}
	}

	return
}

func bagger(rules map[string]map[string]int, color string) (bags int) {
	for c, count := range rules[color] {
		bags += count * (1 + bagger(rules, c))
	}
	return
}

func combined(in []string) (count, bags int) {
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
		}
	}

	return len(contains(contained, "shiny gold")), bagger(rules, "shiny gold")
}

func part1(in []string) (count int) {
	count, _ = combined(in)
	return
}

func part2(in []string) (count int) {
	_, count = combined(in)
	return
}

const day = "https://adventofcode.com/2020/day/7"

func main() {
	println(day)

	println("------- part 1")

	aoc.Test("test", part1(aoc.Strings("test")), 4)
	aoc.Run("run", part1(aoc.Strings(day)))

	println("------- part 2")

	aoc.Test("test", part2(aoc.Strings("test")), 32)
	aoc.Test("test2", part2(aoc.Strings("test2")), 126)
	aoc.Run("run", part2(aoc.Strings(day)))

	println("------- combined")

	t1, t2 := combined(aoc.Strings("test"))
	aoc.TestX("test", t1, t2, 4, 32)

	r1, r2 := combined(aoc.Strings(day))
	aoc.RunX("part", r1, r2)
}
