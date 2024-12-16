package utils

import "math"

type Point struct {
	X, Y int
}

var Directions = []*Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type Vector struct {
	X, Y, Dir int
}

func (v *Vector) RotateLeft() {
	v.Dir = rotateLeft(v.Dir)
}

func (v *Vector) RotateRight() {
	v.Dir = rotateRight(v.Dir)
}

func (v *Vector) Rotations(dirIdx int) int {
	rotations := int(math.Abs(float64(v.Dir - dirIdx)))
	if rotations > 2 {
		// diff between 0 & 3 is just 1
		rotations = 1
	}
	return rotations
}

func (v *Vector) NextPoint() (int, int) {
	return v.X + Directions[v.Dir].X, v.Y + Directions[v.Dir].Y
}

func (v *Vector) RightPoint() (int, int) {
	rdir := rotateRight(v.Dir)
	return v.X + Directions[rdir].X, v.Y + Directions[rdir].Y
}

func (v *Vector) Move() {
	v.X, v.Y = v.NextPoint()
}

func rotateLeft(dir int) int {
	return (len(Directions) + dir - 1) % len(Directions)
}

func rotateRight(dir int) int {
	return (dir + 1) % len(Directions)
}
