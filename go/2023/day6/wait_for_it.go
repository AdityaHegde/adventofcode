package day6

import (
  "regexp"

  "AdityaHegde/adventofcode/go/utils"
)

type race struct {
	time   int
	record int
}

func (r *race) beat() int {
	beg := 1
	for ; beg < r.time; beg++ {
		if beg*(r.time-beg) > r.record {
			break
		}
	}
	end := r.time
	for ; end >= beg; end-- {
		if end*(r.time-end) > r.record {
			break
		}
	}
	return end - beg + 1
}

var (
	TimeRegex     = regexp.MustCompile(`Time:\s+(.*)$`)
	DistanceRegex = regexp.MustCompile(`Distance:\s+(.*)$`)
)

func parsePartOne(input []string) []race {
	timeMatches := TimeRegex.FindStringSubmatch(input[0])
	times := utils.ParseInts(timeMatches[1])
	distanceMatches := DistanceRegex.FindStringSubmatch(input[1])
	distances := utils.ParseInts(distanceMatches[1])

	races := make([]race, len(times))
	for i, time := range times {
		races[i] = race{
			time,
			distances[i],
		}
	}

	return races
}

func partOne(input []string) int {
	res := 1
	races := parsePartOne(input)
	for _, race := range races {
		res *= race.beat()
	}
	return res
}

func parsePartTwo(input []string) race {
	timeMatches := TimeRegex.FindStringSubmatch(input[0])
	timeNums := utils.NumbersRegex.FindAllStringSubmatch(timeMatches[1], -1)
	timeNum := ""
	for _, num := range timeNums {
		timeNum += num[1]
	}
	distanceMatches := DistanceRegex.FindStringSubmatch(input[1])
	distanceNums := utils.NumbersRegex.FindAllStringSubmatch(distanceMatches[1], -1)
	distanceNum := ""
	for _, num := range distanceNums {
		distanceNum += num[1]
	}

	return race{
		utils.Int(timeNum),
		utils.Int(distanceNum),
	}
}

func partTwo(input []string) int {
	race := parsePartTwo(input)
	return race.beat()
}
