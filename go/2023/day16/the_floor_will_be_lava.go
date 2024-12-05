package day16

import (
	"sync"

	"AdityaHegde/adventofcode/go/utils"
)

type mirrorType int

const (
	horizontal mirrorType = iota
	vertical
	forward
	backwards
)

var charToType = map[int32]mirrorType{
	'|':  vertical,
	'-':  horizontal,
	'/':  forward,
	'\\': backwards,
}

type point struct {
	px, py int
	x, y   int
}

func (p point) bounce(mt mirrorType) []point {
	switch mt {
	case horizontal:
		if p.py == p.y {
			return []point{
				{p.x, p.y, p.x + (p.x - p.px), p.y},
			}
		}
		return []point{
			{p.x, p.y, p.x + 1, p.y},
			{p.x, p.y, p.x - 1, p.y},
		}
	case vertical:
		if p.px == p.x {
			return []point{
				{p.x, p.y, p.x, p.y + (p.y - p.py)},
			}
		}
		return []point{
			{p.x, p.y, p.x, p.y + 1},
			{p.x, p.y, p.x, p.y - 1},
		}
	case forward:
		if p.py == p.y {
			return []point{
				{p.x, p.y, p.x, p.y - (p.x - p.px)},
			}
		}
		return []point{
			{p.x, p.y, p.x - (p.y - p.py), p.y},
		}
	case backwards:
		if p.py == p.y {
			return []point{
				{p.x, p.y, p.x, p.y + (p.x - p.px)},
			}
		}
		return []point{
			{p.x, p.y, p.x + (p.y - p.py), p.y},
		}
	}
	return nil
}

func (p point) key() int {
	return (1 + p.x - p.px) + (1+p.y-p.py)*3
}

type floor struct {
	height, width int
	mirrors       *utils.TwoDGrid[mirrorType]
}

func parse(input []string) *floor {
	mirrors := utils.NewTwoDGrid[mirrorType]()

	for y, line := range input {
		for x, c := range line {
			if c == '.' {
				continue
			}
			mirrors.Set(x, y, charToType[c])
		}
	}

	return &floor{
		len(input),
		len(input[0]),
		mirrors,
	}
}

func (f *floor) move(p point) []point {
	if f.mirrors.Has(p.x, p.y) {
		return p.bounce(f.mirrors.Grid[p.x][p.y])
	}
	return []point{
		{p.x, p.y, p.x + (p.x - p.px), p.y + (p.y - p.py)},
	}
}

func (f *floor) withinBounds(p point) bool {
	return p.x >= 0 && p.x < f.width && p.y >= 0 && p.y < f.height
}

func (f *floor) energise(init point) int {
	ps := utils.LinkedList[point]{}
	ps.Push(init)
	visited := utils.NewThreeDGrid[bool]()
	energised := utils.NewTwoDGrid[bool]()
	energisedCount := 0

	for !ps.Empty() {
		pn := ps.Shift()
		p := pn.Value
		if visited.Has(p.x, p.y, p.key()) {
			continue
		}

		visited.Set(p.x, p.y, p.key(), true)
		if !energised.Has(p.x, p.y) {
			energised.Set(p.x, p.y, true)
			energisedCount++
		}

		nps := f.move(p)
		for _, np := range nps {
			if f.withinBounds(np) {
				ps.Push(np)
			}
		}
	}

	return energisedCount
}

func partOne(input []string) int {
	f := parse(input)
	return f.energise(point{-1, 0, 0, 0})
}

func partTwo(input []string) int {
	res := 0
	f := parse(input)

	l := sync.Mutex{}
	var wg sync.WaitGroup

	energise := func(p point) {
		defer wg.Done()

		e := f.energise(p)
		l.Lock()
		defer l.Unlock()
		if e > res {
			res = e
		}
	}

	for y := 0; y < f.height; y++ {
		wg.Add(2)
		go energise(point{-1, y, 0, y})
		go energise(point{f.width, y, f.width - 1, y})
	}
	for x := 0; x < f.width; x++ {
		wg.Add(2)
		go energise(point{x, -1, x, 0})
		go energise(point{0, f.height, x, f.height - 1})
	}

	wg.Wait()

	return res
}
