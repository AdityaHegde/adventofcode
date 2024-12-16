package day14

import (
	"regexp"

	"AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string, width, height, iterations int) int {
	robots := parseRobots(lines)
	i := 0
	for i < iterations {
		for _, r := range robots {
			r.move(width, height)
		}
		i += 1
	}

	quadrants := []int{0, 0, 0, 0}
	mw := width / 2
	mh := height / 2
	for _, r := range robots {
		if r.x == mw || r.y == mh {
			continue
		}
		x := 0
		if r.x > mw {
			x = 1
		}
		y := 0
		if r.y > mh {
			y = 1
		}
		qi := x + y*2
		quadrants[qi] += 1
	}

	res := 1
	for _, q := range quadrants {
		res *= q
	}
	return res
}

func parseRobots(lines []string) []*robot {
	robots := make([]*robot, len(lines))
	for i, line := range lines {
		robots[i] = newRobot(line)
	}
	return robots
}

type robot struct {
	x, y, dx, dy int
}

var robotRegex = regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

func newRobot(line string) *robot {
	r := robot{}
	matches := robotRegex.FindStringSubmatch(line)
	r.x = utils.Int(matches[1])
	r.y = utils.Int(matches[2])
	r.dx = utils.Int(matches[3])
	r.dy = utils.Int(matches[4])
	return &r
}

func (r *robot) move(width, height int) {
	r.x = (width + (r.x + r.dx)) % width
	r.y = (height + (r.y + r.dy)) % height
}
