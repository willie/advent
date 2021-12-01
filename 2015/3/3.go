package main

import (
	"image"

	"github.com/willie/advent/aoc"
)

type world map[int]map[int]int

func (w world) Add(pt image.Point, s int) {
	// does y exist
	if _, ok := w[pt.X]; !ok {
		w[pt.X] = map[int]int{}
	}

	// set it
	w[pt.X][pt.Y] += s
}

func (w world) Iterate(f func(x, y, s int) bool) bool {
	for x, wy := range w {
		for y, s := range wy {
			if !f(x, y, s) {
				return false
			}
		}
	}
	return true
}

func part1(in string) (result int) {
	w := world{}

	var s image.Point
	w.Add(s, 1)

	for _, c := range in {
		switch c {
		case '^':
			s.Y++
		case 'v':
			s.Y--
		case '>':
			s.X++
		case '<':
			s.X--
		}

		w.Add(s, 1)
	}

	w.Iterate(func(x, y, s int) bool {
		if s > 0 {
			result++
		}
		return true
	})

	return
}

func part2(in string) (result int) {
	w := world{}

	var s, r image.Point
	w.Add(s, 1)

	for i, c := range in {
		pt := &s
		if i%2 != 0 {
			pt = &r
		}

		switch c {
		case '^':
			pt.Y++
		case 'v':
			pt.Y--
		case '>':
			pt.X++
		case '<':
			pt.X--
		}

		w.Add(*pt, 1)
	}

	w.Iterate(func(x, y, s int) bool {
		if s > 0 {
			result++
		}
		return true
	})

	return
}

const day = "https://adventofcode.com/2015/day/3"

func main() {
	println(day)

	test1 := []struct {
		s         string
		expected  int
		expected2 int
	}{
		{">", 2, 3},
		{"^>v<", 4, 3},
		{"^v^v^v^v^v", 2, 11},
	}

	for _, t := range test1 {
		aoc.Test("test", part1(t.s), t.expected)
	}

	test2 := []struct {
		s        string
		expected int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, t := range test2 {
		aoc.Test("test2", part2(t.s), t.expected)
	}

	// test2 := []struct {
	// 	s        string
	// 	expected int
	// }{
	// 	{"^v", 3},
	// 	{"^>v<", 3},
	// 	{"^v^v^v^v^v", 11},
	// }

	// aoc.Test("test1", part1(aoc.LoadInts("test")), 7)
	// aoc.Test("test2", part2(aoc.LoadInts("test")), 5)

	println("-------")

	aoc.Run("part1", part1(aoc.String(day)))
	aoc.Run("part2", part2(aoc.String(day)))
}
