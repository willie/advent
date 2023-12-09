package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
	"golang.org/x/exp/maps"
)

type direction struct {
	left  string
	right string
}

func part1(in []string) (total int) {
	instructions := strings.Split(in[0], "")

	directions := make(map[string]direction)
	for _, line := range in[2:] {
		var from, left, right string
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &from, &left, &right)
		directions[from] = direction{left: left, right: right}
	}

	// loop through the instructions over and over until we reach ZZZ
	current := "AAA"
	for {
		for _, instruction := range instructions {
			total++
			if instruction == "L" {
				current = directions[current].left
			} else {
				current = directions[current].right
			}

			if current == "ZZZ" {
				return
			}
		}
	}
}

func part2(in []string) (total int) {
	instructions := strings.Split(in[0], "")

	directions := make(map[string]direction)
	for _, line := range in[2:] {
		var from, left, right string
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &from, &left, &right)
		directions[from] = direction{left: left, right: right}
	}

	starts := []string{}
	for _, dir := range maps.Keys(directions) {
		if dir[2] == 'A' {
			starts = append(starts, dir)
		}
	}

	steps := []int{}
	for _, current := range starts {
		step := 0

		for current[2] != 'Z' {
			for _, instruction := range instructions {
				step++

				if instruction == "L" {
					current = directions[current].left
				} else {
					current = directions[current].right
				}

				if current[2] == 'Z' {
					steps = append(steps, step)
					fmt.Println(current, step)
					break
				}
			}
		}
	}

	fmt.Println(steps)
	return aoc.LCM(steps...)
}

const day = "https://adventofcode.com/2023/day/8"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 2)
	aoc.Test("test2", part1(aoc.Strings("test2")), 6)
	aoc.Test("test3", part2(aoc.Strings("test3")), 6)
	aoc.Test("test4", part2(aoc.Strings("test4")), 23)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
