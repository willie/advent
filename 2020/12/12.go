package main

import (
	"fmt"
	"image"
	"log"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (first int) {
	loc := image.Pt(0, 0)
	direction := 0

	for _, i := range in {
		instruction := i[:1]
		amount := aoc.AtoI(i[1:])

		facing := []string{"E", "N", "W", "S"}
		step := amount / 90

		if instruction == "F" {
			instruction = facing[direction]
		}

		var delta image.Point
		switch instruction {
		case "N":
			delta = image.Pt(0, amount)
		case "S":
			delta = image.Pt(0, -amount)
		case "E":
			delta = image.Pt(amount, 0)
		case "W":
			delta = image.Pt(-amount, 0)
		case "L":
			direction = (direction + step) % 4

			// direction = (amount + direction) % 360
		case "R":
			direction = (direction + (3 * step)) % 4
			// direction = (amount + direction + 180) % 360
		default:
			log.Fatalln("unknown instruction", instruction)
		}

		fmt.Println(i, instruction, facing[direction], loc, loc.Add(delta))
		loc = loc.Add(delta)
	}

	return aoc.ManhattanDistance(0, 0, loc.X, loc.Y)
}

const day = "https://adventofcode.com/2020/day/12"

func main() {
	println(day)
	aoc.Input(day)

	println("------- part1")

	aoc.Test("test", part1(aoc.Strings("test")), 25)
	fmt.Println()
	aoc.Run("run", part1(aoc.Strings(day)))

	// println("------- part2")

	// g := aoc.LoadGrid("count")
	// x, y := g.FindFirst(empty)
	// aoc.Test("count", countVisibleOccupied(g, x, y), 8)

	// g = aoc.LoadGrid("count2")
	// x, y = g.FindFirst(empty)
	// aoc.Test("count2", countVisibleOccupied(g, x, y), 0)

	// g = aoc.LoadGrid("count3")
	// x, y = g.FindFirst(empty)
	// aoc.Test("count3", countVisibleOccupied(g, x, y), 0)

	// aoc.Test("test", part2(aoc.LoadGrid("test")), 26)

	// aoc.Run("run", part2(aoc.LoadGrid(day)))
}
