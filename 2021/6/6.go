package main

import (
	"strings"

	"github.com/willie/advent/aoc"
)

func splitInts(in string) (ints []int) {
	for _, s := range strings.Split(strings.TrimSpace(in), ",") {
		ints = append(ints, aoc.AtoI(s))
	}
	return
}

func part1(in string, days int) (result int) {
	population := splitInts(in)

	for day := 0; day < days; day++ {
		// fmt.Println(population)

		newPopulation := []int{}

		for i, p := range population {
			if p == 0 {
				population[i] = 6
				newPopulation = append(newPopulation, 8)
				continue
			}

			population[i]--
		}
		population = append(population, newPopulation...)
	}
	// fmt.Println(population)

	return len(population)
}

func part2(in string, days int) (result int64) {
	population := make([]int64, days+9)
	for _, i := range splitInts(in) {
		population[i]++
	}

	for day := 0; day < days; day++ {
		newPopulation := population[0]
		population = population[1:]

		population[6] += newPopulation
		population[8] = newPopulation
	}

	return aoc.Sum(population...)
}

const day = "https://adventofcode.com/2021/day/6"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.String("test"), 18), 26)
	aoc.Test("test1", part1(aoc.String("test"), 80), 5934)
	aoc.Test64("test2", part2(aoc.String("test"), 256), 26984457539)

	println("-------")

	aoc.Run("part1", part1(aoc.String(day), 80))
	aoc.Run("part2", part2(aoc.String(day), 256))
}
