package main

import (
	"fmt"
	"strings"

	"github.com/willie/advent/aoc"
)

func parseTickets(in string) (tickets []aoc.Ints) {
	for _, tix := range strings.Split(in, "\n")[1:] {
		if tix == "" {
			continue
		}

		ticket := aoc.Ints{}
		for _, v := range strings.Split(tix, ",") {
			if strings.Contains(v, ":") || (v == "") {
				continue
			}

			value := aoc.AtoI(v)
			ticket = append(ticket, value)
		}
		tickets = append(tickets, ticket)
	}
	return
}

func part1(in string) (result [2]int) {
	parts := strings.Split(in, "\n\n")

	fieldNames := aoc.StringSet{}
	validation := map[int][]string{}
	// parse fields, ranges
	for _, p := range strings.Split(parts[0], "\n") {
		row := strings.Split(p, ": ")

		field, ranges := row[0], aoc.IntSet{}
		fieldNames.Add(field)

		for _, valid := range strings.Split(row[1], " or ") {
			r := strings.Split(valid, "-")
			ranges.AddMany(aoc.Series(aoc.AtoI(r[0]), aoc.AtoI(r[1])))

			for _, i := range aoc.Series(aoc.AtoI(r[0]), aoc.AtoI(r[1])) {
				validation[i] = append(validation[i], field)
			}
		}

	}
	// fmt.Println(validation)

	var first int

	nearbyTickets := parseTickets(parts[2])
	var validTickets []aoc.Ints

	// check invalid fields
	for _, ticket := range nearbyTickets {
		valid := true

		for _, value := range ticket {
			// check for invalid values
			if _, ok := validation[value]; !ok {
				first += value
				valid = false
			}
		}

		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	result[0] = first

	// part 2
	myTicket := parseTickets(parts[1])[0]
	fmt.Println("my", myTicket)

	// field number to field names to occurence count
	fieldSort := []map[string]int{}
	for range validTickets[0] {
		fieldSort = append(fieldSort, map[string]int{})
	}

	// get the counts
	allTickets := append(validTickets, myTicket)

	for _, ticket := range allTickets {
		for field, value := range ticket {
			for _, name := range validation[value] {
				fieldSort[field][name]++
			}
		}
	}

	fieldNameSorter := map[int]aoc.StringSet{}
	for i := range validTickets[0] {
		fieldNameSorter[i] = aoc.StringSet{}
	}

	for i, nameCount := range fieldSort {
		for name, count := range nameCount {
			if count == len(allTickets) {
				fieldNameSorter[i].Add(name)
			}
		}
	}
	fmt.Println("fieldNameSorter", fieldNameSorter, len(allTickets))

	fieldNameOrder := map[int]string{}
	namesFound := aoc.StringSet{}
	for len(fieldNameOrder) < len(fieldNameSorter) {
		for n, names := range fieldNameSorter {
			if len(names) == 1 {
				name := names.Values()[0]

				fieldNameOrder[n] = name

				namesFound.Add(name)
				// remove from others
				for _, names := range fieldNameSorter {
					names.Remove(name)
				}
			}
		}
	}

	fmt.Println("fieldSort", fieldSort)

	// for i, nameCount := range fieldSort {
	// 	var highName string
	// 	var highCount int

	// 	for name, count := range nameCount {
	// 		if !fieldNames.Contains(name) { // already removed
	// 			continue
	// 		}

	// 		if count > highCount {
	// 			highCount = count
	// 			highName = name
	// 		}
	// 	}

	// 	fieldNameOrder[i] = highName
	// 	fieldNames.Remove(highName)
	// }

	for _, name := range fieldNameOrder {
		fmt.Print(name, " ")
	}
	fmt.Println()

	second := 1
	for i, name := range fieldNameOrder {
		if strings.Contains(name, "departure") {
			second *= myTicket[i]
		}
	}

	result[1] = second

	return
}

const day = "https://adventofcode.com/2020/day/16"

func main() {
	println(day)
	aoc.Input(day)

	fmt.Println("test", part1(aoc.String("test")), 71)
	fmt.Println("test2", part1(aoc.String("test2")), 13)
	fmt.Println("run", part1(aoc.String(day)))
}
