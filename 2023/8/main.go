package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
	"golang.org/x/exp/maps"
)

// It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if you have the camel follow the same instructions, you can escape the haunted wasteland!

// After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where you are now, and you have to follow the left/right instructions until you reach ZZZ.

// This format defines each node of the network individually. For example:

// RL

// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)
// Starting with AAA, you need to look up the next element based on the next left/right instruction in your input. In this example, start with AAA and go right (R) by choosing the right element of AAA, CCC. Then, L means to choose the left element of CCC, ZZZ. By following the left/right instructions, you reach ZZZ in 2 steps.

// Of course, you might not find ZZZ right away. If you run out of left/right instructions, repeat the whole sequence of instructions as necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For example, here is a situation that takes 6 steps to reach ZZZ:

// LLR

// AAA = (BBB, BBB)
// BBB = (AAA, ZZZ)
// ZZZ = (ZZZ, ZZZ)

func part1(in []string) (total int) {
	instructions := strings.Split(in[0], "")

	fmt.Println(instructions)

	type direction struct {
		left  string
		right string
	}

	directions := make(map[string]direction)
	for _, line := range in[2:] {
		line := strings.NewReplacer("=", "", ":", "", "(", "", ")", "", ",", "").Replace(line)

		var from, left, right string
		fmt.Sscanf(line, "%s %s %s", &from, &left, &right)

		d := direction{left: left, right: right}
		directions[from] = d
	}

	// fmt.Println(directions)

	// loop through the instructions over and over until we reach ZZZ
	current := "AAA"
	for {
		for _, instruction := range instructions {
			total++
			if instruction == "L" {
				current = directions[current].left
			} else {
				current = directions[current].right
			}

			// fmt.Println(current, total)

			if current == "ZZZ" {
				return
			}
		}

		if current == "ZZZ" {
			return
		}
	}

	return
}

// After examining the maps a bit longer, your attention is drawn to a curious fact: the number of nodes with names ending in A is equal to the number ending in Z! If you were a ghost, you'd probably just start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z.

// For example:

// LR

// 11A = (11B, XXX)
// 11B = (XXX, 11Z)
// 11Z = (11B, XXX)
// 22A = (22B, XXX)
// 22B = (22C, 22C)
// 22C = (22Z, 22Z)
// 22Z = (22B, 22B)
// XXX = (XXX, XXX)
// Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each left/right instruction, use that instruction to simultaneously navigate away from both nodes you're currently on. Repeat this process until all of the nodes you're currently on end with Z. (If only some of the nodes you're on end with Z, they act like any other node and you continue as normal.) In this example, you would proceed as follows:

// Step 0: You are at 11A and 22A.
// Step 1: You choose all of the left paths, leading you to 11B and 22B.
// Step 2: You choose all of the right paths, leading you to 11Z and 22C.
// Step 3: You choose all of the left paths, leading you to 11B and 22Z.
// Step 4: You choose all of the right paths, leading you to 11Z and 22B.
// Step 5: You choose all of the left paths, leading you to 11B and 22C.
// Step 6: You choose all of the right paths, leading you to 11Z and 22Z.
// So, in this example, you end up entirely on nodes that end in Z after 6 steps.

func part2(in []string) (total int) {
	instructions := strings.Split(in[0], "")

	fmt.Println(instructions)

	type direction struct {
		left  string
		right string
	}

	directions := make(map[string]direction)
	for _, line := range in[2:] {
		line := strings.NewReplacer("=", "", ":", "", "(", "", ")", "", ",", "").Replace(line)

		var from, left, right string
		fmt.Sscanf(line, "%s %s %s", &from, &left, &right)

		d := direction{left: left, right: right}
		directions[from] = d
	}

	// fmt.Println(directions)

	// loop through the instructions over and over until we reach ZZZ
	current := []string{}
	// start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z
	for _, dir := range maps.Keys(directions) {
		if strings.HasSuffix(dir, "A") {
			current = append(current, dir)
		}
	}

	fmt.Println(current)

	for {
		done := true

		for _, instruction := range instructions {
			total++

			for i, c := range current {
				if instruction == "L" {
					current[i] = directions[c].left
				} else {
					current[i] = directions[c].right
				}
			}

			done = true
			for _, c := range current {
				// if !strings.HasSuffix(c, "Z") {
				if c[2] != 'Z' {
					done = false
				} else {
					fmt.Println(c)
				}
			}

			if done {
				return
			}
		}
		if done {
			return
		}
	}

	return
}

func part2LCM(in []string) (total int) {
	instructions := strings.Split(in[0], "")

	fmt.Println(instructions)

	type direction struct {
		left  string
		right string
	}

	directions := make(map[string]direction)
	for _, line := range in[2:] {
		line := strings.NewReplacer("=", "", ":", "", "(", "", ")", "", ",", "").Replace(line)

		var from, left, right string
		fmt.Sscanf(line, "%s %s %s", &from, &left, &right)

		d := direction{left: left, right: right}
		directions[from] = d
	}

	// fmt.Println(directions)

	// loop through the instructions over and over until we reach ZZZ
	starts := []string{}
	// start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z
	for _, dir := range maps.Keys(directions) {
		if strings.HasSuffix(dir, "A") {
			starts = append(starts, dir)
		}
	}

	fmt.Println(starts)

	steps := []int64{}
	for _, current := range starts {
		step := 0

		for current[2] != 'Z' {
			for _, instruction := range instructions {
				step++

				if instruction == "L" {
					current = directions[current].left
				} else {
					current = directions[current].right
				}

				if current[2] == 'Z' {
					steps = append(steps, int64(step))
					break
				}
			}
		}
	}

	return int(lcmSlice(steps))
}

const day = "https://adventofcode.com/2023/day/8"

func main() {
	println(day)

	// aoc.Test("test1", part1(aoc.Strings("test")), 2)
	// aoc.Test("test2", part1(aoc.Strings("test2")), 6)
	aoc.Test("test3", part2(aoc.Strings("test3")), 6)
	aoc.Test("test3", part2LCM(aoc.Strings("test3")), 6)

	println("-------")

	// aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2LCM(aoc.Strings(day)))
}

func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// lcm computes the least common multiple of a and b.
func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

// lcmSlice computes the least common multiple of a slice of int64s.
func lcmSlice(nums []int64) int64 {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for _, num := range nums[1:] {
		result = lcm(result, num)
	}
	return result
}
