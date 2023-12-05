package day5

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"sync"

	"AdityaHegde/adventofcode/go/utils"
)

var (
	SeedsRegex   = regexp.MustCompile(`seeds: (.*)$`)
	NumbersRegex = regexp.MustCompile(`(\d+)`)
)

type numMap struct {
	from, to, rng int
}

func parse(input []string) ([]int, [][]numMap) {
	seedsMatch := SeedsRegex.FindStringSubmatch(input[0])
	if len(seedsMatch) == 0 {
		panic(fmt.Sprintf("failed to parse: %s", input[0]))
	}
	seedNumsMatch := NumbersRegex.FindAllStringSubmatch(seedsMatch[1], -1)
	seeds := make([]int, len(seedNumsMatch))
	for i, seedMatch := range seedNumsMatch {
		seeds[i] = utils.Int(seedMatch[1])
	}

	seedsMap := make([][]numMap, 0)
	for i := 3; i < len(input); i += 2 {
		numMaps := make([]numMap, 0)

		for ; i < len(input) && strings.TrimSpace(input[i]) != ""; i++ {
			numsMatch := NumbersRegex.FindAllStringSubmatch(input[i], -1)
			if len(numsMatch) == 0 {
				panic(fmt.Sprintf("failed to parse: %s", input[i]))
			}
			numMaps = append(numMaps, numMap{
				utils.Int(numsMatch[1][1]),
				utils.Int(numsMatch[0][1]),
				utils.Int(numsMatch[2][1]),
			})
		}

		seedsMap = append(seedsMap, numMaps)
	}

	return seeds, seedsMap
}

func mapSeed(seed int, seedsMap [][]numMap) int {
	for _, numMaps := range seedsMap {
		for _, nm := range numMaps {
			if seed >= nm.from && seed <= nm.from+nm.rng {
				seed = seed - nm.from + nm.to
				break
			}
		}
	}
	return seed
}

func partOne(input []string) int {
	seeds, seedsMap := parse(input)

	res := math.MaxInt64
	for _, seed := range seeds {
		seed = mapSeed(seed, seedsMap)
		if seed < res {
			res = seed
		}
	}

	return res
}

func partTwo(input []string) int {
	seeds, seedsMap := parse(input)

	lk := sync.Mutex{}
	res := math.MaxInt64
	var wg sync.WaitGroup

	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		size := seeds[i+1]

		wg.Add(1)
		go func() {
			defer wg.Done()

			m := math.MaxInt64
			for s := 0; s < size; s++ {
				sd := mapSeed(seed+s, seedsMap)
				if sd < m {
					m = sd
				}
			}

			lk.Lock()
			defer lk.Unlock()
			if m < res {
				res = m
			}
		}()
	}

	wg.Wait()

	return res
}
