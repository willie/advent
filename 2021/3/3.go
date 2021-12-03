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
		fmt.Println(count)

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
	return
}

const day = "https://adventofcode.com/2021/day/3"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 198)
	// aoc.Test("test2", part2(aoc.Strings("test")), 900)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
