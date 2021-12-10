package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (result int) {
	for _, i := range in {
		stack := []string{}

		for _, c := range strings.Split(i, "") {
			open := "([{<"
			close := ")]}>"
			value := map[string]int{
				")": 3,
				"]": 57,
				"}": 1197,
				">": 25137,
			}

			if strings.Contains(open, c) {
				stack = append(stack, c)
				continue
			}

			o := string(open[strings.Index(close, c)])

			if o != stack[len(stack)-1] {
				result += value[c]
				// fmt.Printf("expected %s, found %s\n", stack[len(stack)-1], c)
				break
			}

			stack = stack[:len(stack)-1]
		}
	}

	return
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func part2(in []string) (result int) {
	open := "([{<"
	close := ")]}>"

	scores := []int{}

	for _, i := range in {
		stack := []string{}
		corrupted := false

		for _, c := range strings.Split(i, "") {

			if strings.Contains(open, c) {
				stack = append(stack, c)
				continue
			}

			o := string(open[strings.Index(close, c)])

			if o != stack[len(stack)-1] {
				// result += value[c]
				// fmt.Printf("expected %s, found %s\n", stack[len(stack)-1], c)
				corrupted = true
				break
			}

			stack = stack[:len(stack)-1]
		}

		if corrupted {
			continue
		}

		// complete
		complete := []string{}
		score := 0

		reverse(stack)
		for _, c := range stack {
			idx := strings.Index(open, c)
			complete = append(complete, string(close[idx]))

			score *= 5
			score += idx + 1
		}

		scores = append(scores, score)
	}

	fmt.Println(scores)
	sort.Ints(scores)
	fmt.Println(scores)

	return scores[len(scores)/2]
}

const day = "https://adventofcode.com/2021/day/10"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 26397)
	aoc.Test("test2", part2(aoc.Strings("test")), 288957)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
