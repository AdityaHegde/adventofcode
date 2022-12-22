package monkey_map

import "fmt"

type problem struct {
  grid    [][]bool
  xStarts []int
  yStarts []int
  yEnds   []int
  len     int
}

func newProblem(lines []string) *problem {
  l := len(lines)
  p := &problem{
    grid:    make([][]bool, l),
    xStarts: make([]int, l),
    yStarts: make([]int, 0),
    yEnds:   make([]int, 0),
    len:     l,
  }
  for y, line := range lines {
    start := 0
    p.grid[y] = make([]bool, 0)
    for x, c := range line {
      if x == len(p.yStarts) {
        p.yStarts = append(p.yStarts, y)
      }
      if x == len(p.yEnds) {
        p.yEnds = append(p.yEnds, 0)
      }
      switch c {
      case ' ':
        start++
        if p.yStarts[x] == y {
          p.yStarts[x]++
        }
      case '.':
        p.grid[y] = append(p.grid[y], true)
        if p.yEnds[x] < y {
          p.yEnds[x] = y
        }
      case '#':
        p.grid[y] = append(p.grid[y], false)
        if p.yEnds[x] < y {
          p.yEnds[x] = y
        }
      }
    }
    p.xStarts[y] = start
  }
  return p
}

func (p *problem) moveInDir(x, y, dx, dy, l int) (int, int) {
  for l > 0 {
    ny := y + dy
    nx := x + dx

    if dx != 0 {
      if nx == len(p.grid[ny])+p.xStarts[ny] {
        nx = p.xStarts[ny]
      } else if nx < p.xStarts[ny] {
        nx = len(p.grid[ny]) + p.xStarts[ny] - 1
      }
    } else {
      if ny > p.yEnds[nx] {
        ny = p.yStarts[nx]
      } else if ny < p.yStarts[nx] {
        ny = p.yEnds[nx]
      }
    }

    if nx-p.xStarts[ny] > len(p.grid[ny]) {
      fmt.Println(nx, ny, p.xStarts[ny], len(p.grid[ny]), p.yStarts[nx], p.yEnds[nx])
    }

    if !p.grid[ny][nx-p.xStarts[ny]] {
      break
    }

    x = nx
    y = ny
    l--
  }

  return x, y
}

type solution struct {
  x, y, dx, dy, dir int
}

func (s *solution) xyFromDir() {
  switch s.dir {
  case 0:
    s.dx = 1
    s.dy = 0
  case 1:
    s.dx = 0
    s.dy = 1
  case 2:
    s.dx = -1
    s.dy = 0
  case 3:
    s.dx = 0
    s.dy = -1
  }
}

func (s *solution) turn(dir int) {
  s.dir = (4 + s.dir + dir) % 4
  s.xyFromDir()
}

func (p *problem) move(s *solution, directions string) {
  n := 0
  nc := 1

  moveInDir := func() {
    s.x, s.y = p.moveInDir(s.x, s.y, s.dx, s.dy, n)
    nc = 1
    n = 0
  }

  for _, c := range directions {
    switch c {
    case 'L':
      moveInDir()
      s.turn(-1)
    case 'R':
      moveInDir()
      s.turn(1)
    default:
      n = n*nc + int(c-'0')
      nc *= 10
    }
  }

  moveInDir()
}

func (s *solution) get() int {
  return 1000*(s.y+1) + 4*(s.x+1) + s.dir
}

func (p *problem) print() {
  for y, yg := range p.grid {
    for i := 0; i < p.xStarts[y]; i++ {
      fmt.Print(" ")
    }
    for _, e := range yg {
      if e {
        fmt.Print(".")
      } else {
        fmt.Print("#")
      }
    }
    fmt.Println()
  }
}

func (p *problem) getSolution() *solution {
  return &solution{p.xStarts[0], 0, 1, 0, 0}
}

func partOne(lines []string) int {
  p := newProblem(lines[:len(lines)-2])
  directions := lines[len(lines)-1]

  s := p.getSolution()
  p.move(s, directions)
  return s.get()
}
