package aoc

import (
	"image"
	"log"
	"math"
)

// Sum returns total
func Sum(in ...int) (sum int) {
	for _, i := range in {
		sum += i
	}

	return
}

// Sum returns total
func Sum64(in ...int64) (sum int64) {
	for _, i := range in {
		sum += i
	}

	return
}

// Product multiplies all the numbers together
func Product(ints ...int) (p int) {
	p = 1
	for _, i := range ints {
		p = p * i
	}
	return
}

// Min returns smallest value
func Min(in ...int) (min int) {
	if len(in) == 0 {
		log.Fatalln("no values in array")
	}

	min = in[0]
	for i := 1; i < len(in); i++ {
		if in[i] < min {
			min = in[i]
		}
	}

	return
}

// Max returns largest value
func Max(in ...int) (max int) {
	if len(in) == 0 {
		log.Fatalln("no values in array")
	}

	max = in[0]
	for i := 1; i < len(in); i++ {
		if max < in[i] {
			max = in[i]
		}
	}

	return
}

// GCD returns the greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM returns Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Abs return absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ManhattanDistance returns the "taxicab distance" between 2 points.
// https://computervision.fandom.com/wiki/Manhattan_distance
func ManhattanDistance(x, y, x1, y1 int) (distance int) {
	return Abs(x-x1) + Abs(y-y1)
}

// Permutations returns all the permutations
func Permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

// AngleDistance returns the angle and distance between 2 points
func AngleDistance(a, b image.Point) (angle, distance float64) {
	n := float64(a.Y - b.Y)
	d := float64(a.X - b.X)

	angle = math.Atan2(d, n) * (-180.0 / math.Pi)
	if angle != math.Abs(angle) {
		angle += 360
	}

	distance = math.Sqrt((n * n) + (d * d))

	return angle, distance
}
