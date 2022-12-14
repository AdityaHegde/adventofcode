package regolith_reservoir

import (
  "fmt"
  "strings"

  "AdityaHegde/adventofcode/utils"
)

type problem struct {
  grid map[int]map[int]bool
  minX int
  maxX int
  minY int
  maxY int
}

func newProblem() *problem {
  return &problem{
    make(map[int]map[int]bool),
    99999,
    0,
    99999,
    0,
  }
}

func (p *problem) add(x, y int) {
  if _, ok := p.grid[y]; !ok {
    p.grid[y] = make(map[int]bool)
  }
  p.grid[y][x] = true
  if x > p.maxX {
    p.maxX = x
  } else if x < p.minX {
    p.minX = x
  }
  if y > p.maxY {
    p.maxY = y
  } else if y < p.minY {
    p.minY = y
  }
}

func (p *problem) has(x, y int) bool {
  if _, ok := p.grid[y]; !ok {
    return false
  }
  return p.grid[y][x]
}

func (p *problem) parsePoint(point string, startX, startY int) (int, int) {
  end := strings.Split(point, ",")
  endX := utils.Int(end[0])
  endY := utils.Int(end[1])

  dx := utils.Sign(endX - startX)
  dy := utils.Sign(endY - startY)
  if dx != 0 {
    for x := startX; x != endX; x += dx {
      p.add(x, startY)
    }
  } else {
    for y := startY; y != endY; y += dy {
      p.add(startX, y)
    }
  }

  return endX, endY
}

func (p *problem) parseLine(line string) {
  points := strings.Split(line, " -> ")
  start := strings.Split(points[0], ",")
  startX := utils.Int(start[0])
  startY := utils.Int(start[1])
  for i := 1; i < len(points); i++ {
    startX, startY = p.parsePoint(points[i], startX, startY)
  }
  p.add(startX, startY)
}

func (p *problem) parse(lines []string) {
  for _, line := range lines {
    p.parseLine(line)
  }
  p.minX -= 2
  p.maxX += 2
  p.minY -= 2
  p.maxY += 2
}

func (p *problem) print() {
  for y := p.minY; y <= p.maxY; y++ {
    for x := p.minX; x <= p.maxX; x++ {
      if _, ok := p.grid[y]; ok {
        if p.grid[y][x] {
          fmt.Print("#")
          continue
        }
      }
      fmt.Print(".")
    }
    fmt.Print("\n")
  }
}

const SandX = 500
const SandY = 0

func partOne(p *problem) int {
  i := 0
  for {
    x := SandX
    y := SandY
    for y < p.maxY {
      if p.has(x, y+1) {
        if !p.has(x-1, y+1) {
          x--
          y++
        } else if !p.has(x+1, y+1) {
          x++
          y++
        } else {
          p.add(x, y)
          break
        }
      } else {
        y++
      }
    }
    if y >= p.maxY {
      break
    }
    i++
  }
  return i
}

func partTwo(p *problem) int {
  i := 0
  for {
    x := SandX
    y := SandY
    for y < p.maxY {
      if p.has(x, y+1) {
        if !p.has(x-1, y+1) {
          x--
          y++
        } else if !p.has(x+1, y+1) {
          x++
          y++
        } else {
          p.add(x, y)
          if y == SandY {
            return i + 1
          }
          break
        }
      } else {
        y++
      }
    }
    if y == p.maxY {
      p.add(x, y-1)
      if y-1 == SandY {
        break
      }
    }
    i++
  }
  return i
}
