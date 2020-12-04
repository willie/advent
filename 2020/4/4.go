package main

import (
	"fmt"
	"image/color"
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
	// "cid": true,
}

var eyeColor = aoc.NewStringSet("amb", "blu", "brn", "gry", "grn", "hzl", "oth")

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
			if _, there := passportFields[name]; !there {
				continue
			}

			value := kv[1]
			passport[name] = value
		}

		passports = append(passports, passport)
	}
	return
}

func part1(in string) (valid int) {
	for _, passport := range parse(in) {
		pass := len(passport) == len(passportFields)
		if pass {
			valid++
		}
		// fmt.Println(passport, pass)
	}
	return
}

func part2(in string) (count int) {
	for _, passport := range parse(in) {
		passportValid := true

		for k, v := range passport {
			valid := false

			switch k {
			case "byr":
				year := aoc.AtoI(v)
				valid = 1920 <= year && year <= 2002

			case "iyr":
				year := aoc.AtoI(v)
				valid = 2010 <= year && year <= 2020

			case "eyr":
				year := aoc.AtoI(v)
				valid = 2020 <= year && year <= 2030

			case "hgt":
				var h int
				var unit string
				fmt.Sscanf(v, "%d%s", &h, &unit)
				switch unit {
				case "cm":
					valid = 150 <= h && h <= 193
				case "in":
					valid = 59 <= h && h <= 76
				}

			case "hcl":
				var c color.RGBA
				n, _ := fmt.Sscanf(v, "#%02x%02x%02x", &c.R, &c.G, &c.B)
				valid = n == 3

			case "ecl":
				valid = eyeColor.Contains(v)

			case "pid":
				var pid int
				n, _ := fmt.Sscanf(v, "%9d", &pid)
				valid = n == 1 && len(v) == 9

			default:
				fmt.Println(k, v, "not validated")
				continue
			}

			if valid {
				// fmt.Println(k, "valid:\t", v)
			} else {
				passportValid = false
				// fmt.Println(k, "invalid:\t", v)
			}
		}

		if passportValid && (len(passport) == len(passportFields)) {
			count++
		}

		// fmt.Println()
	}

	return
}

const day = "https://adventofcode.com/2020/day/4"

func main() {
	println(day)

	parse(aoc.String("test"))
	aoc.Test("test1", part1(aoc.String("test")), 2)
	aoc.Run("part1", part1(aoc.String(day)))

	println("-------")

	aoc.Test("test2", part2(aoc.String("invalid")), 0)
	aoc.Test("test2", part2(aoc.String("valid")), 4)
	aoc.Run("part2", part2(aoc.String(day)))
}
