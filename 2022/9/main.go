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

// Returns the closest candidate to A in the candidates slice.
func closestCandidate(A image.Point, candidates []image.Point) image.Point {
	// Set the closest candidate to the first candidate in the slice.
	closest := candidates[0]

	// Iterate over the remaining candidates.
	for _, candidate := range candidates[1:] {
		// If the current candidate is closer to A than the current closest candidate,
		// set the current candidate as the new closest candidate.
		if distance(A, candidate) < distance(A, closest) {
			closest = candidate
		}
	}

	return closest
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
		fmt.Println(dir, steps)

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

				fmt.Println()
				visited[tail] = "#"
			}
		}
	}

	visited.Print(".")
	fmt.Println(len(visited))
}

func main() {
	part1("test.txt")
	part1("input.txt")

	// println("------")

	// part2("test.txt")
	// part2("input.txt")
}
