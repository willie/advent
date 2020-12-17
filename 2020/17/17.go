package main

import (
	"fmt"

	"github.com/willie/advent/aoc"
)

type world map[int]map[int]map[int]string

func (w world) Set(x, y, z int, s string) {
	// does y exist
	if _, ok := w[x]; !ok {
		w[x] = map[int]map[int]string{}
	}

	// does z exist
	if _, ok := w[x][y]; !ok {
		w[x][y] = map[int]string{}
	}

	// set it
	w[x][y][z] = s
}

func (w world) Get(x, y, z int) (s string) {
	s = inactive

	// x
	if _, ok := w[x]; !ok {
		return
	}

	// y
	if _, ok := w[x][y]; !ok {
		return
	}

	// z
	if _, ok := w[x][y][z]; !ok {
		return
	}

	return w[x][y][z]
}

// Iterate the grid, return false if f returns false
func (w world) Iterate(f func(x, y, z int, s string) bool) bool {
	for x, wy := range w {
		for y, wz := range wy {
			for z, s := range wz {
				if !f(x, y, z, s) {
					return false
				}
			}
		}
	}
	return true
}

func (w world) Dimensions() (minX, maxX, minY, maxY, minZ, maxZ int) {
	w.Iterate(func(x, y, z int, s string) bool {
		if x < minX {
			minX = x
		} else if x > maxX {
			maxX = x
		}

		if y < minY {
			minY = y
		} else if y > maxY {
			maxY = y
		}

		if z < minZ {
			minZ = z
		} else if z > maxZ {
			maxZ = z
		}
		return true
	})
	return
}

func (w world) Print() {
	minX, maxX, minY, maxY, minZ, maxZ := w.Dimensions()
	// fmt.Println(minX, maxX, minY, maxY, minZ, maxZ)

	for z := minZ; z <= maxZ; z++ {
		fmt.Println("z=", z)
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				fmt.Print(w.Get(x, y, z))
			}
			println()
		}
		println()
	}
	println()
}

const (
	inactive = "."
	active   = "#"
)

func part1(in []string) (result [2]int) {
	w := world{}

	for y, line := range in {
		for x, c := range line {
			w.Set(x, y, 0, string(c))
		}
	}

	w.Print()

	// cycles
	for i := 0; i < 6; i++ {
		nw := world{}

		// consider neighbors
		// w.Iterate(func(x, y, z int, s string) bool {
		minX, maxX, minY, maxY, minZ, maxZ := w.Dimensions()
		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for z := minZ - 1; z <= maxZ+1; z++ {
					s := w.Get(x, y, z)

					activeNeighbors := 0

					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							for dz := -1; dz <= 1; dz++ {
								if dx == 0 && dy == 0 && dz == 0 {
									continue
								}

								if w.Get(x+dx, y+dy, z+dz) == active {
									activeNeighbors++
								}
							}
						}
					}

					state := inactive

					if s == active && (activeNeighbors == 2 || activeNeighbors == 3) {
						state = active
					}

					if s == inactive && (activeNeighbors == 3) {
						state = active
					}

					if state == active {
						nw.Set(x, y, z, state)
					}
				}
			}
		}
		// 	return true
		// })

		// count active
		a := 0
		w.Iterate(func(x, y, z int, s string) bool {
			if s == active {
				a++
			}
			return true
		})
		// fmt.Println(a)

		w = nw
		fmt.Println("After", i+1, "cycles:")
		println()
		w.Print()
	}

	// count active
	w.Iterate(func(x, y, z int, s string) bool {
		if s == active {
			result[0]++
		}
		return true
	})

	return
}

type space map[int]map[int]map[int]map[int]string

func (ww space) Set(x, y, z, w int, s string) {
	// does y exist
	if _, ok := ww[x]; !ok {
		ww[x] = map[int]map[int]map[int]string{}
	}

	// does z exist
	if _, ok := ww[x][y]; !ok {
		ww[x][y] = map[int]map[int]string{}
	}

	// does w exist
	if _, ok := ww[x][y][z]; !ok {
		ww[x][y][z] = map[int]string{}
	}

	// set it
	ww[x][y][z][w] = s
}

