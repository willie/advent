package main

import (
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (total int) {
	for _, s := range in {
		line := strings.Split(strings.Split(s, ": ")[1], " | ")

		winners := aoc.StringInts(strings.Fields(line[0]))
		card := aoc.StringInts(strings.Fields(line[1]))

		losers := aoc.NewSet(card...).Subtract(aoc.NewSet(winners...))
		diff := len(card) - len(losers.Values())

		score := 0
		for i := 0; i < diff; i++ {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}

		total += score
	}

	return
}

func part2(in []string) (total int) {
	cardWins := []int{}

	for _, s := range in {
		line := strings.Split(strings.Split(s, ": ")[1], " | ")

		winners := aoc.StringInts(strings.Fields(line[0]))
		card := aoc.StringInts(strings.Fields(line[1]))

		losers := aoc.NewSet(card...).Subtract(aoc.NewSet(winners...))
		won := len(card) - len(losers.Values())

		cardWins = append(cardWins, won)
	}

	cards := make([]int, len(cardWins))
	for i := range cards {
		cards[i] = 1
	}

	for i, count := range cards {
		wins := cardWins[i]
		for j := i + 1; j < len(cards) && (wins > 0); j++ {
			cards[j] += count
			wins--
		}

		cards[len(cards)-1] += wins
	}

	return aoc.Sum(cards...)
}

const day = "https://adventofcode.com/2023/day/4"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 13)
	aoc.Test("test2", part2(aoc.Strings("test")), 30)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
