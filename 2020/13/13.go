package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (first int) {
	start := aoc.AtoI(in[0])
	buses := []int{}

	for _, b := range strings.FieldsFunc(in[1], func(c rune) bool { return !unicode.IsNumber(c) }) {
		if b == "x" {
			continue
		}
		buses = append(buses, aoc.AtoI(b))
	}

	for timestamp := start; first == 0; timestamp++ {
		for _, bus := range buses {
			// fmt.Println(timestamp, bus, timestamp%bus)

			if timestamp%bus == 0 {
				first = (timestamp - start) * bus
			}
		}
	}

	return
}

func part2(in string, start int) (first int) {
	buses := []int{}

	for _, b := range strings.FieldsFunc(in, func(c rune) bool { return unicode.IsPunct(c) }) {
		if b == "x" {
			buses = append(buses, 0)
			continue
		}
		buses = append(buses, aoc.AtoI(b))
	}

	fmt.Println(buses)

	for timestamp := start; first == 0; timestamp += buses[0] {
		// for timestamp := 3416; timestamp < 3418; timestamp++ {
		pos := []int{}

		for offset, bus := range buses {
			if bus == 0 {
				pos = append(pos, bus)
				continue
			}

			t := timestamp + offset

			if t%bus == 0 {
				pos = append(pos, bus)
			}
		}

		if len(pos) == len(buses) {
			first = timestamp
		}

		// fmt.Println(timestamp, buses, pos)
	}

	return
}

func part2x(in string, start int) (first int) {
	buses := []int{}

	for _, b := range strings.FieldsFunc(in, func(c rune) bool { return unicode.IsPunct(c) }) {
		if b == "x" {
			buses = append(buses, 0)
			continue
		}
		buses = append(buses, aoc.AtoI(b))
	}

	fmt.Println(buses)

	// values := []int64{}
	for offset, bus := range buses {
		if bus == 0 {
			continue
		}
		fmt.Print("(t+", offset, ") mod ", bus, "= 0, ")

	}
	fmt.Println()
	return

	// fmt.Println(aoc.LCM(values[0], values[1], values[2:]...))

	for timestamp := start; first == 0; timestamp += buses[0] {
		// for timestamp := 3416; timestamp < 3418; timestamp++ {
		pos := []int{}

		for offset, bus := range buses {
			if bus == 0 {
				pos = append(pos, bus)
				continue
			}

			t := timestamp + offset

			if t%bus == 0 {
				pos = append(pos, bus)
			}
		}

		if len(pos) == len(buses) {
			first = timestamp
		}

		// fmt.Println(timestamp, buses, pos)
	}

	return
}

const day = "https://adventofcode.com/2020/day/13"

func main() {
	println(day)
	aoc.Input(day)

	println("------- part1")

	aoc.Test("test", part1(aoc.Strings("test")), 295)
	aoc.Run("run", part1(aoc.Strings(day)))

	println("------- part2")
	aoc.Test("test", part2("17,x,13,19", 0), 3417)
	aoc.Test("test", part2("67,7,59,61", 0), 754018)
	aoc.Test("test", part2("67,x,7,59,61", 0), 779210)
	aoc.Test("test", part2("67,7,x,59,61", 0), 1261476)
	aoc.Test("test", part2("1789,37,47,1889", 0), 1202161486)

	println("------- part2x")
	// aoc.Test("test", part2x("17,x,13,19", 0), 3417)
	// aoc.Test("test", part2x("67,7,59,61", 0), 754018)
	// aoc.Test("test", part2x("67,x,7,59,61", 0), 779210)
	// aoc.Test("test", part2x("67,7,x,59,61", 0), 1261476)
	// aoc.Test("test", part2x("1789,37,47,1889", 0), 1202161486)

	aoc.Run("run", part2x(aoc.Strings(day)[1], 100000000000000))
}
