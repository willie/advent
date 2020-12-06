package aoc

import "log"

// Ints is []int
type Ints []int

// Sum returns total
func (ints Ints) Sum() (sum int) {
	for _, i := range ints {
		sum += i
	}

	return
}

// Product multiplies all the numbers together
func (ints Ints) Product() (p int) {
	p = 1
	for _, i := range ints {
		p = p * i
	}
	return
}

// Min returns smallest value
func (ints Ints) Min() (min int) {
	if len(ints) == 0 {
		log.Fatalln("no values in array")
	}

	min = ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < min {
			min = ints[i]
		}
	}

	return
}

// Max returns largest value
func (ints Ints) Max() (max int) {
	if len(ints) == 0 {
		log.Fatalln("no values in array")
	}

	max = ints[0]
	for i := 1; i < len(ints); i++ {
		if max < ints[i] {
			max = ints[i]
		}
	}

	return
}

// Series returns array of low including high
func Series(low, high int) (series Ints) {
	series = make(Ints, (high-low)+1)

	x := 0
	for i := low; i <= high; i++ {
		series[x] = i
		x++
	}
	return
}
