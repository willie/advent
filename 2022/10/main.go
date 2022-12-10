package main

import (
	"fmt"
	"image"

	"github.com/willie/advent/aoc"
)

type instruction struct {
	op  string
	arg int
}

var intervals = aoc.NewSet(20, 60, 100, 140, 180, 220)

func part1(name string) {
	instructions := []instruction{}
	for _, s := range aoc.Strings(name) {
		var i instruction
		fmt.Sscanf(s, "%s %d", &i.op, &i.arg)

		if i.op == "addx" {
			instructions = append(instructions, instruction{op: "noop"})
		}
		instructions = append(instructions, i)
	}

	cycle, X, strength := 1, 1, 0
	screen := aoc.NewBlankGrid(40, 6, " ")
	current := image.Pt(0, 0)

	for _, i := range instructions {
		if current.X >= 40 {
			current.X = 0
			current.Y++
		}

		if (current.X-1 == X) || (current.X == X) || (current.X+1 == X) {
			screen.Set(current.X, current.Y, "█")
		}

		if intervals.Contains(cycle) {
			strength += cycle * X
			println(cycle, X)
		}

		if i.op == "addx" {
			X += i.arg
		}

		cycle++
		current.X++
	}

	fmt.Println(name, X, strength)
	screen.Print()
	fmt.Println()
}

func main() {
	// part1("test.txt")
	part1("test2.txt")
	part1("input.txt")

	// println("------")

	// part2("test2.txt")
	// part2("input.txt")
}
