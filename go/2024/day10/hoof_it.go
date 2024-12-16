package day10

import (
	"AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string) int {
	res := 0
	for y, line := range lines {
		for x, c := range line {
			if c != '0' {
				continue
			}
			res += search(x, y, 0, lines, utils.NewTwoDGrid[bool]())
		}
	}
	return res
}

func partTwo(lines []string) int {
	res := 0
	for y, line := range lines {
		for x, c := range line {
			if c != '0' {
				continue
			}
			res += search(x, y, 0, lines, nil)
		}
	}
	return res
}

const ZeroCharCode = 48

func search(x, y, curHt int, htMap []string, visited *utils.TwoDGrid[bool]) int {
	if x < 0 || x >= len(htMap[0]) || y < 0 || y >= len(htMap) {
		return 0
	}
	ht := int(htMap[y][x] - ZeroCharCode)
	if ht != curHt {
		return 0
	}
	if curHt == 9 {
		if visited != nil {
			if visited.Has(x, y) {
				return 0
			}
			visited.Set(x, y, true)
		}
		return 1
	}

	return search(x+1, y, curHt+1, htMap, visited) +
		search(x, y+1, curHt+1, htMap, visited) +
		search(x-1, y, curHt+1, htMap, visited) +
		search(x, y-1, curHt+1, htMap, visited)
}
