package hill_climbing

import (
  "container/heap"

  "AdityaHegde/adventofcode/utils"
)

type problem struct {
  grid   [][]int
  startX int
  startY int
  endX   int
  endY   int
}

func parse(lines []string) *problem {
  p := &problem{
    grid: make([][]int, len(lines)),
  }

  for x, line := range lines {
    p.grid[x] = make([]int, len(line))
    for y, c := range line {
      if c == 'S' {
        p.startX = x
        p.startY = y
        p.grid[x][y] = 0
      } else if c == 'E' {
        p.endX = x
        p.endY = y
        p.grid[x][y] = 'z' - 'a'
      } else {
        p.grid[x][y] = int(c - 'a')
      }
    }
  }

  return p
}

type node struct {
  weight int
  dist   int
  x      int
  y      int
  idx    int
}

type nodeHeap []*node

func (nh nodeHeap) Len() int {
  return len(nh)
}

func (nh nodeHeap) Less(x, y int) bool {
  if nh[x].weight == nh[y].weight {
    return nh[x].dist < nh[y].dist
  } else {
    return nh[x].weight < nh[y].weight
  }
}

func (nh nodeHeap) Swap(x, y int) {
  nh[x], nh[y] = nh[y], nh[x]
  nh[x].idx = x
  nh[y].idx = y
}

func (nh *nodeHeap) Push(n interface{}) {
  *nh = append(*nh, n.(*node))
  n.(*node).idx = nh.Len() - 1
}

func (nh *nodeHeap) Pop() interface{} {
  old := *nh
  cnt := len(old)
  n := old[cnt-1]
  *nh = old[0 : cnt-1]
  n.idx = -1
  return n
}

type soln struct {
  p       *problem
  nh      *nodeHeap
  visited map[int]map[int]*node
}

func newSoln(p *problem) *soln {
  s := &soln{
    p:       p,
    nh:      &nodeHeap{},
    visited: map[int]map[int]*node{},
  }
  heap.Push(s.nh, &node{
    weight: 0,
    dist:   0,
    x:      p.startX,
    y:      p.startY,
  })
  return s
}

func (s *soln) add(n *node, x, y int) {
  // bounds check
  if x < 0 || x >= len(s.p.grid) || y < 0 || y >= len(s.p.grid[0]) {
    return
  }

  // height diff check
  if s.p.grid[n.x][n.y]+1 < s.p.grid[x][y] {
    return
  }

  dist := n.dist + 1
  weight := dist + utils.ManhattanDist(x, y, s.p.endX, s.p.endY) + s.p.grid[n.x][n.y] - s.p.grid[x][y]
  if _, ok := s.visited[x]; !ok {
    s.visited[x] = map[int]*node{}
  }

  if existing, ok := s.visited[x][y]; ok {
    if dist < existing.dist || existing.idx == 1 {
      return
    }
    existing.weight = weight
    existing.dist = dist
    heap.Fix(s.nh, existing.idx)
  } else {
    s.visited[x][y] = &node{
      weight: weight,
      dist:   dist,
      x:      x,
      y:      y,
    }
    heap.Push(s.nh, s.visited[x][y])
  }
}

func partOne(p *problem) int {
  s := newSoln(p)

  for s.nh.Len() > 0 {
    n := heap.Pop(s.nh).(*node)
    if n.x == s.p.endX && n.y == s.p.endY {
      return n.dist
    }
    s.add(n, n.x+1, n.y)
    s.add(n, n.x-1, n.y)
    s.add(n, n.x, n.y+1)
    s.add(n, n.x, n.y-1)
  }

  return 0
}
