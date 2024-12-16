package day11

import (
	"AdityaHegde/adventofcode/go/utils"
)

func partOne(input string, steps int) int {
	nums := utils.ParseInts(input)

	cache := utils.NewTwoDGrid[[]int]()
	res := 0
	for _, num := range nums {
		splits := changeNumber(num, steps, cache)
		res += len(splits)
	}

	return res
}

func changeNumber(num, times int, cache *utils.TwoDGrid[[]int]) []int {
	if cache.Has(num, times) {
		return cache.Grid[num][times]
	}
	if times == 0 {
		return []int{num}
	}

	nums := make([]int, 0)

	if num == 0 {
		nums = changeNumber(1, times-1, cache)
	} else {
		digits := utils.Digits(num)
		if len(digits)%2 == 0 {
			left, right := split(digits)
			nums = append(nums, changeNumber(left, times-1, cache)...)
			nums = append(nums, changeNumber(right, times-1, cache)...)
		} else {
			nums = changeNumber(num*2024, times-1, cache)
		}
	}

	cache.Set(num, times, nums)
	return nums
}

func split(digits []int) (int, int) {
	mid := len(digits) / 2
	i := mid - 1
	left := 0
	right := 0

	for i >= 0 {
		left = left*10 + digits[mid+i]
		right = right*10 + digits[i]
		i -= 1
	}

	return left, right
}
