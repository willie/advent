package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/willie/advent/aoc"
)

type mapping struct {
	source int
	dest   int
	length int
}

func part1(in []string) (total int) {
	seeds := aoc.StringInts(strings.Fields(strings.Split(in[0], ": ")[1]))

	name := "seeds"

	mappings := []mapping{}
	for _, s := range in[2:] {
		if m := strings.Split(s, " map:"); len(m) > 1 { // reset
			fmt.Println(name, seeds)

			name = strings.Split(m[0], "-")[2]
			// mapping = map[int]int{}
			continue
		}

		if m := aoc.StringInts(strings.Fields(s)); len(m) == 3 {
			dest, source, length := m[0], m[1], m[2]

			mappings = append(mappings, mapping{source, dest, length})
			fmt.Println(dest, source, length)
			continue
		}

		if s == "" {
			slices.Reverse(mappings)

			for i, seed := range seeds {
				for _, m := range mappings {
					if m.source <= seed && seed <= m.source+m.length {
						seeds[i] = m.dest + (seed - m.source)
						break
					}
				}
				fmt.Println(seeds)
				fmt.Println()
			}

			mappings = []mapping{}
		}
	}

	slices.Reverse(mappings)

	for i, seed := range seeds {
		for _, m := range mappings {
			if m.source <= seed && seed <= m.source+m.length {
				seeds[i] = m.dest + (seed - m.source)
				break
			}
		}
		fmt.Println(seeds)
		fmt.Println()
	}

	fmt.Println(seeds)
	return aoc.Min(seeds...)
}

func part2(in []string) (total int) {
	seeds := []int{}

	input := aoc.StringInts(strings.Fields(strings.Split(in[0], ": ")[1]))
	fmt.Println(input)

	for i := 0; i < len(input); i += 2 {
		start := input[i]
		length := input[i+1]

		for x := start; x < start+length; x++ {
			seeds = append(seeds, x)
		}
	}
	fmt.Println(len(seeds))

	// seeds := aoc.StringInts(strings.Fields(strings.Split(in[0], ": ")[1]))

	name := "seeds"
	mappings := []mapping{}
	for _, s := range in[2:] {
		if m := strings.Split(s, " map:"); len(m) > 1 { // reset
			fmt.Println(name, seeds)

			name = strings.Split(m[0], "-")[2]
			// mapping = map[int]int{}
			continue
		}

		if m := aoc.StringInts(strings.Fields(s)); len(m) == 3 {
			dest, source, length := m[0], m[1], m[2]

			mappings = append(mappings, mapping{source, dest, length})
			fmt.Println(dest, source, length)
			continue
		}

		if s == "" {
			slices.Reverse(mappings)

			for i, seed := range seeds {
				for _, m := range mappings {
					if m.source <= seed && seed < m.source+m.length {
						seeds[i] = m.dest + (seed - m.source)
						break
					}
				}
				// fmt.Println(seeds)
				// fmt.Println()
			}

			mappings = []mapping{}
		}
	}

	slices.Reverse(mappings)

	for i, seed := range seeds {
		for _, m := range mappings {
			if m.source <= seed && seed < m.source+m.length {
				seeds[i] = m.dest + (seed - m.source)
				break
			}
		}
		fmt.Println(seeds)
		fmt.Println()
	}

	fmt.Println(seeds)
	return aoc.Min(seeds...)
}

const day = "https://adventofcode.com/2023/day/5"

func main() {
	println(day)

	// aoc.Test("test1", part1(aoc.Strings("test")), 35)
	aoc.Test("test2", part2(aoc.Strings("test")), 46)

	println("-------")

	// aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
