package day5

import (
	"math"
	"slices"
	"strings"

	"AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string) int {
	order, l := parseOrder(lines)
	res := 0

	for l < len(lines) {
		pageNums := utils.ParseInts(lines[l])
		l += 1

		if hasCorrectOrder(order, pageNums) {
			res += pageNums[int(math.Floor(float64(len(pageNums)/2)))]
		}

	}

	return res
}

func partTwo(lines []string) int {
	order, l := parseOrder(lines)
	res := 0

	for l < len(lines) {
		pageNums := utils.ParseInts(lines[l])
		l += 1

		if hasCorrectOrder(order, pageNums) {
			continue
		}

		slices.SortFunc(pageNums, func(left, right int) int {
			if order.Has(left, right) {
				return -1
			}
			return 1
		})

		res += pageNums[int(math.Floor(float64(len(pageNums)/2)))]
	}

	return res
}

func parseOrder(lines []string) (*utils.TwoDGrid[bool], int) {
	order := utils.NewTwoDGrid[bool]()
	l := 0

	for l < len(lines) {
		if strings.TrimSpace(lines[l]) == "" {
			break
		}

		orderParts := strings.Split(lines[l], "|")
		left := utils.Int(orderParts[0])
		right := utils.Int(orderParts[1])

		order.Set(right, left, true)

		l += 1
	}

	return order, l + 1
}

func hasCorrectOrder(order *utils.TwoDGrid[bool], pageNums []int) bool {
	i := 0
	for i < len(pageNums)-1 {
		j := i + 1
		for j < len(pageNums) {
			left := pageNums[i]
			right := pageNums[j]
			if order.Has(left, right) {
				return false
			}
			j += 1
		}
		i += 1
	}

	return true
}
