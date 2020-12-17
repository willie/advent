package main

import (
	"fmt"

	"github.com/willie/advent/aoc"
)

type world map[int]map[int]map[int]string

func (w world) Set(x, y, z int, s string) {
	// does y exist
	if _, ok := w[x]; !ok {
		w[x] = map[int]map[int]string{}
	}

	// does z exist
	if _, ok := w[x][y]; !ok {
		w[x][y] = map[int]string{}
	}

	// set it
	w[x][y][z] = s
}

const (
	inactive = "."
)

func (w world) Get(x, y, z int) (s string) {
	// x
	if _, ok := w[x]; !ok {
		return
	}

	// y
	if _, ok := w[x][y]; !ok {
		return
	}

	// z
	if _, ok := w[x][y][z]; !ok {
		return
	}

	return w[x][y][z]
}

func (w world) Get2(x, y, z int) (s string) {
	// probably faster, but uglier
	if wy, ok := w[x]; ok {
		if wz, ok := wy[y]; ok {
			if v, ok := wz[z]; ok {
				return v
			}
		}
	}

	return
}

func part1(in []string) (result [2]int) {
	return
}

const day = "https://adventofcode.com/2020/day/17"

func main() {
	w := world{}
	w.Set(1, 1, 1, "#")

	fmt.Println(w.Get(1, 1, 1), w.Get(0, 0, 0))
	fmt.Println(w.Get2(1, 1, 1), w.Get2(0, 0, 0))

	println(day)
	aoc.Input(day)

	fmt.Println("test", part1(aoc.Strings("test")), 71)
	// fmt.Println("run", part1(aoc.Strings(day)))
}
