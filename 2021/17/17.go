package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/willie/advent/aoc"
)

func part1(in string) (maxY int, velocityCount int) {
	in = strings.NewReplacer("target area: x=", "", "..", " ", ", y=", " ").Replace(in)
	var x1, x2, y1, y2 int
	fmt.Sscanf(in, "%d %d %d %d", &x1, &x2, &y1, &y2)

	target := image.Rect(x1, y1, x2, y2)
	target.Max.X++
	target.Max.Y++

	// fmt.Println(target)

	for i := 1; i <= aoc.Max(aoc.Abs(x1), aoc.Abs(x2)); i++ {
		for j := aoc.Min(y1, y2); j <= aoc.Max(aoc.Abs(y1), aoc.Abs(y2)); j++ {
			height := 0

			probe := image.Pt(0, 0)
			velocity := image.Pt(i, j)

			// iterate steps

			for probe.X <= aoc.Max(aoc.Abs(x1), aoc.Abs(x2)) && probe.Y >= aoc.Min(y1, y2) {
				// fmt.Println(probe, velocity)

				probe = probe.Add(velocity)

				if probe.Y > height {
					height = probe.Y
				}

				if velocity.X > 0 {
					velocity.X--
				} else if velocity.X < 0 {
					velocity.X++
				}

				velocity.Y--

				// break
				if probe.In(target) {
					// fmt.Println(i, j)
					if height > maxY {
						maxY = height
					}
					velocityCount++
					break
				}
			}
		}
	}

	// fmt.Println(maxY)

	return
}

const day = "https://adventofcode.com/2021/day/17"

func main() {
	println(day)

	var t1, t2 int
	t1, t2 = part1("target area: x=20..30, y=-10..-5")
	aoc.TestX("test1", t1, t2, 45, 112)

	println("-------")

	t1, t2 = part1(aoc.String(day))
	aoc.RunX("part1", t1, t2)
}
