package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

// Parses a string of the form "a-b" into a Range
func parseRange(s string) Range {
	parts := strings.Split(s, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return Range{start: start, end: end}
}

// fullyContains returns true if r1 fully contains r2, false otherwise
func fullyContains(r1, r2 Range) bool {
	return r1.start <= r2.start && r1.end >= r2.end
}

func rangesOverlap(r1, r2 Range) bool {
	return r1.start <= r2.end && r2.start <= r1.end
}

func compareRanges(r1, r2 Range) (count int, overlap int) {
	if fullyContains(r1, r2) || fullyContains(r2, r1) {
		count++
	}

	if rangesOverlap(r1, r2) {
		overlap++
	}

	return
}

func part1and2(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var count int
	var overlap int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elves := strings.Split(scanner.Text(), ",")

		c, o := compareRanges(parseRange(elves[0]), parseRange(elves[1]))
		count += c
		overlap += o
	}

	fmt.Println(count, overlap)
}

func main() {
	part1and2("test.txt")
	part1and2("input.txt")
}
