package blizzard_basin

import (
  "AdityaHegde/adventofcode/go/utils"
)

type shard struct {
  x, y         int
  max          int
  xSign, ySign int
}

func newShard(x, y, maxX, maxY int, dir uint8) *shard {
  switch dir {
  case '^':
    return &shard{x, y, maxY, 0, -1}
  case '>':
    return &shard{x, y, maxX, 1, 0}
  case 'v':
    return &shard{x, y, maxY, 0, 1}
  case '<':
    return &shard{x, y, maxX, -1, 0}
  }
  return nil
}

func (s *shard) blocks(x, y, min int) bool {
  min = min % s.max
  sx := s.x
  sy := s.y
  if s.xSign == 0 {
    sy = (s.max + sy + min*s.ySign) % s.max
  } else {
    sx = (s.max + sx + min*s.xSign) % s.max
  }
  return sx == x && sy == y
}

type problem struct {
  columnShards [][]*shard
  rowShards    [][]*shard
  maxX, maxY   int
}

func newProblem(lines []string) *problem {
  maxX := len(lines[0]) - 2
  maxY := len(lines) - 2

  p := &problem{
    make([][]*shard, maxX),
    make([][]*shard, maxY),
    maxX,
    maxY,
  }

  for x := 0; x < maxX; x++ {
    p.columnShards[x] = make([]*shard, 0)
  }

  for y := 0; y < maxY; y++ {
    p.rowShards[y] = make([]*shard, 0)
    for x := 0; x < maxX; x++ {
      c := lines[y+1][x+1]
      if c == '.' {
        continue
      }
      s := newShard(x, y, maxX, maxY, c)
      if s.xSign == 0 {
        p.columnShards[x] = append(p.columnShards[x], s)
      } else {
        p.rowShards[y] = append(p.rowShards[y], s)
      }
    }
  }

  return p
}

func (p *problem) blocks(x, y, min int) bool {
  if x < 0 || x >= p.maxX || y < 0 || y >= p.maxY {
    return true
  }
  for _, s := range p.columnShards[x] {
    if s.blocks(x, y, min) {
      return true
    }
  }
  for _, s := range p.rowShards[y] {
    if s.blocks(x, y, min) {
      return true
    }
  }
  return false
}

var offsets = [][]int{
  {0, -1},
  {1, 0},
  {0, 1},
  {-1, 0},
}

type node struct {
  x, y, min int
}

type solution struct {
  prob    *problem
  list    *utils.LinkedList[*node]
  visited *utils.ThreeDGrid[bool]
}

func newSolution(p *problem) *solution {
  return &solution{p, utils.NewLinkedList[*node](), utils.NewThreeDGrid[bool]()}
}

func (s *solution) solve(p *problem, fromX, fromY, toX, toY, min int) int {
  s.list.Push(&node{fromX, fromY, min})
  for !s.list.Empty() {
    n := s.list.Shift().Value
    //fmt.Println("Visited", n.min, n.x, n.y)
    if n.x == toX && n.y == toY {
      return n.min + 1
    }

    n.min++

    // check 4 directions
    for _, offset := range offsets {
      nn := &node{n.x + offset[0], n.y + offset[1], n.min}
      if !p.blocks(nn.x, nn.y, nn.min) && !s.visited.Has(nn.x, nn.y, nn.min) {
        if nn.x == toX && nn.y == toY {
          return nn.min + 1
        }
        s.visited.Set(nn.x, nn.y, nn.min, true)
        s.list.Push(nn)
      }
    }

    if (n.y == -1 || n.y == p.maxY || !p.blocks(n.x, n.y, n.min)) && !s.visited.Has(n.x, n.y, n.min) {
      // add entry to wait
      s.list.Push(n)
      s.visited.Set(n.x, n.y, n.min, true)
    }
  }
  return 0
}

func partOne(lines []string) int {
  p := newProblem(lines)
  s := newSolution(p)
  return s.solve(p, 0, -1, p.maxX-1, p.maxY-1, 0)
}

func partTwo(lines []string) int {
  p := newProblem(lines)
  s1 := newSolution(p)
  min1 := s1.solve(p, 0, -1, p.maxX-1, p.maxY-1, 0)
  s2 := newSolution(p)
  min2 := s2.solve(p, p.maxX-1, p.maxY, 0, 0, min1)
  s3 := newSolution(p)
  return s3.solve(p, 0, -1, p.maxX-1, p.maxY-1, min2)
}
