package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/willie/advent/aoc"
)

func uint(b string) (i uint64) {
	i, _ = strconv.ParseUint(b, 2, 36)
	return
}

func toString(i uint64) (b string) {
	b = fmt.Sprintf("%036b", i)
	return
}

func applyMask1(value uint64, mask string) (result uint64) {
	v := toString(value)
	var r [36]byte

	for i, c := range mask {
		switch c {
		case 'X':
			r[i] = v[i]
		case '1':
			r[i] = '1'
		case '0':
			r[i] = '0'
		}
	}

	return uint(string(r[:]))
}

func applyMask2(value uint64, mask string) (result []uint64) {
	v := toString(value)
	results := make([][36]byte, 1)

	for i, c := range mask {
		switch c {
		case '0':
			for j := range results {
				results[j][i] = v[i]
			}
		case '1':
			for j := range results {
				results[j][i] = '1'
			}
		case 'X':
			addtional := make([][36]byte, 0)

			for j := range results {
				results[j][i] = '0'

				// add one
				addtional = append(addtional, results[j])
				addtional[len(addtional)-1][i] = '1'
			}
			results = append(results, addtional...)
		}

	}

	for _, r := range results {
		v := uint(string(r[:]))
		result = append(result, v)
	}
	return
}

func combined(in []string) (result [2]uint64) {
	mem := map[uint64]uint64{}
	mem2 := map[uint64]uint64{}

	var mask string

	for _, i := range in {
		switch i[:3] {
		case "mas":
			fmt.Sscanf(i, "mask = %s", &mask)
		case "mem":
			var loc, v uint64
			fmt.Sscanf(i, "mem[%d] = %d", &loc, &v)

			// part 1
			value := applyMask1(v, mask)
			mem[loc] = value

			// part 2
			for _, addr := range applyMask2(loc, mask) {
				mem2[addr] = v
			}

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
	// fmt.Println("test", combined(aoc.Strings("test")), 165)
	fmt.Println("test2", combined(aoc.Strings("test2")), 208)
	fmt.Println("run", combined(aoc.Strings(day)))
}
