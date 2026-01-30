package aoc

import "image"

// Cardinal directions (screen coordinates: Y increases downward)
var (
	Up    = image.Pt(0, -1)
	Down  = image.Pt(0, 1)
	Left  = image.Pt(-1, 0)
	Right = image.Pt(1, 0)

	// Alternate names
	North = Up
	South = Down
	West  = Left
	East  = Right

	// Single character direction names (common in AoC)
	DirN = Up
	DirS = Down
	DirW = Left
	DirE = Right
	DirU = Up
	DirD = Down
	DirL = Left
	DirR = Right
)

// Diagonal directions
var (
	UpLeft    = image.Pt(-1, -1)
	UpRight   = image.Pt(1, -1)
	DownLeft  = image.Pt(-1, 1)
	DownRight = image.Pt(1, 1)

	NorthWest = UpLeft
	NorthEast = UpRight
	SouthWest = DownLeft
	SouthEast = DownRight
)

// Direction slices for iteration
var (
	// FourWay contains the 4 cardinal directions
	FourWay = []image.Point{Up, Down, Left, Right}

	// EightWay contains all 8 directions (cardinal + diagonal)
	EightWay = []image.Point{
		Up, UpRight, Right, DownRight,
		Down, DownLeft, Left, UpLeft,
	}

	// Cardinals is an alias for FourWay
	Cardinals = FourWay

	// Diagonals contains only diagonal directions
	Diagonals = []image.Point{UpLeft, UpRight, DownLeft, DownRight}
)

// DirFromChar converts a direction character to a Point.
// Supports: U/D/L/R, N/S/E/W, ^/v/</>, and arrow symbols.
func DirFromChar(c rune) image.Point {
	switch c {
	case 'U', 'N', '^':
		return Up
	case 'D', 'S', 'v', 'V':
		return Down
	case 'L', 'W', '<':
		return Left
	case 'R', 'E', '>':
		return Right
	default:
		return image.Point{}
	}
}

// DirFromString converts a direction string to a Point.
func DirFromString(s string) image.Point {
	if len(s) == 0 {
		return image.Point{}
	}
	if len(s) == 1 {
		return DirFromChar(rune(s[0]))
	}
	switch s {
	case "up", "UP", "Up", "north", "NORTH", "North":
		return Up
	case "down", "DOWN", "Down", "south", "SOUTH", "South":
		return Down
	case "left", "LEFT", "Left", "west", "WEST", "West":
		return Left
	case "right", "RIGHT", "Right", "east", "EAST", "East":
		return Right
	default:
		return image.Point{}
	}
}

// TurnLeft rotates a direction 90 degrees counter-clockwise.
func TurnLeft(dir image.Point) image.Point {
	return image.Pt(dir.Y, -dir.X)
}

// TurnRight rotates a direction 90 degrees clockwise.
func TurnRight(dir image.Point) image.Point {
	return image.Pt(-dir.Y, dir.X)
}

// TurnAround rotates a direction 180 degrees.
func TurnAround(dir image.Point) image.Point {
	return image.Pt(-dir.X, -dir.Y)
}

// Neighbors4 returns the 4 cardinal neighbors of a point.
func Neighbors4(p image.Point) []image.Point {
	return []image.Point{
		p.Add(Up),
		p.Add(Down),
		p.Add(Left),
		p.Add(Right),
	}
}

// Neighbors8 returns all 8 neighbors of a point.
func Neighbors8(p image.Point) []image.Point {
	return []image.Point{
		p.Add(Up),
		p.Add(UpRight),
		p.Add(Right),
		p.Add(DownRight),
		p.Add(Down),
		p.Add(DownLeft),
		p.Add(Left),
		p.Add(UpLeft),
	}
}
