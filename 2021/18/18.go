package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

type pair struct {
	left   *pair
	right  *pair
	parent *pair
	n      int
}

func (p *pair) magnitude() int {
	if p.left == nil && p.right == nil {
		return p.n
	}

	return 3*p.left.magnitude() + 2*p.right.magnitude()

}
func (p *pair) String() string {
	if p.left == nil && p.right == nil {
		return fmt.Sprintf("%d", p.n)
	}

	return fmt.Sprintf("[%s,%s]", p.left, p.right)
}

func (p *pair) depth() (count int) {
	for x := p.parent; x != nil; x = x.parent {
		count++
	}
	return
}

func (p *pair) addPair(v *pair) {
	if p.left == nil {
		p.left = v
	} else if p.right == nil {
		p.right = v
	} else {
		panic("too many")
	}
}

func (p *pair) maxdepth() (max int) {
	if p == nil {
		return 0
	}

	max = aoc.Max(p.depth(), p.left.maxdepth(), p.left.maxdepth())
	return
}

func part1(in []string) (result int) {
	// var top *pair

	for _, i := range in {
		explode(i)
	}
	return

	for _, i := range in {
		top := parse(i)
		fmt.Println(top, top.maxdepth())
	}

	return
}

func parse(in string) (top *pair) {
	top = &pair{}
	for _, s := range strings.Split(in, "") {
		switch s {
		case "[":
			p := &pair{parent: top}
			top.addPair(p)
			top = p

		case "]":
			top = top.parent

		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			top.addPair(&pair{n: aoc.AtoI(s)})
		}
	}
	top = top.left
	top.parent = nil
	return
}

type num struct {
	n     int
	depth int
}

func number(s string, depth int) num { return num{n: aoc.AtoI(s), depth: depth} }

func explode(in string) string {
	// exploded := false

	fmt.Println(in)

	numbers := []num{}

	depth := 0
	for _, s := range strings.Split(in, "") {
		switch s {
		case "[":
			depth++
		case "]":
			depth--

		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			numbers = append(numbers, number(s, depth))
		}
	}

	lastDepth := 0
	firstValue := false
	for _, n := range numbers {

		switch {
		case n.depth > lastDepth:
			for x := lastDepth; x < n.depth; x++ {
				fmt.Printf("[")
			}
			fmt.Printf("%d", n.n)
			firstValue = true

		case n.depth == lastDepth:
			if firstValue {
				fmt.Printf(",")
				firstValue = false
			} else {
				fmt.Printf("],[")
				firstValue = true
			}
			fmt.Printf("%d", n.n)

		case n.depth < lastDepth:
			for x := n.depth; x < lastDepth; x++ {
				fmt.Printf("]")
			}
			fmt.Printf("%d", n.n)
			firstValue = true
		}

		lastDepth = n.depth
	}
	for x := 0; x < lastDepth; x++ {
		fmt.Printf("]")
	}
	fmt.Println()
	return ""
}

const day = "https://adventofcode.com/2021/day/18"

func main() {
	println(day)

	explode("[[[[[9,8],1],2],3],4]")
	// aoc.Test("explode", ), 0)

	// aoc.Test("explode", explode(parse("[[[[[9,8],1],2],3],4]")), 0)
	aoc.Test("test1", part1(aoc.Strings("test")), 0)

	println("-------")

	// 	t1 = part1(aoc.Strings(day))
	// 	aoc.RunX("part1", t1)
}
