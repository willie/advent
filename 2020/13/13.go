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

func part2(in string) (first int64) {
	buses := []int64{}

	for _, b := range strings.FieldsFunc(in, func(c rune) bool { return unicode.IsPunct(c) }) {
		if b == "x" {
			buses = append(buses, 0)
			continue
		}
		buses = append(buses, int64(aoc.AtoI(b)))
	}

	fmt.Println(buses)

	step := int64(1)                 // step is multiples of the first bus
	timestamp := int64(0)            // start at time 0
	for offset, bus := range buses { // offset, bus number
		if bus == 0 { // I substituted 0 for x
			continue
		}

		for (timestamp+int64(offset))%bus != 0 { // loop over this bus until you get 0
			timestamp += step
		}

		step *= bus // all multiples from this point will divisible
	}

	first = timestamp

	return
}

func part2mathematica(in string) (s string) {
	buses := []int{}

	for _, b := range strings.FieldsFunc(in, func(c rune) bool { return unicode.IsPunct(c) }) {
		if b == "x" {
			buses = append(buses, 0)
			continue
		}
		buses = append(buses, aoc.AtoI(b))
	}

	out := []string{}
	for offset, bus := range buses {
		if bus == 0 {
			continue
		}
		out = append(out, fmt.Sprint("(t+", offset, ") mod ", bus, "= 0"))
	}

	s = fmt.Sprint("\"", strings.Join(out, ", "), "\"")
	return
}

const day = "https://adventofcode.com/2020/day/13"

func main() {
	println(day)
	aoc.Input(day)

	println("------- part1")
	aoc.Test("test", part1(aoc.Strings("test")), 295)
	aoc.Run("run", part1(aoc.Strings(day)))

	println()
	println("------- part2")
	fmt.Println("test", part2("17,x,13,19"), 3417)
	fmt.Println("test", part2("67,7,59,61"), 754018)
	fmt.Println("test", part2("67,x,7,59,61"), 779210)
	fmt.Println("test", part2("67,7,x,59,61"), 1261476)
	fmt.Println("test", part2("1789,37,47,1889"), 1202161486)
	fmt.Println("run", part2(aoc.Strings(day)[1]))

	println()
	println("------- part2, paste into Mathematica")
	fmt.Println(part2mathematica("17,x,13,19"), 3417)
	fmt.Println(part2mathematica("67,7,59,61"), 754018)
	fmt.Println(part2mathematica("67,x,7,59,61"), 779210)
	fmt.Println(part2mathematica("67,7,x,59,61"), 1261476)
	fmt.Println(part2mathematica("1789,37,47,1889"), 1202161486)

	fmt.Println("run", part2mathematica(aoc.Strings(day)[1]))
}
