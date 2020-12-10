package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/willie/advent/aoc"
)

func permutation(current int, in aoc.Ints) (perms []aoc.Ints) {
	adapters := aoc.NewIntSet(in...)

	for i, v := range in {
		for x := v; x < v+3; x++ {
			if adapters.Contains(x) {
				rest := permutation(x, //range of ints after x)
				for _, r := range rest {
				combined := append(// range of ints after x, r...)
				perms = append(perms, combined)
				}
			}
		}

		current = v
	}

	return
}

func combined(in aoc.Ints) (first, second int) {
	sort.Ints(in)

	differences := map[int]int{3: 1}
	current := 0

	for i, v := range in {
		diff := v - current

		if diff > 3 {
			log.Fatal("diff too big", i, v, diff)
		}

		differences[diff]++
		current = v
	}

	fmt.Println(len(in), differences)
	first = differences[1] * differences[3]
	return
}

const day = "https://adventofcode.com/2020/day/10"

func main() {
	println(day)
	aoc.Input(day)

	println("------- combined")

	t1, t2 := combined(aoc.LoadInts("test"))
	aoc.TestX("test", t1, t2, 7*5, 8)

	t1, t2 = combined(aoc.LoadInts("test2"))
	aoc.TestX("test2", t1, t2, 22*10, 19208)

	r1, _ := combined(aoc.LoadInts(day))
	aoc.RunX("part", r1)
}
