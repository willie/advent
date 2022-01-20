package main

import (
	"fmt"
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
	word  string
	score int
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

	wordScore := make(map[string]int)
	for _, word := range words {
		for i, c := range strings.Split(word, "") {
			wordScore[word] += letterCounts[c][i]
		}
	}

	// print out the words in score order
	scores := []wordscore{}
	for word, score := range wordScore {
		scores = append(scores, wordscore{word, score})
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	// fmt.Println(scores)

	maxscore := 0
	maxwords := ""
	for i, score := range scores {
		f := aoc.NewStringSet(strings.Split(score.word, "")...)

		for _, score2 := range scores[i+1:] {
			f2 := aoc.NewStringSet(strings.Split(score2.word, "")...)

			if len(f.Values()) != 5 || len(f2.Values()) != 5 {
				continue
			}

			if len(f.Subtract(f2).Values()) == 5 {
				if score.score+score2.score > maxscore {
					maxscore = score.score + score2.score
					maxwords = score.word + " " + score2.word
					fmt.Println(maxwords, maxscore)
					continue
				}
			}
		}
	}

	for i, score := range scores {
		f := aoc.NewStringSet(strings.Split(score.word, "")...)
		if len(f.Values()) != 5 {
			continue
		}

		candidates := scores[i+1:]
		for j, score2 := range candidates {
			f2 := aoc.NewStringSet(strings.Split(score2.word, "")...)
			if len(f2.Values()) != 5 {
				continue
			}

			letters := aoc.NewStringSet().AddSet(f).AddSet(f2)
			if len(letters) != 10 {
				continue
			}

			candidates2 := candidates[j+1:]
			for _, score3 := range candidates2 {
				f3 := aoc.NewStringSet(strings.Split(score3.word, "")...)
				if len(f3.Values()) != 5 {
					continue
				}

				f3.AddSet(letters)
				if len(f3.Values()) == 15 {
					if score.score+score2.score+score3.score > maxscore {
						maxscore = score.score + score2.score + score3.score
						maxwords = score.word + " " + score2.word + " " + score3.word
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
