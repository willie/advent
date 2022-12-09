package main

import (
	"fmt"
	"image"
	"math"

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
		if distance(A, candidate) < distance(A, closest) {
			closest = candidate
		}
	}

	return
}

// Returns the Euclidean distance between two image.Points, p and q.
func distance(p, q image.Point) float64 {
	dx := p.X - q.X
	dy := p.Y - q.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
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
			// fmt.Println("head", head, "tail", tail)

			// current := aoc.Grid2[string]{image.Pt(0, 0): "s", tail: "T", head: "H"}
			// current.Print(".")
			// println()

			touching := false
			for _, a := range adjacent {
				if tail.Add(a) == head { // we are touching
					touching = true
				}
			}

			if !touching {
				candidates := []image.Point{}
				for _, n := range surrounding {
					candidates = append(candidates, tail.Add(n))
				}
				tail = closestCandidate(head, candidates)

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
			// fmt.Println("head", head, "tail", tail)

			next := head
			for t := 0; t < len(tails); t++ {

				touching := false
				for _, a := range adjacent {
					if tails[t].Add(a) == next { // we are touching
						touching = true
					}
				}

				if !touching {
					// candidates := []image.Point{}
					// for _, n := range surrounding {
					// 	candidates = append(candidates, tails[t].Add(n))
					// }
					// tails[t] = closestCandidate(next, candidates)

					tails[t] = closestCandidate(next, Map(surrounding, func(n image.Point) image.Point { return tails[t].Add(n) }))

					if t == len(tails)-1 {
						visited[tails[t]] = "#"
					}
				}

				next = tails[t]
			}

			// current := aoc.Grid2[string]{image.Pt(0, 0): "s"}
			// for a := len(tails) - 1; a >= 0; a++ {
			// 	current[tails[a]] = fmt.Sprint(i)
			// }

			// current[head] = "H"
			// current.Print(".")
			// println()

		}
	}

	// visited.Print(".")
	fmt.Println(len(visited))
}

func Map[T any, V any](in []T, f func(T) V) (out []V) {
	out = make([]V, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return
}

func main() {
	part1("test.txt")
	part1("input.txt")

	println("------")

	part2("test2.txt")
	part2("input.txt")
}
