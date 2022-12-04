package main

import (
	"fmt"

	"github.com/willie/advent/aoc"
)

func part1and2(name string) {
	var count, overlap int

	for _, s := range aoc.Strings(name) {
		var r1, r2 Range
		fmt.Sscanf(s, "%d-%d,%d-%d", &r1.start, &r1.end, &r2.start, &r2.end)

		if r1.Contains(r2) || r2.Contains(r1) {
			count++
		}

		if r1.Overlaps(r2) {
			overlap++
		}
	}

	fmt.Println(count, overlap)
}

func main() {
	part1and2("test.txt")
	part1and2("input.txt")
}
