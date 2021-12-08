package main

import (
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (result int) {
	for _, i := range in {
		parts := strings.Split(i, " | ")

		for _, p := range strings.Split(parts[1], " ") {
			switch len(p) {
			case 3:
				// 7
				result += 1
			case 4:
				// 4
				result += 1
			case 2:
				// 1
				result += 1
			case 7:
				// 8
				result += 1
			}
		}
	}

	return
}

// acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf
// 8 5 2 3 7 9 6 4 0 1 | 5 3 5 3

func part2(in []string) (result int) {
	for _, i := range in {
		// figure out 1, 7, 4, 8
		parts := strings.Split(strings.ReplaceAll(i, "| ", ""), " ")
		if len(parts) != 14 {
			continue
		}

		output := make([]int, 14)
		for i := range output {
			output[i] = -1
		}

		var seven, four aoc.StringSet

		// build known
		for i, p := range parts {
			n := 0

			switch len(p) {
			case 2:
				n = 1
			case 3:
				n = 7
				seven = aoc.NewStringSet(strings.Split(p, "")...)
			case 4:
				n = 4
				four = aoc.NewStringSet(strings.Split(p, "")...)
			case 7:
				n = 8
			default:
				continue
			}

			output[i] = n
		}

		// build unknown
		for i, v := range output {
			if v != -1 {
				continue
			}

			p := aoc.NewStringSet(strings.Split(parts[i], "")...)

			switch len(p) {
			case 5: // 2, 3, 5
				// assume 2
				output[i] = 2

				// 3
				if p.ContainsAll(seven.Values()) {
					output[i] = 3
					continue
				}

				// 5
				count := 0
				for _, v := range p.Values() {
					if four.Contains(v) {
						count++
					}
				}
				if count == 3 {
					output[i] = 5
				}

			case 6: // 0, 6, 9
				// assume 6
				output[i] = 6

				// 6
				if p.ContainsAll(four.Values()) {
					output[i] = 9
					continue
				}

				// 0
				if p.ContainsAll(seven.Values()) {
					output[i] = 0
					continue
				}
			}
		}

		total := 0
		for i := 10; i < 14; i++ {
			total *= 10
			total += output[i]
		}
		result += total
	}

	return
}

const day = "https://adventofcode.com/2021/day/8"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 26)
	aoc.Test("test2", part2([]string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"}), 5353)
	aoc.Test("test2", part2(aoc.Strings("test")), 61229)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
