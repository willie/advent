package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

// var required = map[string]

var passportFields = map[string]bool{
	"byr": false,
	"iyr": false,
	"eyr": false,
	"hgt": false,
	"hcl": false,
	"ecl": false,
	"pid": false,
	"cid": true,
}

func part1(in []string) (valid int) {
	var current []string

	for _, i := range in {
		// fmt.Println(i)
		fields := strings.Split(i, " ")
		// fmt.Println(len(i), len(fields))
		current = append(current, fields...)

		if len(i) == 0 { // end of passport
			count := 0
			opt := false

			fields := []string{}
			// check
			for _, j := range current {
				kv := strings.Split(j, ":")
				name := kv[0]
				// value = kv[1]
				// fmt.Println(name)
				fields = append(fields, name)

				if optional, ok := passportFields[name]; ok {
					count++
					if optional {
						opt = true
					}
				}
			}

			qualifies := false

			switch {
			case count == len(passportFields):
				qualifies = true
			case (count == len(passportFields)-1) && opt:
				qualifies = true
			}

			fmt.Println(fields)

			current = []string{}
		}
	}

	return
}

const day = "https://adventofcode.com/2020/day/4"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 2)
	aoc.Run("part1", part1(aoc.Strings(day)))

	// println("-------")

	// aoc.Test("test2", part2(aoc.Strings("test")), 336)
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
