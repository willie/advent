package main

import (
	"fmt"
	"image"

	"github.com/willie/advent/aoc"
)

var intervals = aoc.NewSet(20, 60, 100, 140, 180, 220)

func part1(name string) {
	values := []int{}
	for _, s := range aoc.Strings(name) {
		op, arg := "", 0
		fmt.Sscanf(s, "%s %d", &op, &arg)
		if op == "addx" {
			values = append(values, 0) // noop
		}
		values = append(values, arg)
	}

	cycle, X, strength := 1, 1, 0
	screen := aoc.NewBlankGrid(40, 6, " ")
	current := image.Pt(0, 0)

	for _, arg := range values {
		if (current.X-1 == X) || (current.X == X) || (current.X+1 == X) {
			screen.Set(current.X, current.Y, "█")
		}

		if intervals.Contains(cycle) {
			strength += cycle * X
		}

		X += arg
		cycle++

		current.X++
		if current.X >= 40 {
			current.X = 0
			current.Y++
		}
	}

	fmt.Println(strength)
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
