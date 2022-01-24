package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/willie/advent/aoc"
)

/*

	// now we have a map of letter counts at each position
	// for each letter, find the most common position
	maxPositions := make(map[rune][]int)

	for r, letterCount := range letterCounts {
		maxCount := -1
		maxPositions[r] = []int{}
		for position, count := range letterCount {
			if count > maxCount {
				maxCount = count
				maxPositions[r] = []int{position}
			} else if count == maxCount {
				maxPositions[r] = append(maxPositions[r], position)
			}
		}
	}

	// now we have a map of letter positions
	// for each letter, find the most common position
	maxPositions2 := make(map[rune][]int)

	for r, positions := range maxPositions {
		maxCount := -1
		maxPositions2[r] = []int{}
		for _, position := range positions {
			if _, ok := maxPositions2[r]; !ok {
				maxPositions2[r] = []int{position}
			} else if position < maxPositions2[r][0] {
				maxPositions2[r] = []int{position}
			} else if position == maxPositions2[r][0] {
				maxPositions2[r] = append(maxPositions2[r], position)
			}
		}
	}
*/

type wordscore struct {
	word    string
	score   int
	letters aoc.StringSet
}

func process(words []string) {
	letterCounts := make(map[string]map[int]int)

	for _, word := range words {
		for i, c := range strings.Split(word, "") {
			if _, ok := letterCounts[c]; !ok {
				letterCounts[c] = make(map[int]int)
			}

			letterCounts[c][i]++
		}
	}

	// vowels := aoc.NewStringSet("a", "e", "i", "o", "u")

	multiplier := make(map[string]map[int]int)

	for c, letterCount := range letterCounts {

		lowCount := math.MaxInt
		maxCount := 0
		for _, count := range letterCount {
			if count > maxCount {
				maxCount = count
			}

			if count < lowCount {
				lowCount = count
			}
		}

		rangeCount := (maxCount - lowCount) + 1
		rangeCount = int(float64(rangeCount) * 0.5)
		total := 0

		multiplier[c] = make(map[int]int)
		for i, count := range letterCount {
			factor := 1
			// if count > (rangeCount / 2) {
			// 	factor = 2
			// }

			// factor := 1
			// if count == maxCount && !vowels.Contains(c) {
			// 	factor = 2
			// }

			multiplier[c][i] = factor
			total += count
		}

		fmt.Println(c, lowCount, maxCount, rangeCount, letterCount, ",", total)

	}

	wordScore := make(map[string]int)
	for _, word := range words {
		for i, c := range strings.Split(word, "") {
			// if vowels.Contains(c) {
			// 	continue
			// }

			// wordScore[word] += letterCounts[c][i] * multiplier[c][i]
			lc := letterCounts[c][i]
			// lc *= lc * lc
			// lc *= lc * multiplier[c][i]

			wordScore[word] += lc
		}
	}

	// print out the words in score order
	scores := []wordscore{}
	for word, score := range wordScore {
		letters := aoc.NewStringSet(strings.Split(word, "")...)
		// no dupes
		if len(letters) != 5 {
			continue
		}

		scores = append(scores, wordscore{word, score, letters})
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	for i := 0; i < 10; i++ {
		fmt.Println(scores[i].word, scores[i].score)
	}

	maxscore := 0
	maxwords := ""
	for i, score := range scores {
		f := score.letters
		if len(f) != 5 {
			continue
		}

		for _, score2 := range scores[i+1:] {
			f2 := score2.letters
			if len(f2) != 5 {
				continue
			}

			if len(f.Subtract(f2)) == len(f2) {
				if score.score+score2.score > maxscore {
					maxscore = score.score + score2.score
					maxwords = fmt.Sprintf("%s %d, %s %d", score.word, score.score, score2.word, score2.score)
					fmt.Println(maxwords, maxscore)
					continue
				}
			}
		}
	}

	// return

	maxscore = 0
	maxwords = ""

	for i, score := range scores {
		if score.score*3 <= maxscore {
			continue
		}

		f := score.letters
		if len(f) != 5 {
			continue
		}

		candidates := scores[i+1:]
		for j, score2 := range candidates {
			f2 := score2.letters
			if len(f2) != 5 {
				continue
			}

			letters := aoc.NewStringSet().AddSet(f).AddSet(f2)
			if len(letters) != 10 {
				continue
			}

			candidates2 := candidates[j+1:]
			for _, score3 := range candidates2 {
				if score.score+score2.score+score3.score >= maxscore {
					f3 := aoc.NewStringSet(score3.letters.Values()...)
					if len(f3) != 5 {
						continue
					}

					f3.AddSet(letters)
					if len(f3) == 15 {
						maxscore = score.score + score2.score + score3.score
						maxwords = fmt.Sprintf("%s %d, %s %d, %s %d", score.word, score.score, score2.word, score2.score, score3.word, score3.score)
						fmt.Println(maxwords, maxscore)
						continue
					}
				}
			}
		}
	}

	// fmt.Println(maxwords, maxscore)

	// fmt.Println(letterCounts)
}

func main() {
	process(aoc.Strings("words"))
}
