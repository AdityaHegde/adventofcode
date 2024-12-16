package day16

import (
	"math"

	"AdityaHegde/adventofcode/go/utils"
)

func partOne(input []string) (int, int) {
	m := newMaze(input)
	list := utils.NewLinkedList[*node]()
	cost := utils.NewTwoDGrid[int]()
	list.Push(&node{
		utils.Vector{
			X:   m.start.X,
			Y:   m.start.Y,
			Dir: 1,
		},
		nil,
		0,
	})

	paths := make([]*node, 0)
	endCost := math.MaxInt

	for !list.Empty() {
		n := list.Pop()
		v := n.Value.vec
		if !cost.Has(v.X, v.Y) || cost.Grid[v.X][v.Y] > n.Value.cost {
			cost.Set(v.X, v.Y, n.Value.cost)
		}
		if v.X == m.end.X && v.Y == m.end.Y {
			if endCost > n.Value.cost {
				paths = make([]*node, 0)
				endCost = n.Value.cost
			}
			if n.Value.cost == endCost {
				paths = append(paths, n.Value)
			}
			continue
		}

		for dirIdx, dir := range utils.Directions {
			dx := v.X + dir.X
			dy := v.Y + dir.Y
			rotations := v.Rotations(dirIdx)
			if !m.withinBounds(dx, dy) || m.walls.Has(dx, dy) || rotations == 2 {
				continue
			}
			newCost := n.Value.cost + 1 + v.Rotations(dirIdx)*1000

			if cost.Has(dx, dy) &&
				// add a buffer of a turn to make sure we get 2 paths entering a node with different directions
				// but going forward one will turn one will not.
				cost.Grid[dx][dy] < newCost-1000 {
				continue
			}

			list.Push(&node{
				utils.Vector{
					X:   dx,
					Y:   dy,
					Dir: dirIdx,
				},
				n.Value,
				newCost,
			})
		}
	}

	visited := utils.NewTwoDGrid[bool]()
	ct := 0
	for _, n := range paths {
		for n != nil {
			if !visited.Has(n.vec.X, n.vec.Y) {
				visited.Set(n.vec.X, n.vec.Y, true)
				ct += 1
			}
			n = n.prev
		}
	}

	return cost.Grid[m.end.X][m.end.Y], ct
}

type maze struct {
	walls         *utils.TwoDGrid[bool]
	width, height int
	start, end    utils.Point
}

func newMaze(lines []string) *maze {
	m := maze{
		walls:  utils.NewTwoDGrid[bool](),
		width:  len(lines[0]),
		height: len(lines),
	}
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case '#':
				m.walls.Set(x, y, true)
			case 'S':
				m.start = utils.Point{X: x, Y: y}
			case 'E':
				m.end = utils.Point{X: x, Y: y}
			}
		}
	}
	return &m
}

func (m *maze) withinBounds(x, y int) bool {
	return x >= 0 && x < m.width && y >= 0 && y < m.height
}

type node struct {
	vec  utils.Vector
	prev *node
	cost int
}
