package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(name string) {
	for i, pairs := range strings.Split(aoc.String(name), "\n\n") {
		pair := strings.Split(pairs, "\n")

		var first, second any
		if err := json.Unmarshal([]byte(pair[0]), &first); err != nil {
			panic(err)
		}
		if err := json.Unmarshal([]byte(pair[1]), &second); err != nil {
			panic(err)
		}

		// all = append(all, first, second)

		// if equal(first, second) <= 0 {
		// 	sum += idx + 1
		// }

		fmt.Println(i, first, second)
	}
}

func main() {
	part1("test.txt")
	// part1("input.txt")
}
