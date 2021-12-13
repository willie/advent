package main

import (
	"strings"

	"github.com/willie/advent/aoc"
)

type path []string

func (p path) duplicate() (dup path) {
	dup = make(path, len(p))
	copy(dup, p)
	return
}

func (p path) count(c string) (result int) {
	if c == strings.ToLower(c) {
		for _, i := range p {
			if c == i {
				result++
			}
		}
	}
	return
}

type graph map[string]path

func follow(g graph, current path, c string) (paths []path) {
	if current.count(c) > 0 {
		return
	}

	current = append(current, c)

	if c == "end" {
		paths = append(paths, current)
		return
	}

	for _, f := range g[c] {
		paths = append(paths, follow(g, current.duplicate(), f)...)
	}
	return
}

func part1(in []string) (result int) {
	g := make(graph)
	for _, i := range in {
		path := strings.Split(i, "-")
		start, end := path[0], path[1]

		g[start] = append(g[start], end)
		g[end] = append(g[end], start)
	}

	for k, v := range g {
		g[k] = aoc.NewStringSet(v...).Remove("start").Values()
	}

	p := follow(g, path{}, "start")

	return len(p)
}

func follow2(g graph, current path, c string) (paths []path) {
	current = append(current, c)

	gt2 := 0
	count := map[string]int{}
	for _, p := range current {
		if p == strings.ToLower(p) {
			count[p] += 1

			if count[p] > 2 {
				return
			}

			if count[p] == 2 {
				gt2++
				if gt2 > 1 {
					return
				}
			}
		}
	}

	if c == "end" {
		paths = append(paths, current)
		return
	}

	for _, f := range g[c] {
		paths = append(paths, follow2(g, current.duplicate(), f)...)
	}

	return
}

func part2(in []string) (result int) {
	g := make(graph)
	for _, i := range in {
		path := strings.Split(i, "-")
		start, end := path[0], path[1]

		g[start] = append(g[start], end)
		g[end] = append(g[end], start)
	}

	for k, v := range g {
		g[k] = aoc.NewStringSet(v...).Remove("start").Values()
	}

	p := follow2(g, path{}, "start")

	return len(p)
}

const day = "https://adventofcode.com/2021/day/12"

func main() {
	println(day)

	aoc.Test("test1", part1(aoc.Strings("test")), 10)
	aoc.Test("test1", part1(aoc.Strings("test1")), 19)
	aoc.Test("test1", part1(aoc.Strings("test2")), 226)

	aoc.Test("test2", part2(aoc.Strings("test")), 36)
	aoc.Test("test2", part2(aoc.Strings("test1")), 103)
	aoc.Test("test2", part2(aoc.Strings("test2")), 3509)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
