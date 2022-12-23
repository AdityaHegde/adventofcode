package unstable_diffusion

import (
  "fmt"
  "math"

  "AdityaHegde/adventofcode/go/utils"
)

type problem struct {
  grid       *utils.TwoDGrid[bool]
  minX, minY int
  maxX, maxY int
  elves      int
  orderIdx   int
}

var dirToOffset = map[int][]int{
  0: {0, -1},
  1: {1, -1},
  2: {1, 0},
  3: {1, 1},
  4: {0, 1},
  5: {-1, 1},
  6: {-1, 0},
  7: {-1, -1},
}
var checkOrder = []int{0, 4, 6, 2}

func newProblem(lines []string) *problem {
  p := &problem{
    utils.NewTwoDGrid[bool](),
    math.MaxInt,
    math.MaxInt,
    math.MinInt,
    math.MinInt,
    0,
    0,
  }
  for y, line := range lines {
    for x, c := range line {
      if c == '#' {
        p.grid.Set(x, y, true)
        p.updateMinMax(x, y)
        p.elves++
      }
    }
  }
  return p
}

func (p *problem) updateMinMax(x, y int) {
  if x < p.minX {
    p.minX = x
  }
  if x > p.maxX {
    p.maxX = x
  }
  if y < p.minY {
    p.minY = y
  }
  if y > p.maxY {
    p.maxY = y
  }
}

func (p *problem) resetMinMax() {
  p.minX = math.MaxInt
  p.minY = math.MaxInt
  p.maxX = math.MinInt
  p.maxY = math.MinInt
}

func (p *problem) checkDir(x, y, dir int) bool {
  offsets := dirToOffset[dir]
  return p.grid.Has(x+offsets[0], y+offsets[1])
}

func (p *problem) getMoveDir(x, y int) int {
  elves := 0
  moveDir := -1
  for i := 0; i < 4; i++ {
    dir := checkOrder[(i+p.orderIdx)%4]
    if !p.checkDir(x, y, (8+dir-1)%8) && !p.checkDir(x, y, dir) && !p.checkDir(x, y, dir+1) {
      elves++
      if moveDir == -1 {
        moveDir = dir
      }
    }
  }
  if elves == 4 {
    return -1
  }
  return moveDir
}

func (p *problem) print() {
  fmt.Println("*************")
  for y := p.minY; y <= p.maxY; y++ {
    for x := p.minX; x <= p.maxX; x++ {
      if p.grid.Has(x, y) {
        fmt.Print("#")
      } else {
        fmt.Print(".")
      }
    }
    fmt.Println()
  }
}

type iteration struct {
  moves     *utils.TwoDGrid[[]int]
  dedupe    *utils.TwoDGrid[int]
  moveCount int
}

func newIteration() *iteration {
  return &iteration{
    utils.NewTwoDGrid[[]int](),
    utils.NewTwoDGrid[int](),
    0,
  }
}

func (i *iteration) phaseOne(p *problem) {
  p.grid.ForEach(func(x, y int, _ bool) {
    moveDir := p.getMoveDir(x, y)
    if moveDir == -1 {
      p.updateMinMax(x, y)
      return
    }

    offsets := dirToOffset[moveDir]
    nx := x + offsets[0]
    ny := y + offsets[1]
    i.moves.Set(x, y, []int{nx, ny})
    prev := 0
    if i.dedupe.Has(nx, ny) {
      prev = i.dedupe.Grid[nx][ny]
    }
    i.dedupe.Set(nx, ny, prev+1)
  })
}

func (i *iteration) phaseTwo(p *problem) {
  i.moves.ForEach(func(x, y int, newXY []int) {
    nx := newXY[0]
    ny := newXY[1]
    if i.dedupe.Grid[nx][ny] == 1 {
      p.grid.Delete(x, y)
      p.grid.Set(nx, ny, true)
      p.updateMinMax(nx, ny)
      i.moveCount++
    } else {
      p.updateMinMax(x, y)
    }
  })
}

func partOne(lines []string) int {
  p := newProblem(lines)

  for i := 0; i < 10; i++ {
    p.resetMinMax()
    iter := newIteration()
    iter.phaseOne(p)
    iter.phaseTwo(p)
    p.orderIdx = (p.orderIdx + 1) % 4
  }

  return (p.maxX-p.minX+1)*(p.maxY-p.minY+1) - p.elves
}

func partTwo(lines []string) int {
  p := newProblem(lines)

  for i := 0; true; i++ {
    p.resetMinMax()
    iter := newIteration()
    iter.phaseOne(p)
    iter.phaseTwo(p)
    p.orderIdx = (p.orderIdx + 1) % 4
    if iter.moveCount == 0 {
      return i + 1
    }
  }

  return -1
}
