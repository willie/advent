package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in string) (count int) {
	groups := []aoc.StringSet{}
	in = strings.TrimSpace(in)
	in = strings.ReplaceAll(in, "\n", "|")
	in = strings.ReplaceAll(in, "||", "\n")
	in = strings.ReplaceAll(in, "|", "")
	for _, p := range strings.Split(in, "\n") {
		g := aoc.NewStringSet()

		for _, c := range p {
			letter := string(c)
			g.Add(letter)
		}

		groups = append(groups, g)
	}

	for _, g := range groups {
		count += len(g)
	}
	return
}

func part2(in string) (total int) {
	// | is my record delimiter, seemed easiest
	in = strings.TrimSpace(in)
	in = strings.ReplaceAll(in, "\n\n", "|")
	in += "|"

	personCount := 1
	answerCount := map[rune]int{}

	for _, s := range in {
		switch s {
		case '|': // new group
			for _, x := range answerCount {
				if x == personCount {
					total++
				}
			}

			personCount = 1
			answerCount = map[rune]int{}

		case '\n': // new person
			personCount++

		default: // add to the count map
			answerCount[s]++
		}
	}

	return
}

func part2scan(in string) (total int) {
	groupTotals := []int{}

	// groups
	for _, i := range strings.Split(in, "\n\n") {
		people := 0
		answerCount := map[rune]int{}
		groupTotals = append(groupTotals, 0)

		// per user
		scanner := bufio.NewScanner(strings.NewReader(i))
		for scanner.Scan() {
			people++

			for _, c := range scanner.Text() {
				answerCount[c]++
			}
		}

		// tally
		for _, answers := range answerCount {
			if answers == people {
				groupTotals[len(groupTotals)-1]++
			}
		}
	}

	fmt.Println(groupTotals)
	return aoc.Sum(groupTotals...)
}

func part2set(in string) (total int) {
	groupTotals := []int{}

	// groups
	for _, i := range strings.Split(in, "\n\n") {
		people := 0
		commonAnswers := aoc.StringSet{}

		// per user
		scanner := bufio.NewScanner(strings.NewReader(i))
		for scanner.Scan() {
			answers := aoc.StringSet{}

			for _, c := range scanner.Text() {
				answers.Add(string(c))
			}

			if people == 0 {
				commonAnswers = answers
			} else {
				commonAnswers = commonAnswers.Subtract(answers)
			}

			people++
		}

		groupTotals = append(groupTotals, len(commonAnswers))
	}

	fmt.Println(groupTotals)
	return aoc.Sum(groupTotals...)
}

const day = "https://adventofcode.com/2020/day/6"

func main() {
	println(day)

	println("------- part 1")

	aoc.Test("test1", part1(aoc.String("test")), 11)
	aoc.Run("part1", part1(aoc.String(day)))

	println("------- part 2")

	aoc.Test("test2", part2(aoc.String("test")), 6)
	aoc.Run("part2", part2(aoc.String(day)))

	println("------- redo using scanner")
	aoc.Test("test2", part2scan(aoc.String("test")), 6)
	aoc.Run("part2", part2scan(aoc.String(day)))

	println("------- redo using set")
	aoc.Test("test2", part2scan(aoc.String("test")), 6)
	aoc.Run("part2", part2scan(aoc.String(day)))
}
