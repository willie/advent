package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (total int) {
	times := aoc.StringInts(strings.Fields(strings.Split(in[0], ": ")[1]))
	record := aoc.StringInts(strings.Fields(strings.Split(in[1], ": ")[1]))

	total = 1
	for i, time := range times {
		distance := 0

		for button := 1; button < time; button++ {
			traveled := (time - button) * button

			if traveled > record[i] {
				distance++
			}
		}

		total *= distance
	}

	return
}

func part2(in []string) (total int) {
	times := []int{aoc.AtoI(strings.ReplaceAll(strings.Split(in[0], ": ")[1], " ", ""))}
	record := []int{aoc.AtoI(strings.ReplaceAll(strings.Split(in[1], ": ")[1], " ", ""))}

	fmt.Println(times, record)

	total = 1
	for i, time := range times {
		distance := 0

		for button := 1; button < time; button++ {
			traveled := (time - button) * button

			if traveled > record[i] {
				distance++
			}
		}

		total *= distance
	}

	return
}

const day = "https://adventofcode.com/2023/day/6"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 288)
	aoc.Test("test2", part2(aoc.Strings("test")), 71503)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
