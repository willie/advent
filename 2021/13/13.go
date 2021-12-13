package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/willie/advent/aoc"
)

const (
	invalid = "-1"
)

func part1(in string) (result int) {
	parts := strings.Split(in, "\n\n")

	dots := []image.Point{}
	for _, i := range strings.Split(parts[0], "\n") {
		x, y := 0, 0
		// fmt.Println(i)
		fmt.Sscanf(i, "%d,%d", &x, &y)
		dots = append(dots, image.Pt(x, y))
	}

	x, y := aoc.MaxXY(dots)
	// fmt.Println(x, y)
	g := aoc.NewBlankGrid(x+1, y+1, ".")
	// g.Print()

	for _, i := range dots {
		// fmt.Println(i.X, i.Y)
		g.Set(i.X, i.Y, "#")
	}

	// g.Print()

	// fold?
	folds := []image.Point{}
	for _, i := range strings.Split(parts[1], "\n") {
		n := 0
		axis := ""

		i = strings.ReplaceAll(i, "=", " = ")
		fmt.Sscanf(i, "fold along %s = %d", &axis, &n)
		fmt.Println(axis, n)
		if axis == "x" {
			folds = append(folds, image.Point{n, y})
		} else {
			folds = append(folds, image.Point{x, n})
		}
	}

	// fmt.Println(folds)

	// translate?
	fold := 7

	g2 := aoc.NewBlankGrid(g.Width(), fold, ".")
	g2.Iterate(func(x, y int, s string) (next bool) {
		v := g.Get(x, y, "-1")
		g2.Set(x, y, v)
		return true
	})

	for i := 1; i <= fold; i++ {
		for j := 0; j <= x; j++ {
			s := g.Get(j, fold+i, "-1")
			if s == "#" {
				g2.Set(j, fold-i, s)
			}
		}
	}

	// g2.Print()

	g2.Iterate(func(x, y int, s string) (next bool) {
		if s == "#" {
			result++
		}
		return true
	})
	return
}

func part2(in string) (result int) {
	return
}

const day = "https://adventofcode.com/2021/day/13"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.String("test")), 17)
	// aoc.Test("test2", part2(aoc.String("test")), 195)

	println("-------")

	aoc.Run("part1", part1(aoc.String(day)))
	// aoc.Run("part2", part2(aoc.String(day)))
}
