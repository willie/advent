package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func HashMD5(buffer []byte) (hex string) {
	hash := md5.New()
	hash.Write(buffer)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func part1(in []string) (result int) {
	for _, i := range in {
		result += nice(i)
	}
	return
}

func part2(in []string) (result int) {
	for _, i := range in {
		result += nice2(i)
	}
	return
}

func nice(in string) (nice int) {
	// count vowels
	vowels := 0
	for _, i := range in {
		for _, v := range "aeiou" {
			if i == v {
				vowels++
			}
		}
	}
	if vowels < 3 {
		// println("vowels", vowels)
		return 0
	}

	// look for doubles
	double := false
	var prev rune
	for _, i := range in {
		if i == prev {
			double = true
			break
		}
		prev = i
	}
	if !double {
		// println("no doubles")
		return 0
	}

	// look for weird
	weird := false
	for _, s := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(in, s) {
			weird = true
			break
		}
	}
	if weird {
		// println("weird")
		return 0
	}

	return 1
}

func nice2(in string) (nice int) {
	pairs := false
	for _, a := range "abcdefghijklmnopqrstuvwxyz" {
		for _, b := range "abcdefghijklmnopqrstuvwxyz" {
			ab := string(a) + string(b)

			first := strings.Index(in, ab)
			last := strings.LastIndex(in, ab)

			if first+1 < last {
				pairs = true
				break
			}
		}
	}
	if !pairs {
		// println("no pair")
		return 0
	}

	repeat := false
	for i := 0; i < len(in)-2; i++ {
		if in[i] == in[i+2] {
			repeat = true
			break
		}
	}
	if !repeat {
		// println("no repeat")
		return 0
	}

	return 1
}

const day = "https://adventofcode.com/2015/day/5"

func main() {
	println(day)

	test1 := []struct {
		s    string
		nice int
	}{
		{"ugknbfddgicrmopn", 1},
		{"aaa", 1},
		{"jchzalrnumimnmhp", 0},
		{"haegwjzuvuyypxyu", 0},
		{"dvszwmarrgswjxmb", 0},
	}

	for _, t := range test1 {
		aoc.Test("test", nice(t.s), t.nice)
	}

	test2 := []struct {
		s    string
		nice int
	}{
		{"qjhvhtzxzqqjkmpb", 1},
		{"xxyxx", 1},
		{"uurcxstgmygtbstg", 0},
		{"ieodomkazucvgmuy", 0},
	}

	for _, t := range test2 {
		// println(t.s)
		aoc.Test("test2", nice2(t.s), t.nice)
	}

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
