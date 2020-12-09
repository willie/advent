package main

import (
	"github.com/willie/advent/aoc"
)

func combined(in aoc.Ints, window int) (first, second int) {
	for i, next := range in[window:] {
		preamble := in[i : i+window]

		isSum := false

		for _, j := range preamble {
			for _, k := range preamble {
				if j == k || isSum {
					continue
				}

				if next == j+k {
					isSum = true
				}
			}
		}

		if !isSum {
			first = next
			break
		}
	}

	for start := 0; start < len(in); start++ {
		for end := start + 1; end < len(in); end++ {
			candidate := in[start:end]
			if first == candidate.Sum() {
				second = candidate.Min() + candidate.Max()
				return
			}
		}
	}

	return
}

const day = "https://adventofcode.com/2020/day/9"

func main() {
	println(day)
	aoc.Input(day)

	println("------- combined")

	t1, t2 := combined(aoc.LoadInts("test"), 5)
	aoc.TestX("test", t1, t2, 127, 62)

	r1, r2 := combined(aoc.LoadInts(day), 25)
	aoc.RunX("part", r1, r2)
}
