package day24

import (
	"strings"

	"AdityaHegde/adventofcode/go/utils"
)

type point struct {
	x, y, z int
}

func parsePoint(str string) point {
	parts := strings.Split(str, ", ")
	return point{
		x: utils.Int(parts[0]),
		y: utils.Int(parts[1]),
		z: utils.Int(parts[2]),
	}
}

type hail struct {
	pos point
	vel point
}

func parseHail(line string) *hail {
	parts := strings.Split(line, " @ ")
	return &hail{
		pos: parsePoint(parts[0]),
		vel: parsePoint(parts[1]),
	}
}
