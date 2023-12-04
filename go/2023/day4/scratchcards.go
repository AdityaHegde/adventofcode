package day4

import (
	"fmt"
	"math"
	"regexp"

	"AdityaHegde/adventofcode/go/utils"
)

var (
	ScratchCardsRegex = regexp.MustCompile(`^Card\s*\d+:(.*)\|(.*)$`)
	NumbersRegex      = regexp.MustCompile(`(\d+)`)
)

func parse(line string) int {
	numberMatches := ScratchCardsRegex.FindStringSubmatch(line)
	if len(numberMatches) == 0 {
		panic(fmt.Sprintf("Failed to parse: %s", line))
	}

	winnerMatches := NumbersRegex.FindAllStringSubmatch(numberMatches[2], -1)
	winners := map[int]bool{}
	for _, w := range winnerMatches {
		winners[utils.Int(w[1])] = true
	}

	scratchMatches := NumbersRegex.FindAllStringSubmatch(numberMatches[1], -1)
	res := 0
	for _, s := range scratchMatches {
		if !winners[utils.Int(s[1])] {
			continue
		}
		res++
	}
	return res
}

func partOne(input []string) int {
	res := 0

	for _, line := range input {
		res += int(math.Pow(2.0, float64(parse(line)-1)))
	}

	return res
}

func partTwo(input []string) int {
	il := len(input)
	cards := make([]int, il)
	for i, line := range input {
		cards[i]++
		wins := parse(line)
		for j := 1; j <= wins && j < il; j++ {
			cards[j+i] += cards[i]
		}
	}

	res := 0
	for _, c := range cards {
		res += c
	}
	return res
}
