package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	maxCalories := 0
	maxElf := 0

	// Keep track of the current elf and their total calories
	currentElf := 0
	currentTotal := 0

	// Read in each line from the input
	for scanner.Scan() {
		line := scanner.Text()

		// If the line is empty, we have reached the end of the current elf's inventory
		if line == "" {
			// If the current elf has more calories than the current max, update the max
			if currentTotal > maxCalories {
				maxCalories = currentTotal
				maxElf = currentElf
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
	fmt.Printf("Elf %d has the most calories with a total of %d\n", maxElf+1, maxCalories)
}
