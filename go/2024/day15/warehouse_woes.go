package day15

import (
	"fmt"
	"strings"

	"AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string) int {
	w := newWarehouse(lines)
	i := w.height + 1

	for i < len(lines) {
		for _, d := range lines[i] {
			//w.print()
			w.move(dirToDxDy(d))
		}
		i += 1
	}

	//w.print()
	return w.gps()
}

type warehouse struct {
	floor                 *utils.TwoDGrid[int]
	width, height, rx, ry int
}

var (
	wall     = 1
	boxLeft  = 2
	boxRight = 3
)

func newWarehouse(lines []string) *warehouse {
	w := warehouse{
		floor: utils.NewTwoDGrid[int](),
		width: len(lines[0]) * 2,
	}

	for y, line := range lines {
		if strings.TrimSpace(line) == "" {
			w.height = y
			break
		}

		for x, c := range line {
			switch c {
			case '#':
				w.floor.Set(2*x, y, wall)
				w.floor.Set(2*x+1, y, wall)

			case 'O':
				w.floor.Set(2*x, y, boxLeft)
				w.floor.Set(2*x+1, y, boxRight)

			case '@':
				w.rx = 2 * x
				w.ry = y
			}
		}
	}

	return &w
}

func (w *warehouse) move(dx, dy int) {
	nx := w.rx + dx
	ny := w.ry + dy

	x := nx
	y := ny
	// simple handling for horizontal movement
	if dy == 0 {
		for w.floor.Has(x, y) && w.floor.Grid[x][y] != wall {
			x += dx
			y += dy
		}

		// cannot move
		if w.floor.Grid[x][y] == wall {
			return
		}

		// nudge each item starting from the furthest away
		for x != nx || y != ny {
			px := x - dx
			py := y - dy

			w.floor.Set(x, y, w.floor.Grid[px][py])

			x = px
			y = py
		}
		w.floor.Delete(nx, ny)
	} else {
		moves := [][]*move{
			{{w.rx, w.ry, 0}},
		}
		canMove := true
		mi := 0
		for len(moves[mi]) != 0 && canMove {
			mj := 0
			curMoves := make([]*move, 0)
			seen := utils.NewTwoDGrid[bool]()
			for mj < len(moves[mi]) {
				mx := moves[mi][mj].x + dx
				my := moves[mi][mj].y + dy
				mj += 1

				// empty space or already processed
				if !w.floor.Has(mx, my) || seen.Has(mx, my) {
					continue
				}

				item := w.floor.Grid[mx][my]
				seen.Set(mx, my, true)
				// hit a wall for some box
				if item == wall {
					canMove = false
					break
				}

				curMoves = append(curMoves, &move{mx, my, item})

				// other part of box
				ox := mx + 1
				oy := my
				if item == boxRight {
					ox = mx - 1
				}

				if seen.Has(ox, oy) {
					continue
				}
				seen.Set(ox, oy, true)
				otherItem := w.floor.Grid[ox][oy]
				curMoves = append(curMoves, &move{ox, oy, otherItem})
			}
			mi += 1
			moves = append(moves, curMoves)
		}

		if !canMove {
			return
		}

		i := len(moves) - 1
		for i >= 0 {
			for _, m := range moves[i] {
				if m.boxDir == 0 {
					continue
				}
				w.floor.Set(m.x+dx, m.y+dy, m.boxDir)
				w.floor.Delete(m.x, m.y)
			}
			i -= 1
		}
	}

	w.rx = nx
	w.ry = ny
}

func (w *warehouse) gps() int {
	g := 0
	w.floor.ForEach(func(x, y int, v int) {
		if v != boxLeft {
			return
		}
		g += x + 100*y
	})
	return g
}

func (w *warehouse) print() {
	y := 0
	for y < w.height {
		x := 0
		for x < w.width {
			if w.floor.Has(x, y) {
				switch w.floor.Grid[x][y] {
				case wall:
					fmt.Print("#")
				case boxLeft:
					fmt.Print("[")
				case boxRight:
					fmt.Print("]")
				}
			} else if x == w.rx && y == w.ry {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
			x += 1
		}
		y += 1
		fmt.Println()
	}
}

func dirToDxDy(dir int32) (int, int) {
	switch dir {
	case '^':
		return 0, -1
	case '>':
		return 1, 0
	case 'v':
		return 0, 1
	case '<':
		return -1, 0
	}

	return 0, 0
}

type move struct {
	x, y, boxDir int
}
