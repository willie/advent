package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (result int) {
	// for _, i := range in {
	// 	// result += nice(i)
	// }
	return
}

func part2(in []string) (result int) {
	// for _, i := range in {
	// 	// result += nice2(i)
	// }
	return
}

func lights(g aoc.Grid, in string) (result int) {
	in = strings.Replace(in, "turn ", "", 1)

	var cmd string
	var start, end image.Point
	fmt.Sscanf(in, "%s %d,%d through %d,%d", &cmd, &start.X, &start.Y, &end.X, &end.Y)

	fmt.Println(start, end)

	for x := start.X; x <= end.X; x++ {
		for y := start.Y; y <= end.Y; y++ {

			g.Set(x, y, "on")
		}
	}

	return g.Count("on")
}

const day = "https://adventofcode.com/2015/day/6"

func main() {
	println(day)

	test1 := []struct {
		s    string
		nice int
	}{
		{"turn on 0,0 through 999,999", 1000 * 1000},
		{"toggle 0,0 through 999,0", 1000},
		{"turn off 499,499 through 500,500", 4},
	}

	g := aoc.NewBlankGrid(1000, 1000, "off")

	for _, t := range test1 {
		aoc.Test("test", lights(g, t.s), t.nice)
	}

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
