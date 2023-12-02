package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
	"golang.org/x/exp/maps"
)

func part1(in []string) (total int) {
	limit := map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, s := range in {
		line := strings.Split(s, ": ")
		var gameID int
		fmt.Sscanf(line[0], "Game %d:", &gameID)

		overLimit := false

		for _, subs := range strings.Split(line[1], "; ") {
			for _, pair := range strings.Split(subs, ", ") {
				var amount int
				var color string
				fmt.Sscanf(pair, "%d %s", &amount, &color)

				if amount > limit[color] {
					overLimit = true
				}
			}
		}

		if !overLimit {
			total += gameID
		}
	}

	return
}

func part2(in []string) (total int) {

	for _, s := range in {
		line := strings.Split(s, ": ")
		cubesNeeded := make(map[string]int)

		for _, subs := range strings.Split(line[1], "; ") {
			for _, pair := range strings.Split(subs, ", ") {
				var amount int
				var color string
				fmt.Sscanf(pair, "%d %s", &amount, &color)

				if current := cubesNeeded[color]; amount > current {
					cubesNeeded[color] = amount
				}
			}
		}

		total += aoc.Product(maps.Values(cubesNeeded)...)
	}

	return
}

const day = "https://adventofcode.com/2023/day/2"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 8)
	aoc.Test("test2", part2(aoc.Strings("test2")), 2286)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
