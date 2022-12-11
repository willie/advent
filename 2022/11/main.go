package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/willie/advent/aoc"
)

type Monkey struct {
	items       []int
	op          func(old int) int
	test        func(item int) bool
	throw       map[bool]int
	inspections int
}

func part1(name string) {

	monkeys := []*Monkey{}

	for _, desc := range strings.Split(string(aoc.Input(name)), "\n\n") {
		record := strings.Split(desc, "\n")[1:]
		for i, s := range record {
			record[i] = strings.TrimSpace(strings.SplitAfter(s, ": ")[1])
		}

		m := &Monkey{
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

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				worry := monkey.op(item) / 3

				dest := monkey.throw[monkey.test(worry)]
				fmt.Println(item, worry, worry, dest)

				monkeys[dest].items = append(monkeys[dest].items, worry)
				monkey.inspections++
			}
			monkey.items = []int{}
		}
	}

	inspections := []int{}
	for i, m := range monkeys {
		fmt.Println("Monkey", i, ":", m.items, m.inspections)
		inspections = append(inspections, m.inspections)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	fmt.Println(inspections[0] * inspections[1])
}

func part2(name string) {

	monkeys := []*Monkey{}
	worryMod := 1

	for _, desc := range strings.Split(string(aoc.Input(name)), "\n\n") {
		record := strings.Split(desc, "\n")[1:]
		for i, s := range record {
			record[i] = strings.TrimSpace(strings.SplitAfter(s, ": ")[1])
		}

		m := &Monkey{
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

		worryMod *= divisible

		// fmt.Println(m)
		monkeys = append(monkeys, m)

		// fmt.Println(len(record), record)
	}

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				worry := monkey.op(item) % worryMod

				dest := monkey.throw[monkey.test(worry)]
				// fmt.Println(item, worry, worry, dest)

				monkeys[dest].items = append(monkeys[dest].items, worry)
				monkey.inspections++
			}
			monkey.items = []int{}
		}
	}

	inspections := []int{}
	for i, m := range monkeys {
		fmt.Println("Monkey", i, ":", m.items, m.inspections)
		inspections = append(inspections, m.inspections)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	fmt.Println(inspections, inspections[0]*inspections[1])
}

func main() {
	part1("test.txt")
	// part1("test2.txt")
	part1("input.txt")

	println("------")

	part2("test.txt")
	part2("input.txt")
}
