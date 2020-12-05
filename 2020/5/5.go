package main

import (
	"fmt"
	"sort"

	"github.com/willie/advent/aoc"
)

func calcPass(pass string) (r, c, seatID int) {
	rl := 0
	rh := 127

	cl := 0
	ch := 7

	for _, c := range pass {
		p := string(c)

		rd := ((rh + 1) - rl) / 2
		cd := ((ch + 1) - cl) / 2

		switch p {
		case "F":
			rh -= rd
		case "B":
			rl += rd
		case "R":
			cl += cd
		case "L":
			ch -= cd
		}
	}

	seatID = rl*8 + cl
	fmt.Println(rl, rh, cl, ch, seatID)
	return
}

func part1(in []string) (seatID int) {
	for _, pass := range in {
		_, _, id := calcPass(pass)
		if id > seatID {
			seatID = id
		}
	}

	return
}

func part2(in []string) (seatID int) {
	seatIDs := []int{}

	for _, pass := range in {
		_, _, id := calcPass(pass)
		seatIDs = append(seatIDs, id)
	}

	sort.Ints(seatIDs)
	fmt.Println(seatIDs)

	assigned := aoc.NewIntSet(seatIDs...)
	for i := seatIDs[1]; i < seatIDs[len(seatIDs)-2]; i++ {
		if assigned.Contains(i-1) && !assigned.Contains(i) && assigned.Contains(i+1) {
			return i
		}
	}

	return
}

const day = "https://adventofcode.com/2020/day/5"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 820)
	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
