package main

import (
	"fmt"

	"github.com/willie/advent/aoc"
)

func combined(in []string) (result [2]int) {
	return
}

const day = "https://adventofcode.com/2020/day/15"

func main() {
	println(day)
	aoc.Input(day)

	println("------- combined")
	fmt.Println("test", combined(aoc.Strings("test")), 165)
	fmt.Println("run", combined(aoc.Strings(day)))
}
