package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in []string) (result int) {
	var gamma string
	for i := 0; i < len(in[0]); i++ {
		count := make(map[byte]int)

		for _, s := range in {
			count[s[i]-'0']++
		}

		most := "0"
		if count[1] > count[0] {
			most = "1"
		}
		// fmt.Println(count)

		gamma += most
	}

	g, err := strconv.ParseUint(gamma, 2, 30)
	if err != nil {
		log.Fatalln(err)
	}

	epsilon := strings.NewReplacer("1", "0", "0", "1").Replace(gamma)
	e, err := strconv.ParseUint(epsilon, 2, 30)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(gamma, epsilon)
	fmt.Println(g, e)

	return int(g * e)
}

func part2(in []string) (result int) {
	g := aoc.NewGrid(in)
	prefix := ""
	for i := 0; i < g.Width(); i++ {
		c := strings.Join(g.Column(i), "")

		mcv := "1"
		if strings.Count(c, "0") > strings.Count(c, "1") {
			mcv = "0"
		}
		prefix += mcv

		newG := []string{}
		for y := 0; y < g.Height(); y++ {
			r := strings.Join(g.Row(y), "")
			if strings.Index(r, prefix) == 0 {
				newG = append(newG, r)
			}
		}

		g = aoc.NewGrid(newG)

		if g.Height() == 1 {
			break
		}
	}

	O2 := strings.Join(g.Row(0), "")
	println(O2)

	g = aoc.NewGrid(in)
	prefix = ""
	for i := 0; i < g.Width(); i++ {
		c := strings.Join(g.Column(i), "")

		lcv := "0"
		if strings.Count(c, "1") < strings.Count(c, "0") {
			lcv = "1"
		}
		prefix += lcv

		newG := []string{}
		for y := 0; y < g.Height(); y++ {
			r := strings.Join(g.Row(y), "")
			if strings.Index(r, prefix) == 0 {
				newG = append(newG, r)
			}
		}

		g = aoc.NewGrid(newG)

		if g.Height() == 1 {
			break
		}
	}

	CO2 := strings.Join(g.Row(0), "")
	println(CO2)

	r1, _ := strconv.ParseUint(O2, 2, 30)
	r2, _ := strconv.ParseUint(CO2, 2, 30)
	println(r1, r2)

	return int(r1 * r2)
}

const day = "https://adventofcode.com/2021/day/3"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 198)
	aoc.Test("test2", part2(aoc.Strings("test")), 230)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
