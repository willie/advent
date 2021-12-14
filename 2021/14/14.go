package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

type pair struct {
	find, insert string
}

func part1(in string, iteration int) (result int) {
	// parse file
	parts := strings.Split(in, "\n\n")

	polymer := parts[0]
	// fmt.Println(polymer)
	rules := []pair{}

	// parse rules
	for _, i := range strings.Split(parts[1], "\n") {
		var p pair
		fmt.Sscanf(i, "%s -> %s", &p.find, &p.insert)
		// p.insert = p.find[:1] + p.insert + p.find[1:]
		rules = append(rules, p)
	}

	// fmt.Println(rules)

	// find locations before replace
	for i := 0; i < iteration; i++ {
		var dest string

		src := strings.Split(polymer, "")
		// fmt.Println(src)
		for i := 0; i < len(src)-1; i++ {
			dest += src[i]

			s := src[i] + src[i+1]
			for _, r := range rules {
				if r.find == s {
					dest += r.insert
				}
			}
		}
		dest += src[len(src)-1]

		polymer = dest

		// fmt.Println(polymer)

		var min, max string
		counts := map[string]int{}
		for _, i := range strings.Split(polymer, "") {
			counts[i]++

			min, max = i, i
		}

		// fmt.Println(min, max)
		for k, v := range counts {
			// fmt.Println(k, v)
			if v < counts[min] {
				min = k
			}

			if v > counts[max] {
				max = k
			}
		}
		// fmt.Println(min, max)

		result = counts[max] - counts[min]
	}

	return
}

func part2(in string, iteration int) (result int64) {
	// parse file
	parts := strings.Split(in, "\n\n")

	polymer := map[string]int64{}
	letters := map[string]int64{parts[0][:1]: 1} // add one for first letter -- it never changes

	// parse into pairs
	for i := 0; i < len(parts[0])-1; i++ {
		pair := parts[0][i : i+2]
		polymer[pair]++
	}

	// fmt.Println(polymer)
	rules := map[string]string{}

	// parse rules
	for _, i := range strings.Split(parts[1], "\n") {
		var find, insert string
		fmt.Sscanf(i, "%s -> %s", &find, &insert)
		rules[find] = insert
	}

	// fmt.Println(rules)

	// mutate
	for i := 0; i < iteration; i++ {
		p2 := map[string]int64{}

		for pair, count := range polymer {
			if insert, ok := rules[pair]; ok {
				first, second := pair[:1], pair[1:]

				p2[first+insert] += count
				p2[insert+second] += count
			}
		}

		// fmt.Println(polymer)
		polymer = p2
	}

	// fmt.Println(polymer)

	// count
	for pair, count := range polymer {
		letters[pair[1:]] += count
	}

	counts := []int64{}
	for _, count := range letters {
		counts = append(counts, count)
	}

	// fmt.Println(letters)

	max, min := aoc.Max64(counts...), aoc.Min64(counts...)
	result = max - min
	// result = (max - min) / 2
	// if (max-min)%2 == 1 {
	// 	result += 1
	// }

	return
}

const day = "https://adventofcode.com/2021/day/14"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.String("test"), 10), 1588)
	aoc.Test64("test2", part2(aoc.String("test"), 10), 1588)

	println("-------")

	aoc.Run("part1", part1(aoc.String(day), 10))
	aoc.Run64("part2", part2(aoc.String(day), 40))
}
