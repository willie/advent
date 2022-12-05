package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func part2(name string) {
	stacks := map[int]*stack[string]{}

	for _, s := range aoc.Strings(name) {
		if strings.Contains(s, "[") {
			c := 1
			for i := 1; i < len(s); i += 4 {
				if stacks[c] == nil {
					stacks[c] = &stack[string]{}
				}

				value := s[i : i+1]
				if value != " " {
					stacks[c].pushBottom(value)
				}
				c++
			}
		} else if strings.Contains(s, "move") {
			var count, from, to int
			fmt.Sscanf(s, "move %d from %d to %d", &count, &from, &to)

			crates := stacks[from].popN(count)
			stacks[to].push(crates...)
		}
	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i].top())
	}
	fmt.Println()
}

func part1(name string) {
	stacks := map[int]*stack[string]{}

	for _, s := range aoc.Strings(name) {
		if strings.Contains(s, "[") {
			c := 1
			for i := 1; i < len(s); i += 4 {
				if stacks[c] == nil {
					stacks[c] = &stack[string]{}
				}

				value := s[i : i+1]
				if value != " " {
					stacks[c].pushBottom(value)
				}
				c++
			}
		} else if strings.Contains(s, "move") {
			var count, from, to int
			fmt.Sscanf(s, "move %d from %d to %d", &count, &from, &to)

			for i := 0; i < count; i++ {
				x := stacks[from].pop()
				stacks[to].push(x)
			}
		}

	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i].top())
	}
	fmt.Println()
}

func main() {
	part1("test.txt")
	part1("input.txt")

	fmt.Println("------")

	part2("test.txt")
	part2("input.txt")
}
