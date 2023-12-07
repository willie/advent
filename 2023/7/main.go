package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/willie/advent/aoc"
	"golang.org/x/exp/maps"
)

const strength = "23456789TJQKA"

// Every hand is exactly one type. From strongest to weakest, they are:

// Five of a kind, where all five cards have the same label: AAAAA
// Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// High card, where all cards' labels are distinct: 23456

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

	switch len(cardinality) {
	case 1:
		return FiveOfAKind
	case 2:
		// either FourOfAKind or FullHouse
		for _, v := range cardinality {
			if v == 4 {
				return FourOfAKind
			}
		}

		return FullHouse
	case 3:
		// either ThreeOfAKind or TwoPair
		for _, v := range cardinality {
			if v == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	case 5:
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

	fmt.Println("what:", values, wild)

	most := values[0]
	second := 0
	if len(values) > 1 {
		second = values[1]
	}

	// for i, v := range cardinality {
	// 	if v == most {
	// 		cardinality[i] += wild
	// 		break
	// 	}
	// }

	most += wild

	// fmt.Println(wild, cardinality)

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

const strength2 = "J23456789TQKA"

func CompareHands2(a, b Hand) int {
	if a.HandType2() < b.HandType2() {
		return -1
	} else if a.HandType2() > b.HandType2() {
		return 1
	} else {
		for i := 0; i < len(a.cards); i++ {
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

	fmt.Println(handType2("JAAJJ"), handType2("8JJJJ"), handType2("J3JJA"), handType2("TTJTT"))

	fmt.Println(hands)
	slices.SortFunc(hands, CompareHands2)
	// for _, hand := range hands {
	// 	fmt.Println(hand, handType2(hand.cards))
	// }

	// fmt.Println(hands)
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	return
}

const day = "https://adventofcode.com/2023/day/7"

func main() {
	println(day)

	// aoc.Test("test1", part1(aoc.Strings("test")), 6440)
	aoc.Test("test2", part2(aoc.Strings("test")), 5905)

	println("-------")

	// aoc.Run("part1", part1(aoc.Strings(day)))
	aoc.Run("part2", part2(aoc.Strings(day)))
	// aoc.Run("part2", part2(aoc.Strings(day)))
}
