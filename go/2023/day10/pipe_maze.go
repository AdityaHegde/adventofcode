package day10

import (
	"fmt"

	"AdityaHegde/adventofcode/go/utils"
)

type pipeType int

const (
	Ground pipeType = iota
	Vertical
	Horizontal
	NorthEast
	NorthWest
	SouthWest
	SouthEast
)

type point struct {
	x, y int
}

var (
	North = point{0, -1}
	South = point{0, 1}
	East  = point{1, 0}
	West  = point{-1, 0}
)

var pipeTypeMap = map[rune]pipeType{
	'.': Ground,
	'|': Vertical,
	'-': Horizontal,
	'L': NorthEast,
	'J': NorthWest,
	'7': SouthWest,
	'F': SouthEast,
}
var pipeConnectionsMap = map[pipeType][]point{
	Ground: {
		{0, 0},
		{0, 0},
	},
	Vertical: {
		North,
		South,
	},
	Horizontal: {
		West,
		East,
	},
	NorthEast: {
		North,
		East,
	},
	NorthWest: {
		West,
		North,
	},
	SouthEast: {
		East,
		South,
	},
	SouthWest: {
		South,
		West,
	},
}

func (pt pipeType) connects(px, py, x, y int) (int, bool) {
	if pipeConnectionsMap[pt][0].connects(px, py, x, y) {
		return 0, true
	}
	if pipeConnectionsMap[pt][1].connects(px, py, x, y) {
		return 1, true
	}
	return 0, false
}

func (pt pipeType) progress(idx, x, y int) (int, int) {
	//fmt.Printf("%d,%d =(%d/%d)> %d,%d (%v)\n", x, y, pt, idx, x+pipeConnectionsMap[pt][idx].x, y+pipeConnectionsMap[pt][idx].y, pipeConnectionsMap[pt][idx])
	return x + pipeConnectionsMap[pt][idx].x, y + pipeConnectionsMap[pt][idx].y
}

func (pt pipeType) divide() ([]point, []point) {
	switch pt {
	case Vertical:
		return []point{East}, []point{West}
	case Horizontal:
		return []point{North}, []point{South}
	default:
		c := pipeConnectionsMap[pt]
		return []point{
				{c[0].x + c[1].x, c[0].y + c[1].y},
			}, []point{
				{-c[0].x, -c[0].y},
				{-c[1].x, -c[1].y},
				{-c[0].x - c[1].x, -c[0].y - c[1].y},
			}
	}
}

func (pt pipeType) toString() string {
	for k, p := range pipeTypeMap {
		if p == pt {
			return string(k)
		}
	}
	return "."
}

func (pc point) connects(px, py, x, y int) bool {
	return (x+pc.x == px) && (y+pc.y == py)
}

type pipeGrid [][]pipeType

func parse(input []string, spt pipeType) (*pipeGrid, int, int) {
	var sx, sy int
	grid := make(pipeGrid, len(input))

	for y, line := range input {
		grid[y] = make([]pipeType, len(line))
		for x, c := range line {
			if c == 'S' {
				sx = x
				sy = y
			} else {
				grid[y][x] = pipeTypeMap[c]
			}
		}
	}

	grid[sy][sx] = spt

	return &grid, sx, sy
}

func (pg *pipeGrid) progress(px, py, cx, cy int) (int, int, int, int, int) {
	cpt := (*pg)[cy][cx]
	idx, ok := cpt.connects(px, py, cx, cy)
	if !ok {
		panic(fmt.Sprintf("failed to move. %d,%d => %d,%d (%v)", px, py, cx, cy, cpt))
	}

	nidx := idx
	if idx == 0 {
		nidx = 1
	} else {
		nidx = 0
	}

	nx, ny := cpt.progress(nidx, cx, cy)
	return cx, cy, nx, ny, idx
}

func (pg *pipeGrid) trace(cb func(x, y, idx int) bool, sx, sy int) {
	x0 := sx
	px0 := sx
	y0 := sy
	py0 := sy
	x0, y0 = (*pg)[sy][sx].progress(0, px0, py0)
	idx := -1

	for x0 != sx || y0 != sy {
		if cb(px0, py0, idx) {
			break
		}
		px0, py0, x0, y0, idx = pg.progress(px0, py0, x0, y0)
	}
	cb(px0, py0, idx)
}

func (pg *pipeGrid) traceDivisions(cb func(x, y int, left, right []point) bool, sx, sy int) {
	pg.trace(func(x, y, idx int) bool {
		if idx == -1 {
			return false
		}
		pt := (*pg)[y][x]
		left, right := pt.divide()
		if idx == 1 {
			left, right = right, left
		}
		return cb(x, y, left, right)
	}, sx, sy)
}

func (pg *pipeGrid) checkOutOfBounds(x, y int) bool {
	return x < 0 || x >= len((*pg)[0]) || y < 0 || y >= len(*pg)
}

func partOne(input []string, spt pipeType) int {
	pipes, sx, sy := parse(input, spt)
	ct := 0
	pipes.trace(func(_, _, _ int) bool {
		ct++
		return false
	}, sx, sy)
	return (ct + 1) / 2
}

type soln struct {
	pg          *pipeGrid
	sx, sy      int
	leftOutside bool
	path        *utils.TwoDGrid[bool]
	outside     *utils.TwoDGrid[bool]
}

func newSolution(input []string, spt pipeType) *soln {
	pg, sx, sy := parse(input, spt)
	path := utils.NewTwoDGrid[bool]()
	pg.trace(func(x, y, _ int) bool {
		path.Set(x, y, true)
		return false
	}, sx, sy)

	s := &soln{
		pg:      pg,
		sx:      sx,
		sy:      sy,
		path:    path,
		outside: utils.NewTwoDGrid[bool](),
	}
	s.markAbsoluteOutside()
	s.identifyLeftRight()
	s.traceInsideOutside()

	return s
}

