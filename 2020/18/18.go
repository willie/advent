package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func evalulate(in string) (result int) {
	in = strings.ReplaceAll(in, " ", "")

	var sub func(string) (int, string)
	sub = func(s string) (int, string) {
		op := '+'
		x := 0
		r := 0

		for len(s) > 0 && s[0] != ')' {
			c := s[0]
			s = s[1:]

			switch c {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				x = int(c - '0')
			case '+', '*':
				op = rune(c)
				continue

			case '(':
				x, s = sub(s)
				s = s[1:]
			}

			switch op {
			case '+':
				r += x
			case '*':
				r *= x
			}
		}

		return r, s
	}

	r, _ := sub(in)
	return r
}

func part1(in []string) (result int) {
	for _, i := range in {
		result += evalulate(i)
	}
	return
}

const day = "https://adventofcode.com/2020/day/18"

func main() {
	println(day)
	aoc.Input(day)

	fmt.Println("test", evalulate("1 + 2 * 3 + 4 * 5 + 6"), 71)
	fmt.Println("test", evalulate("2 * 3 + (4 * 5)"), 26)
	fmt.Println("test", evalulate("1 + (2 * 3) + (4 * (5 + 6))"), 51)
	fmt.Println("test", evalulate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"), 13632)
	fmt.Println("test", part1(aoc.Strings("test")), 13632)
	fmt.Println("run", part1(aoc.Strings(day)))

	// fmt.Println("test2", part2(aoc.Strings("test")), 848)
	// fmt.Println("run", part2(aoc.Strings(day)))
}
