package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in string, turns int) (result [2]int) {
	starting := aoc.Ints{}

	for _, i := range strings.Split(in, ",") {
		starting = append(starting, aoc.AtoI(i))
	}

	spoken := aoc.Ints{}
	spoken = append(spoken, starting...)

	last := spoken.Last()

	for turn := len(spoken) + 1; turn <= turns; turn++ {
		this := 0

		all := spoken.AllIndex(last)
		// fmt.Println(turn, last, spoken, all)
		switch {
		case len(all) <= 1: // never been spoken
			this = 0

		default:
			// fmt.Println(turn, last, spoken)
			// fmt.Println(all)

			this = all[len(all)-1] - all[len(all)-2]
		}

		spoken = append(spoken, this)
		last = this

		// fmt.Println(turn, last)

	}

	result[0] = last

	return
}

const day = "https://adventofcode.com/2020/day/15"

func main() {
	println(day)
	aoc.Input(day)

	println("------- combined")
	fmt.Println("test", part1(aoc.String("test"), 2020), 436)
	fmt.Println("run", part1(aoc.String(day), 2020))
}
