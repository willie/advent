package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

type Monkey struct {
	items []int
	op    func(old int) int
	test  func(item int) bool
	throw map[bool]int
}

func part1(name string) {

	monkeys := []Monkey{}

	for _, desc := range strings.Split(string(aoc.Input(name)), "\n\n") {
		record := strings.Split(desc, "\n")[1:]
		for i, s := range record {
			record[i] = strings.TrimSpace(strings.SplitAfter(s, ": ")[1])
		}

		m := Monkey{
			items: aoc.Map(aoc.AtoI, strings.Split(record[0], ", ")),
		}

		op, value := "", ""
		fmt.Sscanf(record[1], "new = old %s %s", &op, &value)

		if value == "old" {
			m.op = func(old int) int { return old * old }
		} else if op == "*" {
			m.op = func(old int) int { return old * aoc.AtoI(value) }
		} else if op == "+" {
			m.op = func(old int) int { return old + aoc.AtoI(value) }
		}

		var divisible, t, f int
		fmt.Sscanf(record[2], "divisible by  %d", &divisible)
		fmt.Sscanf(record[3], "throw to monkey %d", &t)
		fmt.Sscanf(record[4], "throw to monkey %d", &f)

		m.test = func(item int) bool { return item%divisible == 0 }
		m.throw = map[bool]int{true: t, false: f}

		// fmt.Println(m)
		monkeys = append(monkeys, m)

		fmt.Println(len(record), record)
	}

	// for _, s := range aoc.Strings(name) {

	// }

	fmt.Println(monkeys)
}

func main() {
	part1("test.txt")
	// part1("test2.txt")
	// part1("input.txt")

	// println("------")

	// part2("test2.txt")
	// part2("input.txt")
}
