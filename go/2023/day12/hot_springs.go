package day12

import (
	"regexp"
	"strings"
	"sync"

	"AdityaHegde/adventofcode/go/utils"
)

type record struct {
	rec           string
	damaged       []int
	possibilities int
}

var (
	recordRegex = regexp.MustCompile(`^([.?#]*) (.*)$`)
)

func parseRecord(line string) *record {
	recordMatches := recordRegex.FindStringSubmatch(line)
	damagedStrs := strings.Split(recordMatches[2], ",")
	damaged := make([]int, len(damagedStrs))
	for i, str := range damagedStrs {
		damaged[i] = utils.Int(str)
	}
	return &record{
		rec:           recordMatches[1],
		damaged:       damaged,
		possibilities: 0,
	}
}

func (r *record) calculate(recIdx, damIdx, brokenLen int, cache *utils.ThreeDGrid[int]) {
	if recIdx == len(r.rec) {
		if damIdx == len(r.damaged) ||
			(r.damaged[damIdx] == brokenLen && damIdx == len(r.damaged)-1) {
			r.possibilities++
		}
		return
	}
	if cache.Has(recIdx, damIdx, brokenLen) {
		r.possibilities += cache.Grid[recIdx][damIdx][brokenLen]
		return
	}

	prevPos := r.possibilities
	switch r.rec[recIdx] {
	case '?':
		// all damaged blocks used up
		if damIdx == len(r.damaged) {
			r.calculate(recIdx+1, damIdx, 0, cache)
			break
		}

		// already in broken streak
		if brokenLen > 0 {
			if brokenLen < r.damaged[damIdx] {
				// increment brokenLen if current streak is less than the record
				r.calculate(recIdx+1, damIdx, brokenLen+1, cache)
			} else if brokenLen == r.damaged[damIdx] {
				// increment damaged streak index and reset brokenLen
				r.calculate(recIdx+1, damIdx+1, 0, cache)
			}
		} else {
			r.calculate(recIdx+1, damIdx, 0, cache)
			r.calculate(recIdx+1, damIdx, 1, cache)
		}

	case '.':
		if brokenLen > 0 {
			if brokenLen < r.damaged[damIdx] {
				// invalid
				break
			} else if brokenLen == r.damaged[damIdx] {
				// increment damaged streak index and reset brokenLen
				r.calculate(recIdx+1, damIdx+1, 0, cache)
			}
		} else {
			r.calculate(recIdx+1, damIdx, 0, cache)
		}

	case '#':
		// all damaged blocks used up, invalid
		if damIdx == len(r.damaged) {
			break
		}

		if brokenLen > 0 {
			if brokenLen < r.damaged[damIdx] {
				// increment brokenLen if current streak is less than the record
				r.calculate(recIdx+1, damIdx, brokenLen+1, cache)
			} else if brokenLen == r.damaged[damIdx] {
				// invalid
			}
		} else {
			r.calculate(recIdx+1, damIdx, 1, cache)
		}
	}
	cache.Set(recIdx, damIdx, brokenLen, r.possibilities-prevPos)
}

func partOne(input []string) int {
	res := 0
	for _, line := range input {
		r := parseRecord(line)
		r.calculate(0, 0, 0, utils.NewThreeDGrid[int]())
		res += r.possibilities
	}
	return res
}

func partTwo(input []string) int {
	res := 0
	rem := len(input)
	lk := sync.Mutex{}
	var wg sync.WaitGroup

	for _, line := range input {
		line := line
		wg.Add(1)

		go func() {
			defer wg.Done()
			r := parseRecord(line)

			recs := make([]string, 0)
			dam := r.damaged
			r.damaged = make([]int, 0)
			for i := 0; i < 5; i++ {
				recs = append(recs, r.rec)
				r.damaged = append(r.damaged, dam...)
			}
			r.rec = strings.Join(recs, "?")

			r.calculate(0, 0, 0, utils.NewThreeDGrid[int]())
			lk.Lock()
			defer lk.Unlock()
			res += r.possibilities
			rem--
		}()
	}

	wg.Wait()
	return res
}
