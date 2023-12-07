package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/willie/advent/aoc"
	"golang.org/x/exp/maps"
)

const (
	HighCard = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func cardinality(hand string) (cardinality map[string]int) {
	cardinality = make(map[string]int)
	for _, h := range hand {
		cardinality[string(h)]++
	}
	return
}

func handType(hand string) (score int) {
	cardinality := cardinality(hand)

	values := maps.Values(cardinality)
	slices.Sort(values)
	slices.Reverse(values)

	most := values[0]
	second := 0
	if len(values) > 1 {
		second = values[1]
	}

	switch most {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if second == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if second == 2 {
			return TwoPair
		}
		return OnePair
	case 1:
		return HighCard
	}

	return 0
}

type Hand struct {
	cards string
	bid   int
}

func (h Hand) HandType() int  { return handType(h.cards) }
func (h Hand) HandType2() int { return handType2(h.cards) }

func CompareHands(a, b Hand) int {
	if a.HandType() < b.HandType() {
		return -1
	} else if a.HandType() > b.HandType() {
		return 1
	} else {
		for i := 0; i < len(a.cards); i++ {
			const strength = "23456789TJQKA"

			aa := strings.IndexByte(strength, a.cards[i])
			bb := strings.IndexByte(strength, b.cards[i])

			if aa == bb {
				continue
			} else if aa < bb {
				return -1
			} else if aa > bb {
				return 1
			}
		}
	}
	fmt.Println("WTF", a, b)
	return 0
}

func part1(in []string) (total int) {
	hands := []Hand{}

	for _, line := range in {
		hand := Hand{}
		fmt.Sscanf(line, "%s %d", &hand.cards, &hand.bid)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, CompareHands)

	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	return
}

func handType2(hand string) (score int) {
	cardinality := cardinality(hand)

	wild := 0
	if w, ok := cardinality["J"]; ok && w != 5 {
		wild = w
		delete(cardinality, "J")
	}

	values := maps.Values(cardinality)
	slices.Sort(values)
	slices.Reverse(values)

	most := values[0] + wild
	second := 0
	if len(values) > 1 {
		second = values[1]
	}

	switch most {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if second == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if second == 2 {
			return TwoPair
		}
		return OnePair
	case 1:
		return HighCard
	}

	return 0
}

func CompareHands2(a, b Hand) int {
	if a.HandType2() < b.HandType2() {
		return -1
	} else if a.HandType2() > b.HandType2() {
		return 1
	} else {
		for i := 0; i < len(a.cards); i++ {
			const strength2 = "J23456789TQKA"

			aa := strings.IndexByte(strength2, a.cards[i])
			bb := strings.IndexByte(strength2, b.cards[i])

			if aa == bb {
				continue
			} else if aa < bb {
				return -1
			} else if aa > bb {
				return 1
			}
		}
	}
	fmt.Println("WTF", a, b)
	return 0
}

func part2(in []string) (total int) {
	hands := []Hand{}

	for _, line := range in {
		hand := Hand{}
		fmt.Sscanf(line, "%s %d", &hand.cards, &hand.bid)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, CompareHands2)

	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	return
}

const day = "https://adventofcode.com/2023/day/7"

func main() {
	println(day)

	// fmt.Println(handType2("JAAJJ"), handType2("8JJJJ"), handType2("J3JJA"), handType2("TTJTT"))

	aoc.Test("test1", part1(aoc.Strings("test")), 6440)
	aoc.Test("test2", part2(aoc.Strings("test")), 5905)

	println("-------")

	aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
}
