package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(name string) {
	sum := 0
	packets := []any{}

	for i, pairs := range strings.Split(aoc.String(name), "\n\n") {
		pair := strings.Split(pairs, "\n")

		var left, right any
		json.Unmarshal([]byte(pair[0]), &left)
		json.Unmarshal([]byte(pair[1]), &right)

		if compare(left, right) < 0 {
			sum += i + 1
		}

		packets = append(packets, left, right)
	}

	dividers := []any{[]any{[]any{float64(2)}}, []any{[]any{float64(6)}}} // either this or more unmarshalling

	packets = append(packets, dividers...)
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	key := 1
	for i, packet := range packets {
		for _, divider := range dividers {
			if compare(packet, divider) == 0 {
				key *= i + 1
			}
		}
	}

	fmt.Println(sum, key)
}

func jsonNumberToI(in any) (i int, ok bool) {
	if f, ok := in.(float64); ok {
		return int(f), true
	}
	return
}

func getOrMakeList(in any) (list []any) {
	switch in.(type) {
	case []any, []float64:
		list = in.([]any)
	case float64:
		list = []any{in}
	}
	return
}

// compare returns < 0 if right, 0 if equal, > 0 if wrong
func compare(left, right any) int {
	if l, ok := jsonNumberToI(left); ok {
		if r, ok := jsonNumberToI(right); ok {
			return l - r
		}
	}

	ll, rl := getOrMakeList(left), getOrMakeList(right)
	for i := 0; i < len(ll) && i < len(rl); i++ {
		if c := compare(ll[i], rl[i]); c != 0 {
			return c
		}
	}

	return len(ll) - len(rl)
}

func main() {
	part1("test.txt")
	part1("input.txt")
}
