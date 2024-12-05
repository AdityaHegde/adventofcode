package day1

import (
	"math"
	"slices"
	"strings"

	"AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string) int {
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left = append(left, utils.Int(parts[0]))
		right = append(right, utils.Int(parts[1]))
	}
	slices.Sort(left)
	slices.Sort(right)

	dist := 0
	for i := range left {
		dist += int(math.Abs(float64(left[i] - right[i])))
	}
	return dist
}

func partTwo(lines []string) int {
	left := make([]int, 0)
	right := make(map[int]int)
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left = append(left, utils.Int(parts[0]))
		r := utils.Int(parts[1])
		right[r] = right[r] + 1
	}

	dist := 0
	for _, l := range left {
		dist += l * right[l]
	}
	return dist
}
