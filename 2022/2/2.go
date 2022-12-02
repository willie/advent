package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var moves = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func losingMove(x int) (y int) {
	y = x - 1
	if y == 0 {
		y = 3
	}
	return
}

func winningMove(x int) (y int) {
	y = x + 1
	if y == 4 {
		y = 1
	}
	return
}

func part1(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var totalScore int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		opponent := moves[fields[0]]
		you := moves[fields[1]]
		score := you

		switch {
		case opponent == you:
			score += 3
		case losingMove(opponent) == you:
			score += 0
		default:
			score += 6
		}

		totalScore += score
	}

	fmt.Println(totalScore)
}

const (
	lose = iota + 1
	draw
	win
)

func part2(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var totalScore int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		opponent := moves[fields[0]]
		endState := moves[fields[1]]
		score := 0

		switch endState {
		case lose:
			score = losingMove(opponent) + 0
		case draw:
			score = opponent + 3
		case win:
			score = winningMove(opponent) + 6
		}

		totalScore += score
	}

	fmt.Println(totalScore)
}

const day = "https://adventofcode.com/2022/day/2"

func main() {
	println(day)

	part1("test.txt")
	part1("input.txt")

	println("-------")

	part2("test.txt")
	part2("input.txt")
}
