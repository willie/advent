package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in string) (result [2]int) {
	parts := strings.Split(in, "\n\n")

	validation := map[int][]string{}

	// parse fields, ranges
	for _, p := range strings.Split(parts[0], "\n") {
		row := strings.Split(p, ": ")

		field := row[0]
		ranges := aoc.IntSet{}

		for _, valid := range strings.Split(row[1], " or ") {
			r := strings.Split(valid, "-")
			ranges.AddMany(aoc.Series(aoc.AtoI(r[0]), aoc.AtoI(r[1])))

			for _, i := range aoc.Series(aoc.AtoI(r[0]), aoc.AtoI(r[1])) {
				validation[i] = append(validation[i], field)
			}
		}

		// fmt.Println(validation)
	}

	var first int

	for _, ticket := range strings.Split(parts[2], "\n") {
		for _, v := range strings.Split(ticket, ",") {
			if strings.Contains(v, ":") {
				continue
			}

			if v == "" {
				continue
			}

			// println(v)
			value := aoc.AtoI(v)
			if _, ok := validation[value]; !ok {
				first += value
			}
		}
	}

	result[0] = first
	// parse your ticket
	// parse nearby tickets

	return
}

const day = "https://adventofcode.com/2020/day/16"

func main() {
	println(day)
	aoc.Input(day)

	println("------- part 1")
	fmt.Println("test", part1(aoc.String("test")), 71)
	fmt.Println("run", part1(aoc.String(day)))

	println("------- part 2")
	// fmt.Println("test", part2(aoc.String("test"), 2020), 436)
	// fmt.Println("test2", part2(aoc.String("test"), 9), 4)
	// fmt.Println("run", part2(aoc.String(day), 2020))
	// fmt.Println("run", part2(aoc.String(day), 30000000))
}
