package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func indexAll(s, substr string) (pos []int) {
	p := 0

	for {
		i := strings.Index(s, substr)
		if i == -1 {
			return
		}

		p += i
		pos = append(pos, p)
		p += len(substr)

		s = s[i+len(substr):]
	}
}

type pair struct {
	find, insert string
}

func part1(in string, iteration int) (result int) {
	// parse file
	parts := strings.Split(in, "\n\n")

	polymer := parts[0]
	fmt.Println(polymer)
	rules := []pair{}

	// parse rules
	for _, i := range strings.Split(parts[1], "\n") {
		var p pair
		fmt.Sscanf(i, "%s -> %s", &p.find, &p.insert)
		// p.insert = p.find[:1] + p.insert + p.find[1:]
		rules = append(rules, p)
	}

	fmt.Println(rules)

	// find locations before replace
	for i := 0; i < iteration; i++ {
		var dest string

		src := strings.Split(polymer, "")
		fmt.Println(src)
		for i := 0; i < len(src)-1; i++ {
			dest += src[i]

			s := src[i] + src[i+1]
			for _, r := range rules {
				if r.find == s {
					dest += r.insert
				}
			}
		}
		dest += src[len(src)-1]

		polymer = dest

		fmt.Println(polymer)

		var min, max string
		counts := map[string]int{}
		for _, i := range strings.Split(polymer, "") {
			counts[i]++

			min, max = i, i
		}

		// fmt.Println(min, max)
		for k, v := range counts {
			// fmt.Println(k, v)
			if v < counts[min] {
				min = k
			}

			if v > counts[max] {
				max = k
			}
		}
		// fmt.Println(min, max)

		result = counts[max] - counts[min]
	}

	return
}

func part2(in string) (result int) {
	return
}

const day = "https://adventofcode.com/2021/day/14"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.String("test"), 10), 1588)
	// aoc.Test("test2", part2(aoc.Strings("test")), 12)

	// println("-------")

	aoc.Run("part1", part1(aoc.String(day), 10))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
