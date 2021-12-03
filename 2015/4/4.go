package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func HashMD5(buffer []byte) (hex string) {
	hash := md5.New()
	hash.Write(buffer)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func part1(in string, match string) (result int) {
	for {
		candidate := fmt.Sprintf("%s%d", in, result)
		if strings.Index(HashMD5([]byte(candidate)), match) == 0 {
			break
		}
		result++
	}

	return
}

const day = "https://adventofcode.com/2015/day/4"

func main() {
	println(day)

	test1 := []struct {
		s        string
		expected int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, t := range test1 {
		aoc.Test("test", part1(t.s, "00000"), t.expected)
	}

	println("-------")

	aoc.Run("part1", part1("ckczppom", "00000"))
	aoc.Run("part2", part1("ckczppom", "000000"))
}
