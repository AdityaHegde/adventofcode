package day8

import (
	"AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string) int {
	points := parseAntenna(lines)
	visisted := utils.NewTwoDGrid[bool]()
	res := 0

	withinBounds := func(x, y int) bool {
		return x >= 0 && x < len(lines[0]) && y >= 0 && y < len(lines)
	}

	for _, samePoints := range points {
		i := 0
		for i < len(samePoints)-1 {
			j := i + 1
			for j < len(samePoints) {
				p1, p2 := getOtherPoints(samePoints[i], samePoints[j])
				if !visisted.Has(p1.x, p1.y) && withinBounds(p1.x, p1.y) {
					res += 1
					visisted.Set(p1.x, p1.y, true)
				}
				if !visisted.Has(p2.x, p2.y) && withinBounds(p2.x, p2.y) {
					res += 1
					visisted.Set(p2.x, p2.y, true)
				}
				j += 1
			}
			i += 1
		}
	}

	return res
}

func partTwo(lines []string) int {
	points := parseAntenna(lines)
	visisted := utils.NewTwoDGrid[bool]()
	res := 0

	withinBounds := func(x, y int) bool {
		return x >= 0 && x < len(lines[0]) && y >= 0 && y < len(lines)
	}

	for _, samePoints := range points {
		i := 0
		for i < len(samePoints)-1 {
			j := i + 1
			for j < len(samePoints) {
				p1 := samePoints[i]
				p2 := samePoints[j]
				dx := p2.x - p1.x
				dy := p2.y - p1.y

				po := point{p1.x, p1.y}
				for withinBounds(po.x, po.y) {
					if !visisted.Has(po.x, po.y) {
						res += 1
						visisted.Set(po.x, po.y, true)
					}
					po = point{po.x - dx, po.y - dy}
				}

				po = point{p2.x, p2.y}
				for withinBounds(po.x, po.y) {
					if !visisted.Has(po.x, po.y) {
						res += 1
						visisted.Set(po.x, po.y, true)
					}
					po = point{po.x + dx, po.y + dy}
				}

				j += 1
			}
			i += 1
		}
	}

	return res
}

func parseAntenna(lines []string) [][]point {
	samePoints := make(map[int32]int)
	points := make([][]point, 0)

	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				continue
			}
			idx, ok := samePoints[c]
			if !ok {
				idx = len(points)
				points = append(points, []point{})
				samePoints[c] = idx
			}

			points[idx] = append(points[idx], point{x, y})
		}
	}

	return points
}

type point struct {
	x, y int
}

func getOtherPoints(p1, p2 point) (point, point) {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return point{p1.x - dx, p1.y - dy}, point{p2.x + dx, p2.y + dy}
}
