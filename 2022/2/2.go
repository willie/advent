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

func lose(x int) (y int) {
	y = x - 1
	if y == 0 {
		y = 3
	}
	return
}

func win(x int) (y int) {
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

		opponentMove := moves[fields[0]]
		yourMove := moves[fields[1]]
		score := yourMove

		switch {
		case opponentMove == yourMove:
			score += 3
		case lose(opponentMove) == yourMove:
			score += 0
		default:
			score += 6
		}

		totalScore += score
	}

	fmt.Println(totalScore)
}

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

		opponentMove := moves[fields[0]]
		endState := fields[1]
		score := 0

		switch moves[endState] {
		case 1: // lose
			score = lose(opponentMove) + 0
		case 2: // draw
			score = opponentMove + 3
		case 3: // win
			score = win(opponentMove) + 6
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
