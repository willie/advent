package aoc

// Series returns a slice of integers from low to high (inclusive).
func Series(low, high int) []int {
	if high < low {
		return nil
	}
	series := make([]int, high-low+1)
	for i := range series {
		series[i] = low + i
	}
	return series
}
