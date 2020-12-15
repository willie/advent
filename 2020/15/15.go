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
		switch {
		case len(all) <= 1: // never been spoken
			this = 0

		default:
			this = all[len(all)-1] - all[len(all)-2]
		}

		spoken = append(spoken, this)
		last = this
	}

	result[0] = last

	return
}

func part2(in string, turns int) (result [2]int) {
	spoken := map[int]int{}
	var last int

	for i, x := range strings.Split(in, ",") {
		last := aoc.AtoI(x)
		spoken[last] = i + 1
	}

	for turn := len(spoken) + 1; turn < turns; turn++ {
		this := 0 // assume it's never been spoken

		if prev, ok := spoken[last]; ok { // it has been spoken
			this = turn - prev // turn - previous turn
		}

		spoken[last] = turn // update the last time it was spoken
		last = this
	}

	result[0] = last

	return
}

const day = "https://adventofcode.com/2020/day/15"

func main() {
	println(day)
	aoc.Input(day)

	println("------- part 1")
	fmt.Println("test", part1(aoc.String("test"), 2020), 436)
	fmt.Println("test", part1(aoc.String("test"), 9), 4)
	// fmt.Println("run", part1(aoc.String(day), 2020))

	println("------- part 2")
	fmt.Println("test", part2(aoc.String("test"), 2020), 436)
	fmt.Println("test2", part2(aoc.String("test"), 9), 4)
	fmt.Println("run", part2(aoc.String(day), 2020))
	fmt.Println("run", part2(aoc.String(day), 30000000))
}
