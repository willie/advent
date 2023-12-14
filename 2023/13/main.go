package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func loadGrids(in []string) (grids []aoc.Grid) {
	p := []string{}
	for _, i := range in {
		if i == "" {
			grids = append(grids, aoc.NewGrid(p))
			p = []string{}
			continue
		}

		p = append(p, i)
	}
	if len(p) > 0 {
		grids = append(grids, aoc.NewGrid(p))
	}
	return
}

func reverseString(s string) (reversed string) {
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return
}

func reflection(in []string) (reflected int) {
	reflected = -1
	reflections := map[int]int{}

	for _, s := range in {
		for i := 1; i < len(s); i++ {
			// iterate over the string, looking for a reflection, however small
			l, r := reverseString(s[0:i]), s[i:]
			// fmt.Println(l, r)
			m := min(len(l), len(r))
			l, r = l[:m], r[:m]

			if l == r {
				// fmt.Println(l, r)
				reflections[i]++
			}
		}
	}

	// fmt.Println(reflections)
	// max := aoc.Max(maps.Values(reflections)...)

	wtf := []int{}
	for k, v := range reflections {
		if v == len(in) {
			// return k
			wtf = append(wtf, k)
		}
	}
	if len(wtf) > 0 {
		// fmt.Println("wtf", wtf)
		return wtf[0]
	}

	return
}

func part1(in []string) (total int) {
	grids := loadGrids(in)
	for _, g := range grids {
		g.Print()
		fmt.Println()
	}
	fmt.Println(len(grids))

	var rowR, colR int

	for _, g := range grids {
		// find a perfect reflection across either a horizontal line between two rows or across a vertical line between two columns.
		rows := []string{}
		for _, row := range g.Rows() {
			rows = append(rows, strings.Join(row, ""))
		}
		rr := reflection(rows)
		if rr != -1 {
			rowR += rr
			fmt.Println(rowR)
		}

		cols := []string{}
		for _, col := range g.Columns() {
			cols = append(cols, strings.Join(col, ""))
		}
		cc := reflection(cols)
		if cc != -1 {
			colR += cc
			fmt.Println(colR)
		}
	}

	total = rowR + colR*100

	return
}

func part2(in []string) (total int) {
	grids := loadGrids(in)
	for _, g := range grids {
		g.Print()
		fmt.Println()
	}
	fmt.Println(len(grids))

	var rowR, colR int

	for _, g := range grids {
		// find a perfect reflection across either a horizontal line between two rows or across a vertical line between two columns.
		rr, cc := -1, -1

		rows := []string{}
		for _, row := range g.Rows() {
			rows = append(rows, strings.Join(row, ""))
		}
		rr = reflection(rows)

		if rr == -1 {
			cols := []string{}
			for _, col := range g.Columns() {
				cols = append(cols, strings.Join(col, ""))
			}
			cc = reflection(cols)
		}

		pixels := g.Width() * g.Height()
		for xy := 0; xy < pixels; xy++ {
			n := g.Copy()
			x, y := xy%n.Width(), xy/n.Width()
			if n.At(x, y) == "#" {
				n.Set(x, y, ".")
			} else {
				n.Set(x, y, "#")
			}

			rows := []string{}
			for _, row := range n.Rows() {
				rows = append(rows, strings.Join(row, ""))
			}
			nrr := reflection(rows)
			if nrr != -1 && nrr != rr {
				rowR += nrr
				fmt.Println("row", rowR)
				n.Print()

				break
			}

			cols := []string{}
			for _, col := range n.Columns() {
				cols = append(cols, strings.Join(col, ""))
			}
			ncc := reflection(cols)
			if ncc != -1 && ncc != cc {
				colR += ncc
				fmt.Println("col", colR)
				n.Print()

				break
			}
		}
		// }

	}

	total = rowR + colR*100

	return
}

const day = "https://adventofcode.com/2023/day/13"

func main() {
	println(day)

	aoc.Test("test", part1(aoc.Strings("test")), 405)
	// aoc.Test("test", part2(aoc.Strings("test")), 400)
	aoc.Test("test2", part1(aoc.Strings("test3")), 405)
	aoc.Test("test2", part2(aoc.Strings("test3")), 405)
	// aoc.Test("test1", part1(aoc.Strings("test1")), 405)
	// aoc.Test("test2", part1(aoc.Strings("test2")), 405)

	println("-------")

	// aoc.Run("part1", part1(aoc.Strings(day)))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}

// You note down the patterns of ash (.) and rocks (#) that you see as you walk (your puzzle input); perhaps by carefully analyzing these patterns, you can figure out where the mirrors are!

// For example:

// #.##..##.
// ..#.##.#.
// ##......#
// ##......#
// ..#.##.#.
// ..##..##.
// #.#.##.#.

// #...##..#
// #....#..#
// ..##..###
// #####.##.
// #####.##.
// ..##..###
// #....#..#
// To find the reflection in each pattern, you need to find a perfect reflection across either a horizontal line between two rows or across a vertical line between two columns.

// In the first pattern, the reflection is across a vertical line between two columns; arrows on each of the two columns point at the line between the columns:

// 123456789
//     ><
// #.##..##.
// ..#.##.#.
// ##......#
// ##......#
// ..#.##.#.
// ..##..##.
// #.#.##.#.
//     ><
// 123456789
// In this pattern, the line of reflection is the vertical line between columns 5 and 6. Because the vertical line is not perfectly in the middle of the pattern, part of the pattern (column 1) has nowhere to reflect onto and can be ignored; every other column has a reflected column within the pattern and must match exactly: column 2 matches column 9, column 3 matches 8, 4 matches 7, and 5 matches 6.

// The second pattern reflects across a horizontal line instead:

// 1 #...##..# 1
// 2 #....#..# 2
// 3 ..##..### 3
// 4v#####.##.v4
// 5^#####.##.^5
// 6 ..##..### 6
// 7 #....#..# 7
// This pattern reflects across the horizontal line between rows 4 and 5. Row 1 would reflect with a hypothetical row 8, but since that's not in the pattern, row 1 doesn't need to match anything. The remaining rows match: row 2 matches row 7, row 3 matches row 6, and row 4 matches row 5.

// To summarize your pattern notes, add up the number of columns to the left of each vertical line of reflection; to that, also add 100 multiplied by the number of rows above each horizontal line of reflection. In the above example, the first pattern's vertical line has 5 columns to its left and the second pattern's horizontal line has 4 rows above it, a total of 405.

// Find the line of reflection in each of the patterns in your notes. What number do you get after summarizing all of your notes?
