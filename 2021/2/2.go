package main

import (
	"image"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (result int) {
	p := image.Point{}

	for _, s := range in {
		cmd := strings.Split(s, " ")
		amount := aoc.AtoI(cmd[1])

		switch cmd[0] {
		case "forward":
			p.X += amount
		case "up":
			p.Y -= amount
		case "down":
			p.Y += amount
		}
	}

	return p.X * p.Y
}

func part2(in []string) (result int) {
	p := image.Point{}
	aim := 0

	for _, s := range in {
		cmd := strings.Split(s, " ")
		amount := aoc.AtoI(cmd[1])

		switch cmd[0] {
		case "forward":
			p.X += amount
			p.Y += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		}
	}

	return p.X * p.Y
}

const day = "https://adventofcode.com/2021/day/2"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 150)
	aoc.Test("test2", part2(aoc.Strings("test")), 900)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
