package aoc

import (
	"cmp"
	"image"
	"log"
	"math"
	"slices"
)

// Integer is a constraint for integer types.
// Note: cmp.Ordered from stdlib is more general (includes floats/strings).
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Sum returns total of all values.
func Sum[T Integer](in ...T) (sum T) {
	for _, i := range in {
		sum += i
	}
	return
}

// Product multiplies all values together.
func Product[T Integer](in ...T) (p T) {
	p = 1
	for _, i := range in {
		p = p * i
	}
	return
}

// Min returns smallest value. Panics if empty.
// Note: For slices, use slices.Min from stdlib directly.
func Min[T cmp.Ordered](in ...T) T {
	if len(in) == 0 {
		log.Fatalln("no values in array")
	}
	return slices.Min(in)
}

// Max returns largest value. Panics if empty.
// Note: For slices, use slices.Max from stdlib directly.
func Max[T cmp.Ordered](in ...T) T {
	if len(in) == 0 {
		log.Fatalln("no values in array")
	}
	return slices.Max(in)
}

// GCD returns the greatest common divisor (GCD) via Euclidean algorithm.
func GCD[T Integer](a, b T) T {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns Least Common Multiple (LCM) via GCD.
func LCM[T Integer](nums ...T) (lcm T) {
	if len(nums) == 0 {
		return
	}

	lcm = nums[0]
	for _, num := range nums[1:] {
		lcm = lcm * num / GCD(lcm, num)
	}
	return
}

// Signed is a constraint for signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Abs returns absolute value.
func Abs[T Signed](x T) T {
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

// Permutations returns all the permutations, written with help from ChatGPT
func Permutations[T any](in []T) [][]T {
	if len(in) == 0 {
		return [][]T{}
	}

	if len(in) == 1 {
		return [][]T{in}
	}

	var out [][]T
	for i, x := range in {
		rest := make([]T, len(in)-1)
		copy(rest, in[:i])
		copy(rest[i:], in[i+1:])
		for _, p := range Permutations(rest) {
			out = append(out, append([]T{x}, p...))
		}
	}
	return out
}

// PermutationsString returns all the permutations of strings.
// Deprecated: Use Permutations[string] instead.
func PermutationsString(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
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

// Distance returns the Euclidean distance between 2 points
func Distance(p, q image.Point) float64 {
	dx, dy := p.X-q.X, p.Y-q.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func ComparePoints(a, b image.Point) int {
	if a.Y < b.Y {
		return -1
	}
	if a.Y > b.Y {
		return 1
	}

	if a.X < b.X {
		return -1
	}
	if a.X > b.X {
		return 1
	}
	return 0
}

// LessThan returns true if a point is less than another point
func LessThan(a, b image.Point) bool { return ComparePoints(a, b) == -1 }
