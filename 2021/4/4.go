package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func isBingo(called aoc.StringSet, board aoc.Grid) (bingo bool, winner []string) {
	possibilities := append(board.Columns(), board.Rows()...)

	for _, poss := range possibilities {
		if called.ContainsAll(poss) {
			return true, poss
		}
	}

	return
}

func part1(in string) (result int) {
	sections := strings.Split(in, "\n\n")
	numbers := strings.Split(sections[0], ",")

	boards := []aoc.Grid{}
	for i := 1; i < len(sections); i++ {
		g := aoc.NewBlankGrid(5, 5, ".")

		for y, row := range strings.Split(sections[i], "\n") {
			for x, f := range strings.Fields(row) {
				g.Set(x, y, f)
			}
		}

		boards = append(boards, g)
	}

	called := aoc.NewStringSet()
	var winningBoard aoc.Grid
	var poss []string
	var last int

	for _, turn := range numbers {
		last = aoc.AtoI(turn)
		called.Add(turn)

		bingo := false

		for _, board := range boards {
			bingo, poss = isBingo(called, board)
			if bingo {
				winningBoard = board
				break
			}
		}

		if bingo {
			fmt.Println(poss)
			break
		}
	}

	unmarked := 0
	winningBoard.Iterate(func(x, y int, s string) bool {
		if !called.Contains(s) {
			unmarked += aoc.AtoI(s)
		}
		return true
	})

	fmt.Println(unmarked, last)
	return unmarked * last
}

func part2(in string) (result int) {
	sections := strings.Split(in, "\n\n")
	numbers := strings.Split(sections[0], ",")

	boards := []aoc.Grid{}
	for i := 1; i < len(sections); i++ {
		g := aoc.NewBlankGrid(5, 5, ".")

		for y, row := range strings.Split(sections[i], "\n") {
			for x, f := range strings.Fields(row) {
				g.Set(x, y, f)
			}
		}

		boards = append(boards, g)
	}

	called := aoc.NewStringSet()
	var winningBoard aoc.Grid
	winningBoards := aoc.NewIntSet()
	var last int

	for _, turn := range numbers {
		last = aoc.AtoI(turn)
		called.Add(turn)

		for b, board := range boards {
			if winningBoards.Contains(b) {
				continue
			}

			bingo, poss := isBingo(called, board)
			if bingo {
				if len(winningBoards.Values()) == len(boards)-1 {
					winningBoard = board
					fmt.Println(poss)
				}

				winningBoards.Add(b)
			}
		}

		if len(winningBoards.Values()) == len(boards) {
			break
		}
	}

	unmarked := 0
	winningBoard.Iterate(func(x, y int, s string) bool {
		if !called.Contains(s) {
			unmarked += aoc.AtoI(s)
		}
		return true
	})

	fmt.Println(unmarked, last)
	return unmarked * last
}

const day = "https://adventofcode.com/2021/day/4"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.String("test")), 4512)
	aoc.Test("test2", part2(aoc.String("test")), 1924)

	println("-------")

	aoc.Run("part1", part1(aoc.String(day)))
	aoc.Run("part2", part2(aoc.String(day)))
}
