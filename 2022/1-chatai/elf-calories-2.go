package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// coded by instructing ChatGPT

type Elf struct {
	ID       int
	Calories int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// Use a slice of Elf to keep track of the top elves
	topElves := make([]Elf, 0, 3)

	// Keep track of the current elf and their total calories
	currentElf := 0
	currentTotal := 0

	topElves = append(topElves, Elf{}) // I added this line to fix a range issue

	// Read in each line from the input
	for scanner.Scan() {
		line := scanner.Text()

		// If the line is empty, we have reached the end of the current elf's inventory
		if line == "" {
			// Add the current elf to the top elves if they have more calories than the current minimum
			if currentTotal > topElves[len(topElves)-1].Calories {
				topElves = append(topElves, Elf{ID: currentElf, Calories: currentTotal})

				// Sort the top elves in descending order of calories
				sort.Slice(topElves, func(i, j int) bool {
					return topElves[i].Calories > topElves[j].Calories
				})

				// If we have more than 3 top elves, remove the lowest calorie elf
				if len(topElves) > 3 {
					topElves = topElves[:3]
				}
			}

			// Reset the current elf and total
			currentElf++
			currentTotal = 0
		} else {
			// Parse the calories from the current line and add it to the current elf's total
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			currentTotal += calories
		}
	}

	// Check if there was an error while reading the input
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Print the results
	totalCalories := 0
	for i, elf := range topElves {
		totalCalories += elf.Calories
		fmt.Printf("Elf %d is ranked #%d with %d calories\n", elf.ID+1, i+1, elf.Calories)
	}
	fmt.Printf("The top three elves have a total of %d calories\n", totalCalories)
}
