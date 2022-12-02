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

func part1(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Keep track of the total score
	var totalScore int

	// Read each line of the input
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		opponentMove := fields[0]
		yourMove := fields[1]
		score := moves[yourMove]

		switch {
		case moves[opponentMove] == moves[yourMove]:
			score += 3
		case (opponentMove == "A" && yourMove == "Z") || (opponentMove == "B" && yourMove == "X") || (opponentMove == "C" && yourMove == "Y"):
			score += 0
		default:
			score += 6
		}

		totalScore += score
	}

	fmt.Println(totalScore)
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

func part2(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalScore int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		opponentMove := fields[0]
		endState := fields[1]

		var score int

		switch moves[endState] {
		case 1: // lose
			score = lose(moves[opponentMove]) + 0
		case 2: // draw
			score = moves[opponentMove] + 3
		case 3: // win
			score = win(moves[opponentMove]) + 6
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
