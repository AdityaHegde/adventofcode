package day2

import (
	"math"
	"strings"

	"AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string) int {
	count := 0
	for _, line := range lines {
		levels := strings.Split(line, " ")

		if checkLevels(levels, len(levels)) {
			count++
		}
	}
	return count
}

func partTwo(lines []string) int {
	count := 0
	for _, line := range lines {
		levels := strings.Split(line, " ")

		if checkLevels(levels, -1) {
			count++
		}
	}
	return count
}

func checkLevels(levels []string, skipIdx int) bool {
	prev := utils.Int(levels[0])
	prevSign := 1
	hasPrevSign := false
	startIdx := 1
	if skipIdx == 0 {
		prev = utils.Int(levels[1])
		startIdx = 2
	}

	for i, lvl := range levels[startIdx:] {
		ie := i + startIdx
		if ie == skipIdx {
			continue
		}

		l := utils.Int(lvl)
		sign := utils.Sign(prev - l)
		diff := int(math.Abs(float64(prev - l)))

		prev = l

		if diff < 1 || diff > 3 || (hasPrevSign && prevSign != sign) {
			if skipIdx != -1 {
				return false
			}
			if ie == len(levels)-1 {
				return true
			}
			if ie == 2 && checkLevels(levels, 0) {
				return true
			}

			return checkLevels(levels, ie) || checkLevels(levels, ie-1)
		}

		hasPrevSign = true
		prevSign = sign
	}

	return true
}