func (s *soln) markAbsoluteOutside() {
	// calculate absolute outside, coming from left-right and right-left
	for y := 0; y < len(*s.pg); y++ {
		for x := 0; x < len((*s.pg)[0]); x++ {
			if s.path.Has(x, y) {
				break
			}
			s.outside.Set(x, y, true)
		}

		for x := len((*s.pg)[0]) - 1; x >= 0; x-- {
			if s.path.Has(x, y) {
				break
			}
			s.outside.Set(x, y, true)
		}
	}
	// calculate absolute outside, coming top-bottom and bottom-top
	for x := 0; x < len((*s.pg)[0]); x++ {
		for y := 0; y < len(*s.pg); y++ {
			if s.path.Has(x, y) {
				break
			}
			s.outside.Set(x, y, true)
		}

		for y := len(*s.pg) - 1; y >= 0; y-- {
			if s.path.Has(x, y) {
				break
			}
			s.outside.Set(x, y, true)
		}
	}
}

func (s *soln) identifyLeftRight() {
	// identify if left is outside or right
	s.pg.traceDivisions(func(x, y int, left, right []point) bool {
		for _, lc := range left {
			if s.path.Has(x+lc.x, y+lc.y) || !s.outside.Has(x+lc.x, y+lc.y) {
				continue
			}
			s.leftOutside = true
			return true
		}
		for _, rc := range right {
			if s.path.Has(x+rc.x, y+rc.y) || !s.outside.Has(x+rc.x, y+rc.y) {
				continue
			}
			s.leftOutside = false
			return true
		}

		return false
	}, s.sx, s.sy)
}

func (s *soln) traceInsideOutside() {
	// trace and mark left/right as outside
	s.pg.traceDivisions(func(x, y int, left, right []point) bool {
		//fmt.Println(x, y, (*s.pg)[y][x].toString())
		for _, lc := range left {
			if s.path.Has(x+lc.x, y+lc.y) || s.outside.Has(x+lc.x, y+lc.y) {
				//fmt.Println("Left seen", x+lc.x, y+lc.y)
				continue
			}
			//fmt.Println("left", x+lc.x, y+lc.y, s.leftOutside)
			s.outside.Set(x+lc.x, y+lc.y, s.leftOutside)
		}
		for _, rc := range right {
			if s.path.Has(x+rc.x, y+rc.y) || s.outside.Has(x+rc.x, y+rc.y) {
				//fmt.Println("Right seen", x+rc.x, y+rc.y)
				continue
			}
			//fmt.Println("right", x+rc.x, y+rc.y, !s.leftOutside)
			s.outside.Set(x+rc.x, y+rc.y, !s.leftOutside)
		}

		return false
	}, s.sx, s.sy)
}

func (s *soln) print() {
	for y := 0; y < len(*s.pg); y++ {
		for x := 0; x < len((*s.pg)[0]); x++ {
			if s.path.Has(x, y) {
				fmt.Print((*s.pg)[y][x].toString())
			} else if s.outside.Has(x, y) {
				o := s.outside.Grid[x][y]
				if o {
					fmt.Print("O")
				} else {
					fmt.Print("I")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func partTwo(input []string, spt pipeType) int {
	s := newSolution(input, spt)

	area := 0
	areaMap := utils.NewTwoDGrid[bool]()
	floodFill := func(x, y int) {
		if s.pg.checkOutOfBounds(x, y) || s.path.Has(x, y) {
			return
		}

		visited := make([]point, 0)
		seen := utils.NewTwoDGrid[bool]()
		list := utils.NewLinkedList[point]()
		list.Push(point{x, y})
		add := func(x, y int) {
			if s.pg.checkOutOfBounds(x, y) || seen.Has(x, y) || s.path.Has(x, y) {
				return
			}
			//fmt.Println("add", x, y)
			list.Push(point{x, y})
			seen.Set(x, y, true)
		}

		//fmt.Println("floodFill", x, y)
		for !list.Empty() {
			p := list.Shift().Value
			//fmt.Println("visit", p, (*s.pg)[p.y][p.x].toString())
			if areaMap.Has(p.x, p.y) {
				// flood fill hit a visited node.
				// mark all visited node with the same inside/outside value as p
				io := areaMap.Grid[p.x][p.y]
				for _, vp := range visited {
					areaMap.Set(vp.x, vp.y, io)
					//fmt.Println(vp, io)
					if !io {
						area++
					}
				}
				break
			}
			if s.outside.Has(p.x, p.y) {
				// flood fill hit an outside node.
				// mark all visited node with the same inside/outside value as p
				io := s.outside.Grid[p.x][p.y]
				areaMap.Set(p.x, p.y, io)
				//fmt.Println(p, io)
				if !io {
					area++
				}
				for _, vp := range visited {
					areaMap.Set(vp.x, vp.y, io)
					//fmt.Println(vp, io)
					if !io {
						area++
					}
				}
				break
			}

			visited = append(visited, p)
			add(p.x-1, p.y)
			add(p.x+1, p.y)
			add(p.x, p.y-1)
			add(p.x, p.y+1)
		}
	}

	for y := 0; y < len(*s.pg); y++ {
		for x := 0; x < len((*s.pg)[0]); x++ {
			floodFill(x-1, y)
			floodFill(x+1, y)
			floodFill(x, y+1)
			floodFill(x, y-1)
		}
	}

	return area
}
