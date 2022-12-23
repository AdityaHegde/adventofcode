package boiling_boulders

import (
  "strings"

  utils2 "AdityaHegde/adventofcode/go/utils"
)

type point struct {
  x, y, z int
}

type problem struct {
  grid    *utils2.ThreeDGrid[bool]
  air     *utils2.ThreeDGrid[int]
  visited *utils2.ThreeDGrid[bool]
  area    int
  minX    int
  minY    int
  minZ    int
  maxX    int
  maxY    int
  maxZ    int
}

func newProblem() *problem {
  return &problem{
    grid:    utils2.NewThreeDGrid[bool](),
    air:     utils2.NewThreeDGrid[int](),
    visited: utils2.NewThreeDGrid[bool](),
    minX:    99999,
    minY:    99999,
    minZ:    99999,
    maxX:    -99999,
    maxY:    -99999,
    maxZ:    -99999,
  }
}

func (p *problem) add(x, y, z int) {
  p.grid.Set(x, y, z, true)
  p.air.Delete(x, y, z)

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
  if z < p.minZ {
    p.minZ = z
  }
  if z > p.maxZ {
    p.maxZ = z
  }

  for d := -1; d <= 1; d += 2 {
    p.area += p.checkSide(x+d, y, z)
    p.area += p.checkSide(x, y+d, z)
    p.area += p.checkSide(x, y, z+d)
  }
}

func (p *problem) parseLines(lines []string) {
  for _, line := range lines {
    coords := strings.Split(line, ",")
    x := utils2.Int(coords[0])
    y := utils2.Int(coords[1])
    z := utils2.Int(coords[2])

    p.add(x, y, z)
  }
  p.minX--
  p.minY--
  p.minZ--
  p.maxX++
  p.maxY++
  p.maxZ++
}

func (p *problem) checkSide(x, y, z int) int {
  if p.grid.Has(x, y, z) {
    return -1
  } else {
    air := 0
    if p.air.Has(x, y, z) {
      air = p.air.Grid[x][y][z] + 1
    } else {
      air = 1
    }
    p.air.Set(x, y, z, air)
    return 1
  }
}

func (p *problem) visit() {
  visited := utils2.NewThreeDGrid[bool]()
  visited.Set(p.minX, p.minY, p.minZ, true)

  list := utils2.LinkedList[*point]{}
  list.Push(&point{p.minX, p.minY, p.minZ})

  check := func(cx, cy, cz int) {
    if cx < p.minX || cx > p.maxX ||
      cy < p.minY || cy > p.maxY ||
      cz < p.minZ || cz > p.maxZ ||
      visited.Has(cx, cy, cz) || p.grid.Has(cx, cy, cz) {
      return
    }

    list.Push(&point{cx, cy, cz})
    visited.Set(cx, cy, cz, true)
    return
  }

  for !list.Empty() {
    pt := list.Shift().Value
    if p.air.Has(pt.x, pt.y, pt.z) {
      p.air.Delete(pt.x, pt.y, pt.z)
    }
    for d := -1; d <= 1; d += 2 {
      check(pt.x+d, pt.y, pt.z)
      check(pt.x, pt.y+d, pt.z)
      check(pt.x, pt.y, pt.z+d)
    }
  }
}

func partOne(lines []string) int {
  p := newProblem()
  p.parseLines(lines)
  return p.area
}

func partTwo(lines []string) int {
  p := newProblem()
  p.parseLines(lines)

  p.visit()
  p.air.ForEach(func(x, y, z, air int) {
    p.area -= air
  })

  return p.area
}
