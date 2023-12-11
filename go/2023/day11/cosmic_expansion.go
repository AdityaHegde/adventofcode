package day11

import (
	"fmt"

	"AdityaHegde/adventofcode/go/utils"
)

type galaxy struct {
	x, y int
}

type universe struct {
	width, height int
	grid          *utils.TwoDGrid[bool]
	emptyRows     map[int]bool
	galaxies      []galaxy
	scale         int
}

func (u *universe) distance(g1, g2 galaxy) int {
	return u.horizontalDistance(g1.x, g2.x) + u.verticalDistance(g1.y, g2.y)
}

func (u *universe) horizontalDistance(x1, x2 int) int {
	minx := min(x1, x2)
	maxx := max(x1, x2)
	d := 0
	for x := minx; x < maxx; x++ {
		if _, ok := u.grid.Grid[x]; !ok {
			d += u.scale
		} else {
			d++
		}
	}
	return d
}

func (u *universe) verticalDistance(y1, y2 int) int {
	miny := min(y1, y2)
	maxy := max(y1, y2)
	d := 0
	for y := miny; y < maxy; y++ {
		if _, ok := u.emptyRows[y]; ok {
			d += u.scale
		} else {
			d++
		}
	}
	return d
}

func (u *universe) print() {
	for y := 0; y < u.height; y++ {
		for x := 0; x < u.width; x++ {
			if u.grid.Has(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func parse(input []string, scale int) *universe {
	grid := utils.NewTwoDGrid[bool]()

	emptyRows := map[int]bool{}
	for y, line := range input {
		hasGalaxy := false
		for x, c := range line {
			if c == '#' {
				grid.Set(x, y, true)
				hasGalaxy = true
			}
		}

		if !hasGalaxy {
			emptyRows[y] = true
		}
	}

	galaxies := make([]galaxy, 0)
	grid.ForEach(func(x, y int, _ bool) {
		galaxies = append(galaxies, galaxy{x, y})
	})

	return &universe{
		width:     len(input[0]),
		height:    len(input),
		grid:      grid,
		emptyRows: emptyRows,
		galaxies:  galaxies,
		scale:     scale,
	}
}

func partOne(input []string, scale int) int {
	res := 0
	u := parse(input, scale)
	for i, g := range u.galaxies {
		for j := i + 1; j < len(u.galaxies); j++ {
			res += u.distance(g, u.galaxies[j])
		}
	}
	return res
}
