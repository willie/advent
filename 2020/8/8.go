package main

import (
	"log"
	"strings"

	"github.com/willie/advent/aoc"
)

func combined(in []string) (count, count2 int) {
	return
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
		// fmt.Println(current)

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

		if ip >= len(program) {
			break
		}
	}

	return
}

type line struct {
	instruction string
	argument    int
}

type program []line

func loadProgram(in []string) (program program) {
	for _, i := range in {
		fields := strings.Fields(i)
		if len(fields) != 2 {
			log.Fatalln("fields != 2", fields)
		}

		program = append(program, line{
			instruction: fields[0],
			argument:    aoc.AtoI(fields[1]),
		})
	}
	return
}

func (program program) count(instruction string) (count int) {
	for _, l := range program {
		if l.instruction == instruction {
			count++
		}
	}
	return
}

func (program program) run() (accumulator int, normalTermination bool) {
	ip := 0
	lineCounter := map[int]int{}

	for ip < len(program) {
		current := program[ip]
		// fmt.Println(ip, current)

		switch current.instruction {
		case "acc":
			accumulator += current.argument
			ip++
		case "jmp":
			ip += current.argument
		case "nop":
			ip++
		default:
			log.Fatalln("unrecognized instruction", current.instruction)
		}

		lineCounter[ip]++
		if lineCounter[ip] > 1 {
			return
		}
	}

	normalTermination = true
	return
}

func part2(in string) (accumulator int) {
	in = strings.TrimSpace(in)
	for i := 1; i <= strings.Count(in, "acc"); i++ {

		newIn := replaceNth(in, "jmp", "nop", i)
		p := loadProgram(strings.Split(newIn, "\n"))
		accumulator, normalTermination := p.run()

		if normalTermination {
			return accumulator
		}
	}
	return
}

func part2x(in []string) (accumulator int) {
	tmp := make([]string, len(in))

	for i := range in {
		copy(tmp, in)

		tmp[i] = strings.NewReplacer("jmp", "nop", "nop", "jmp").Replace(tmp[i])
		p := loadProgram(tmp)
		accumulator, normalTermination := p.run()

		if normalTermination {
			return accumulator
		}
	}
	return
}

const day = "https://adventofcode.com/2020/day/8"

func main() {
	println(day)
	aoc.Input(day)

	println("------- part 1")

	aoc.Test("test", part1(aoc.Strings("test")), 5)
	aoc.Run("run", part1(aoc.Strings(day)))

	println("------- part 2")

	aoc.Test("test", part2(aoc.String("test")), 8)
	// aoc.Test("test2", part2(aoc.Strings("test2")), 126)
	aoc.Run("run", part2(aoc.String(day)))
	aoc.Run("run", part2x(aoc.Strings(day)))

	// println("------- combined")

	// t1, t2 := combined(aoc.Strings("test"))
	// aoc.TestX("test", t1, t2, 4, 32)

	// r1, r2 := combined(aoc.Strings(day))
	// aoc.RunX("part", r1, r2)
}

// Replace the nth occurrence of old in s by new.
func replaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}
