package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func evalulate(in string) (result int) {
	op := '+'

	// fmt.Println("in", in)
	for i := 0; i < len(in); i++ {
		var x int

		c := rune(in[i])
		switch c {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			x = aoc.AtoI(string(c))
		case '+', '*':
			op = c
			continue
		case '(':
			basis := in[i+1:]
			p := strings.LastIndex(basis, ")")
			basis = basis[:p+1]

			fmt.Println("basis", basis)
			i += len(basis)
			x = evalulate(basis)

		case ')':
			continue
		case ' ':
			continue
		}

		// fmt.Println(result, string(op), x)

		switch op {
		case '+':
			result += x
		case '*':
			result *= x
		}
	}

	return
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
	// fmt.Println("test", part1(aoc.Strings("test")), 13632)
	// fmt.Println("run", part1(aoc.Strings(day)))

	// fmt.Println("test2", part2(aoc.Strings("test")), 848)
	// fmt.Println("run", part2(aoc.Strings(day)))
}
