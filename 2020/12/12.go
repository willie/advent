package main

import (
	"image"
	"log"

	"github.com/willie/advent/aoc"
)

func delta(direction string, amount int) (d image.Point) {
	switch direction {
	case "N":
		d = image.Pt(0, amount)
	case "S":
		d = image.Pt(0, -amount)
	case "E":
		d = image.Pt(amount, 0)
	case "W":
		d = image.Pt(-amount, 0)
	}
	return
}

func rotate(i image.Point, direction string, amount int) (n image.Point) {
	for n = i; amount != 0; amount = amount - 90 {
		switch direction {
		case "L":
			n.X, n.Y = -n.Y, n.X
		case "R":
			n.X, n.Y = n.Y, -n.X
		}
	}
	return
}

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

		switch instruction {
		case "N", "S", "E", "W":
			loc = loc.Add(delta(instruction, amount))
		case "L":
			direction = (direction + step) % 4
		case "R":
			direction = (direction + (3 * step)) % 4
		default:
			log.Fatalln("unknown instruction", instruction)
		}
	}

	return aoc.ManhattanDistance(0, 0, loc.X, loc.Y)
}

func part2(in []string) (first int) {
	ship := image.Pt(0, 0)
	waypoint := image.Pt(10, 1)

	for _, i := range in {
		instruction := i[:1]
		amount := aoc.AtoI(i[1:])

		switch instruction {
		case "N", "S", "E", "W":
			waypoint = waypoint.Add(delta(instruction, amount))
		case "L", "R":
			waypoint = rotate(waypoint, instruction, amount)
		case "F":
			ship = ship.Add(waypoint.Mul(amount))
		}
	}

	return aoc.ManhattanDistance(0, 0, ship.X, ship.Y)
}

const day = "https://adventofcode.com/2020/day/12"

func main() {
	println(day)
	aoc.Input(day)

	println("------- part1")

	aoc.Test("test", part1(aoc.Strings("test")), 25)
	aoc.Run("run", part1(aoc.Strings(day)))

	println("------- part2")
	aoc.Test("test", part2(aoc.Strings("test")), 286)
	aoc.Run("run", part2(aoc.Strings(day)))
}
