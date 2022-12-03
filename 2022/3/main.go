package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func score(c rune) int {
	if int('a') <= int(c) && int(c) <= int('z') {
		return int(c) - int('a') + 1
	}

	return int(c) - int('A') + 27
}

func findCommonRunes(input []string) []rune {
	// first string
	set := make(map[rune]bool)
	for _, c := range input[0] {
		set[c] = true
	}

	// all the rest
	for i := 1; i < len(input); i++ {
		for c := range set {
			if !strings.ContainsRune(input[i], c) {
				delete(set, c)
			}
		}
	}

	runes := make([]rune, 0, len(set))
	for c := range set {
		runes = append(runes, c)
	}
	return runes
}

func part1(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		commonItems := findCommonRunes([]string{s[:len(s)/2], s[len(s)/2:]})

		for _, c := range commonItems {
			total += score(c)
		}

	}

	fmt.Println(total)
}

func part2(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	total := 0
	rucksacks := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}

	for i := 0; i < len(rucksacks); i += 3 {
		group := rucksacks[i : i+3]
		// badgeItems := findBadgeItems(group)
		badgeItems := findCommonRunes(group)

		for _, c := range badgeItems {
			total += score(c)
		}
	}

	fmt.Println(total)
}

const day = "https://adventofcode.com/2022/day/3"

func main() {
	println(day)

	part1("test.txt")
	part1("input.txt")

	println("-------")

	part2("test.txt")
	part2("input.txt")
}
