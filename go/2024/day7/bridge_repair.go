package day7

import (
	"math"
	"strings"

	"AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string) int {
	return checkTotals(lines, false)
}

func partTwo(lines []string) int {
	return checkTotals(lines, true)
}

func checkTotals(lines []string, concat bool) int {
	eqns := parseEquations(lines)
	res := 0

	for _, e := range eqns {
		totals := e.possibleTotals(concat)
		for _, total := range totals {
			if e.total == total {
				res += e.total
				break
			}
		}
	}

	return res
}

func parseEquations(lines []string) []*eqn {
	eqns := make([]*eqn, len(lines))

	for i, line := range lines {
		e := eqn{}
		eqnParts := strings.Split(line, ": ")
		e.total = utils.Int(eqnParts[0])
		e.nums = utils.ParseInts(eqnParts[1])
		eqns[i] = &e
	}

	return eqns
}

type eqn struct {
	total int
	nums  []int
}

func (e *eqn) possibleTotals(concat bool) []int {
	totals := []int{e.nums[0]}
	digits := make([]int, len(e.nums))
	if concat {
		for i, num := range e.nums {
			digits[i] = int(math.Pow10(int(math.Log10(float64(num))) + 1))
		}
	}

	i := 1
	for i < len(e.nums) {
		newTotals := make([]int, 0)
		for _, total := range totals {
			newTotals = append(newTotals, total+e.nums[i])
			newTotals = append(newTotals, total*e.nums[i])
			if concat {
				newTotals = append(newTotals, total*digits[i]+e.nums[i])
			}
		}
		totals = newTotals
		i += 1
	}

	return totals
}
