package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (total int) {
	limit := map[string]int{"red": 12, "green": 13, "blue": 14}

	games := []int{}
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
					fmt.Println("over limit", gameID, amount, color)
					overLimit = true
				}

				// fmt.Printf("%d: %d %s\n", gameID, amount, color)
			}
		}

		if !overLimit {
			games = append(games, gameID)
		}
	}

	return aoc.Sum(games...)
}

const day = "https://adventofcode.com/2023/day/2"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 8)
	// aoc.Test("test2", part2(aoc.Strings("test2")), 281)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
