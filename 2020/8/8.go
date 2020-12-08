package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/willie/advent/aoc"
)

func combined(in []string) (count, count2 int) {
	return
}

type line struct {
	instruction string
	argument    int
}

func part1(in []string) (accumulator int) {
	program := []line{}

	for _, i := range in {
		fields := strings.Fields(i)
		program = append(program, line{
			instruction: fields[0],
			argument:    aoc.AtoI(fields[1]),
		})
	}

	// fmt.Println(program)

	ip := 0
	lineCounter := map[int]int{}

	for {
		current := program[ip]
		fmt.Println(current)

		switch current.instruction {
		case "nop":
			ip++
		case "acc":
			accumulator += current.argument
			ip++
		case "jmp":
			ip += current.argument
		default:
			log.Fatalln("unrecognized instruction", current.instruction)
		}

		lineCounter[ip]++
		if lineCounter[ip] > 1 {
			break
		}
	}

	return
}

func part2(in []string) (count int) {
	return
}

const day = "https://adventofcode.com/2020/day/8"

func main() {
	println(day)

	println("------- part 1")

	aoc.Test("test", part1(aoc.Strings("test")), 5)
	aoc.Run("run", part1(aoc.Strings(day)))

	// println("------- part 2")

	// aoc.Test("test", part2(aoc.Strings("test")), 32)
	// aoc.Test("test2", part2(aoc.Strings("test2")), 126)
	// aoc.Run("run", part2(aoc.Strings(day)))

	// println("------- combined")

	// t1, t2 := combined(aoc.Strings("test"))
	// aoc.TestX("test", t1, t2, 4, 32)

	// r1, r2 := combined(aoc.Strings(day))
	// aoc.RunX("part", r1, r2)
}
