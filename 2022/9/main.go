package main

import (
	"fmt"
	"image"

	"github.com/willie/advent/aoc"
)

var directions = map[string]image.Point{
	"L": image.Pt(-1, 0),
	"R": image.Pt(1, 0),
	"U": image.Pt(0, 1),
	"D": image.Pt(0, -1),
}

var surrounding = []image.Point{
	{-1, 1}, {0, 1}, {1, 1},
	{-1, 0}, {0, 0}, {1, 0},
	{-1, -1}, {0, -1}, {1, -1},
}

var adjacent = []image.Point{
	{-1, 1}, {0, 1}, {1, 1},
	{-1, 0} /*{0, 0},*/, {1, 0},
	{-1, -1}, {0, -1}, {1, -1},
}

func closestCandidate(A image.Point, candidates []image.Point) (closest image.Point) {
	closest = candidates[0]

	for _, candidate := range candidates[1:] {
		if aoc.Distance(A, candidate) < aoc.Distance(A, closest) {
			closest = candidate
		}
	}

	return
}

func part1(name string) {
	var head, tail image.Point

	visited := aoc.Grid2[string]{image.Pt(0, 0): "#"}

	for _, s := range aoc.Strings(name) {
		var dir string
		var steps int

		fmt.Sscanf(s, "%s %d", &dir, &steps)
		// fmt.Println(dir, steps)

		delta := directions[dir]
		for i := 0; i < steps; i++ {
			head = head.Add(delta)

			touching := false
			for _, a := range adjacent {
				if tail.Add(a) == head {
					touching = true
				}
			}

			if !touching {
				tail = closestCandidate(head, aoc.Map(tail.Add, surrounding))

				visited[tail] = "#"
			}
		}
	}

	// visited.Print(".")
	fmt.Println(len(visited))
}

func part2(name string) {
	var head image.Point
	tails := make([]image.Point, 9)

	visited := aoc.Grid2[string]{image.Pt(0, 0): "#"}

	for _, s := range aoc.Strings(name) {
		var dir string
		var steps int

		fmt.Sscanf(s, "%s %d", &dir, &steps)

		delta := directions[dir]
		for i := 0; i < steps; i++ {
			head = head.Add(delta)

			next := head
			for t := 0; t < len(tails); t++ {

				touching := false
				for _, a := range adjacent {
					if tails[t].Add(a) == next {
						touching = true
					}
				}

				if !touching {
					tails[t] = closestCandidate(next, aoc.Map(tails[t].Add, surrounding))

					if t == len(tails)-1 {
						visited[tails[t]] = "#"
					}
				}

				next = tails[t]
			}
		}
	}

	// visited.Print(".")
	fmt.Println(len(visited))
}

func main() {
	part1("test.txt")
	part1("input.txt")

	println("------")

	part2("test2.txt")
	part2("input.txt")
}
