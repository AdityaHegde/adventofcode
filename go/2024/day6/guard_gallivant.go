package day6

import (
	"AdityaHegde/adventofcode/go/utils"
)

var dirs = []struct {
	x, y int
}{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func partOne(lines []string) int {
	gm := guardMap{}
	gm.parse(lines)

	res := 0
	x := gm.sx
	y := gm.sy
	dir := 0

	visited := utils.NewTwoDGrid[bool]()

	for {
		if !visited.Has(x, y) {
			res += 1
			visited.Set(x, y, true)
		}

		nx := x + dirs[dir].x
		ny := y + dirs[dir].y

		if gm.outOfBounds(nx, ny) {
			break
		}

		if gm.mp.Has(nx, ny) {
			dir = (dir + 1) % len(dirs)
		} else {
			x = nx
			y = ny
		}
	}

	return res
}

func partTwo(lines []string) int {
	gm := guardMap{}
	gm.parse(lines)

	res := 0
	x := gm.sx
	y := gm.sy
	dir := 0

	circularVisited := utils.NewTwoDGrid[bool]()

	for {
		nx := x + dirs[dir].x
		ny := y + dirs[dir].y

		if gm.outOfBounds(nx, ny) {
			break
		}

		if gm.mp.Has(nx, ny) {
			dir = (dir + 1) % len(dirs)
			continue
		}

		rx, ry, rdir := rotateRight(x, y, dir)
		if !gm.isStart(nx, ny) && !gm.outOfBounds(rx, ry) && !circularVisited.Has(nx, ny) && gm.checkCircular(rx, ry, rdir, nx, ny) {
			circularVisited.Set(nx, ny, true)
			res += 1
		}

		x = nx
		y = ny
	}

	return res
}

type guardMap struct {
	mp            *utils.TwoDGrid[bool]
	width, height int
	sx, sy        int
}

func (g *guardMap) parse(lines []string) {
	g.mp = utils.NewTwoDGrid[bool]()
	g.width = len(lines[0])
	g.height = len(lines)

	g.sx = 0
	g.sy = 0

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				g.mp.Set(x, y, true)
			}
			if c == '^' {
				g.sx = x
				g.sy = y
			}
		}
	}
}

func (g *guardMap) checkCircular(x, y, dir, obx, oby int) bool {
	visited := utils.NewThreeDGrid[bool]()

	for {
		if visited.Has(x, y, dir) {
			return true
		}

		visited.Set(x, y, dir, true)

		nx := x + dirs[dir].x
		ny := y + dirs[dir].y

		if g.outOfBounds(nx, ny) {
			return false
		}

		if g.mp.Has(nx, ny) || (nx == obx && ny == oby) {
			dir = (dir + 1) % len(dirs)
		} else {
			x = nx
			y = ny
		}
	}
}

func (g *guardMap) outOfBounds(x, y int) bool {
	return x < 0 || x >= g.width || y < 0 || y >= g.height
}

func (g *guardMap) isStart(x, y int) bool {
	return x == g.sx && y == g.sy
}

func rotateRight(x, y, dir int) (int, int, int) {
	dir = (dir + 1) % len(dirs)
	return x + dirs[dir].x, y + dirs[dir].y, dir
}