func (sp space) Get(x, y, z, w int) (s string) {
	s = inactive

	if sy, ok := sp[x]; ok {
		if sz, ok := sy[y]; ok {
			if sw, ok := sz[z]; ok {
				if v, ok := sw[w]; ok {
					return v
				}
			}
		}
	}

	return
}

func (ww space) GetOld(x, y, z, w int) (s string) {
	s = inactive

	// x
	if _, ok := ww[x]; !ok {
		return
	}

	// y
	if _, ok := ww[x][y]; !ok {
		return
	}

	// z
	if _, ok := ww[x][y][z]; !ok {
		return
	}

	// w
	if _, ok := ww[x][y][z][w]; !ok {
		return
	}

	return ww[x][y][z][w]
}

func (sp space) iterate(f func(x, y, z, w int, s string) bool) bool {
	for x, wy := range sp {
		for y, wz := range wy {
			for z, ww := range wz {
				for w, s := range ww {
					if !f(x, y, z, w, s) {
						return false
					}
				}
			}
		}
	}
	return true
}

func (sp space) Dimensions() (minX, maxX, minY, maxY, minZ, maxZ, minW, maxW int) {
	sp.iterate(func(x, y, z, w int, s string) bool {
		if x < minX {
			minX = x
		} else if x > maxX {
			maxX = x
		}

		if y < minY {
			minY = y
		} else if y > maxY {
			maxY = y
		}

		if z < minZ {
			minZ = z
		} else if z > maxZ {
			maxZ = z
		}

		if w < minW {
			minW = w
		} else if w > maxW {
			maxW = w
		}

		return true
	})
	return
}

func (sp space) print() {
	minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := sp.Dimensions()
	// fmt.Println(minX, maxX, minY, maxY, minZ, maxZ)

	for w := minW; w <= maxW; w++ {
		for z := minZ; z <= maxZ; z++ {
			fmt.Println("z=", z, "w=", w)
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					fmt.Print(sp.Get(x, y, z, w))
				}
				println()
			}
			println()
		}
	}
	println()
}

func part2(in []string) (result int) {
	sp := space{}

	for y, line := range in {
		for x, c := range line {
			sp.Set(x, y, 0, 0, string(c))
		}
	}

	sp.print()

	// cycles
	for i := 0; i < 6; i++ {
		nw := space{}

		// consider neighbors
		// w.Iterate(func(x, y, z int, s string) bool {
		minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := sp.Dimensions()
		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for z := minZ - 1; z <= maxZ+1; z++ {
					for w := minW - 1; w <= maxW+1; w++ {
						s := sp.Get(x, y, z, w)

						activeNeighbors := 0

						for dx := -1; dx <= 1; dx++ {
							for dy := -1; dy <= 1; dy++ {
								for dz := -1; dz <= 1; dz++ {
									for dw := -1; dw <= 1; dw++ {
										if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
											continue
										}

										if sp.Get(x+dx, y+dy, z+dz, w+dw) == active {
											activeNeighbors++
										}
									}
								}
							}
						}

						state := inactive

						if s == active && (activeNeighbors == 2 || activeNeighbors == 3) {
							state = active
						}

						if s == inactive && (activeNeighbors == 3) {
							state = active
						}

						if state == active {
							nw.Set(x, y, z, w, state)
						}
					}
				}
			}
		}

		// 	return true
		// })

		// count active
		a := 0
		sp.iterate(func(x, y, z, w int, s string) bool {
			if s == active {
				a++
			}
			return true
		})
		// fmt.Println(a)

		sp = nw
		fmt.Println("After", i+1, "cycles:")
		println()
		sp.print()
	}

	// count active
	sp.iterate(func(x, y, z, w int, s string) bool {
		if s == active {
			result++
		}
		return true
	})

	return
}

const day = "https://adventofcode.com/2020/day/17"

func main() {
	println(day)
	aoc.Input(day)

	fmt.Println("test", part1(aoc.Strings("test")), 112)
	fmt.Println("run", part1(aoc.Strings(day)))

	fmt.Println("test2", part2(aoc.Strings("test")), 848)
	fmt.Println("run", part2(aoc.Strings(day)))
}
