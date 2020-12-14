package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/willie/advent/aoc"
)

func combined(in []string) (result [2]uint64) {
	mem := map[uint64]uint64{}
	mem2 := map[uint64]uint64{}

	var AND, OR uint64 // AND is mask, OR is value
	var mask string

	for _, i := range in {
		switch i[:3] {
		case "mas":
			fmt.Sscanf(i, "mask = %s", &mask)

			AND, _ = strconv.ParseUint(
				strings.NewReplacer("1", "0", "X", "1").
					Replace(mask), 2, 36)

			OR, _ = strconv.ParseUint(
				strings.NewReplacer("X", "0").
					Replace(mask), 2, 36)

		case "mem":
			var loc, v uint64
			fmt.Sscanf(i, "mem[%d] = %d", &loc, &v)
			value := (v & AND) | OR

			mem[loc] = value

			// loc = (loc & AND) | OR

		default:
			log.Fatalln(in)
		}
	}

	for _, v := range mem {
		result[0] += v
	}

	for _, v := range mem2 {
		result[1] += v
	}

	return
}

const day = "https://adventofcode.com/2020/day/14"

func main() {
	println(day)
	aoc.Input(day)

	println("------- combined")
	fmt.Println("test", combined(aoc.Strings("test")), 165, 208)
	fmt.Println("run", combined(aoc.Strings(day)), 6559449933360)
}
