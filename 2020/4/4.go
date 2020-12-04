package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/willie/advent/aoc"
)

var passportFields = aoc.NewStringSet("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")
var eyeColor = aoc.NewStringSet("amb", "blu", "brn", "gry", "grn", "hzl", "oth")

func parse(in string) (passports []map[string]string) {
	for _, p := range strings.Split(strings.TrimSpace(in), "\n\n") {
		p = strings.ReplaceAll(p, "\n", " ")

		passport := map[string]string{}

		for _, f := range strings.Split(p, " ") {
			kv := strings.Split(f, ":")

			name := kv[0]
			if !passportFields.Contains(name) {
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
		if len(passport) == len(passportFields) {
			valid++
		}
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

			if !valid {
				passportValid = false
			}
		}

		if passportValid && (len(passport) == len(passportFields)) {
			count++
		}
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
