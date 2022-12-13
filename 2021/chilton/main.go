package main

import (
  "container/heap"
  "fmt"

  "AdityaHegde/adventofcode/utils"
)

func main() {
  lines := utils.InputLines()
  grid := parseGrid(lines)
  fmt.Println(partOne(grid))
  fmt.Println(partTwo(grid))
}

func parseGrid(lines []string) [][]int {
  grid := make([][]int, len(lines))
  for i, line := range lines {
    grid[i] = make([]int, len(line))
    for j, num := range line {
      grid[i][j] = int(num - '0')
    }
  }
  return grid
}

type risk struct {
  risk int
  len  int
  i    int
  j    int
  idx  int
}

type riskHeap []*risk

func (r riskHeap) Len() int {
  return len(r)
}

func (r riskHeap) Less(i, j int) bool {
  return r[i].risk < r[j].risk
}

func (r riskHeap) Swap(i, j int) {
  r[i], r[j] = r[j], r[i]
  r[i].idx = i
  r[j].idx = j
}

func (r *riskHeap) Push(x interface{}) {
  *r = append(*r, x.(*risk))
  x.(*risk).idx = r.Len() - 1
}

func (r *riskHeap) Pop() interface{} {
  old := *r
  n := len(old)
  x := old[n-1]
  *r = old[0 : n-1]
  x.idx = -1
  return x
}

type soln struct {
  rh      *riskHeap
  visited map[string]*risk
}

func newSoln() *soln {
  s := &soln{
    rh:      &riskHeap{},
    visited: make(map[string]*risk),
  }
  heap.Push(s.rh, &risk{
    risk: 0,
    len:  0,
    i:    0,
    j:    0,
  })
  return s
}

func (s *soln) add(r *risk, i, j, addnRisk int) {
  newRisk := r.risk + addnRisk
  key := fmt.Sprintf("%d_%d", i, j)
  if existing, ok := s.visited[key]; ok {
    if newRisk > existing.risk || existing.idx == -1 {
      return
    }
    existing.risk = newRisk
    heap.Fix(s.rh, existing.idx)
  } else {
    s.visited[key] = &risk{
      risk: newRisk,
      len:  r.len + 1,
      i:    i,
      j:    j,
    }
    heap.Push(s.rh, s.visited[key])
  }
}

func partOne(grid [][]int) int {
  s := newSoln()

  add := func(r *risk, i, j int) {
    s.add(r, i, j, grid[i][j])
  }

  for s.rh.Len() > 0 {
    r := heap.Pop(s.rh).(*risk)
    if r.len == len(grid)+len(grid[0])-2 {
      return r.risk
    }
    if r.i < len(grid)-1 {
      add(r, r.i+1, r.j)
    }
    if r.j < len(grid)-1 {
      add(r, r.i, r.j+1)
    }
  }
  return 0
}

func partTwo(grid [][]int) int {
  s := newSoln()
  l := len(grid) * 5

  add := func(r *risk, i, j int) {
    ig := i % len(grid)
    ip := i / len(grid)
    jg := j % len(grid)
    jp := j / len(grid)
    addnRisk := (grid[ig][jg]+ip+jp-1)%9 + 1
    s.add(r, i, j, addnRisk)
  }

  for s.rh.Len() > 0 {
    r := heap.Pop(s.rh).(*risk)
    if r.len == l+l-2 {
      return r.risk
    }
    if r.i < l-1 {
      add(r, r.i+1, r.j)
    }
    if r.j < l-1 {
      add(r, r.i, r.j+1)
    }
  }
  return 0
}
