package main

import (
	"fmt"

	"github.com/willie/advent/aoc"
)

func firstMarker(s string, uniqueN int) int {
	for i := range s {
		marker := aoc.NewSet[rune]()
		for _, c := range s[i : i+uniqueN] {
			marker.Add(c)
		}

		if len(marker.Values()) == uniqueN {
			return i + uniqueN
		}
	}

	return 0
}

func part1(name string) {
	for _, s := range aoc.Strings(name) {
		fmt.Println(firstMarker(s, 4))
	}
}

func part2(name string) {
	for _, s := range aoc.Strings(name) {
		fmt.Println(firstMarker(s, 14))
	}
}

func main() {
	part1("test.txt")
	part1("input.txt")

	fmt.Println("------")

	part2("test.txt")
	part2("input.txt")
}
