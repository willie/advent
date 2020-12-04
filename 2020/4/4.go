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

func parse(in string) (passports []map[string]string) {
	lines := strings.Split(in, "\n\n")

	for _, p := range lines {
		p = strings.ReplaceAll(p, "\n", " ")

		passport := map[string]string{}

		fields := strings.Split(p, " ")
		for _, f := range fields {
			if len(f) == 0 {
				continue
			}

			kv := strings.Split(f, ":")

			name := kv[0]
			value := kv[1]

			passport[name] = value
		}

		passports = append(passports, passport)
	}
	return
}

func part1(in string) (valid int) {
	passports := parse(in)

	for _, passport := range passports {
		pass := false

		if len(passport) == len(passportFields) {
			pass = true
		}

		if len(passport)+1 == len(passportFields) {
			if _, there := passport["cid"]; !there {
				pass = true
			}
		}

		if pass {
			valid++
		}

		fmt.Println(passport, pass)
	}

	return
}

const day = "https://adventofcode.com/2020/day/4"

func main() {
	println(day)

	parse(aoc.String("test"))
	aoc.Test("test1", part1(aoc.String("test")), 2)
	aoc.Run("part1", part1(aoc.String(day)))

	// println("-------")

	// aoc.Test("test2", part2(aoc.Strings("test")), 336)
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
