package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

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
	// fmt.Println(rl, rh, cl, ch, seatID)
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
	// fmt.Println(seatIDs)

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

	println("-------")

	fmt.Println("binary solution after some sleep, damnit")

	seatIDs := []int{}
	bc := strings.NewReplacer("F", "0", "B", "1", "R", "1", "L", "0")

	// for _, pass := range aoc.Strings("test") {
	for _, pass := range aoc.Strings(day) {
		pass = bc.Replace(pass) // convert to binary
		seatID, _ := strconv.ParseUint(pass, 2, 10)
		seatIDs = append(seatIDs, int(seatID))

		// r, _ := strconv.ParseUint(pass[:7], 2, 7)
		// c, _ := strconv.ParseUint(pass[7:], 2, 3)
		// fmt.Println(r, c, seatID)
	}

	sort.Ints(seatIDs)
	fmt.Println("part1", seatIDs[len(seatIDs)-1])

	for i, id := range seatIDs {
		if id+1 != seatIDs[i+1] {
			fmt.Println("part2", id+1)
			return
		}
	}
}
