package monkey_map

import (
  "fmt"

  "AdityaHegde/adventofcode/utils"
)

type problem_two struct {
  grid [][]bool
  wrap *utils.ThreeDGrid[[]int]
}

func newProblemTwo(lines []string) *problem_two {
  l := len(lines)
  p := &problem_two{
    grid: make([][]bool, l),
    wrap: utils.NewThreeDGrid[[]int](),
  }
  for y, line := range lines {
    p.grid[y] = make([]bool, 0)
    for _, c := range line {
      switch c {
      case ' ':
        p.grid[y] = append(p.grid[y], false)
      case '.':
        p.grid[y] = append(p.grid[y], true)
      case '#':
        p.grid[y] = append(p.grid[y], false)
      }
    }
  }
  return p
}

func (p *problem_two) addWrap(x1, y1, dir1, x2, y2, dir2 int) {
  p.wrap.Set(x1, y1, dir1, []int{x2, y2, dir2})
}

func (p *problem_two) fillWrap() {
  for dx := 0; dx < 50; dx++ {
    dy := dx

    // x=50-99,  y=-1  ==> x=0, y=150+dx
    // x=-1,y=150-199  ==> x=50+dy,y=0
    p.addWrap(50+dx, -1, 3, 0, 150+dx, 0)
    p.addWrap(-1, 150+dy, 2, 50+dy, 0, 1)

    // x=100-149,y=-1  ==> x=dx,y=199
    // x=0-49,y=200    ==> x=100+dx,y=0
    p.addWrap(100+dx, -1, 3, dx, 199, 3)
    p.addWrap(dx, 200, 1, 100+dx, 0, 1)

    // x=49, y=0-49    ==> x=0, y=149-dy
    // x=-1, y=100-149 ==> x=50,y=49-dy
    p.addWrap(49, dy, 2, 0, 149-dy, 0)
    p.addWrap(-1, 100+dy, 2, 50, 49-dy, 0)

    // x=150,y=0-49    ==> x=99,y=149-dy
    // x=100,y=100-149 ==> x=149,y=49-dy
    p.addWrap(150, dy, 0, 99, 149-dy, 2)
    p.addWrap(100, 100+dy, 0, 149, 49-dy, 2)

    // x=100-149,y=50  ==> x=99,y=50+dx
    // x=100,y=50-99   ==> x=100+dy,y=49
    p.addWrap(100+dx, 50, 1, 99, 50+dx, 2)
    p.addWrap(100, 50+dy, 0, 100+dy, 49, 3)

    // x=50-99,y=150   ==> x=49,y=150+dx
    // x=50,y=150-199  ==> x=50+dy,y=149
    p.addWrap(50+dx, 150, 1, 49, 150+dx, 2)
    p.addWrap(50, 150+dy, 0, 50+dy, 149, 3)

    // x=49,y=50-99    ==> x=dy,y=100
    // x=0-49,y=99     ==> x=50,y=50+dx
    p.addWrap(49, 50+dy, 2, dy, 100, 1)
    p.addWrap(dx, 99, 3, 50, 50+dx, 0)
  }
}

func (p *problem_two) moveInDir(s *solution, l int) {
  fmt.Printf("Move %d,%d in %d for %d\n", s.x, s.y, s.dir, l)
  for l > 0 {
    ny := s.y + s.dy
    nx := s.x + s.dx
    ndir := s.dir

    if p.wrap.Has(nx, ny, ndir) {
      wrapTo := p.wrap.Grid[nx][ny][ndir]
      fmt.Printf("Wrap %d,%d => %d,%d in %d\n", nx, ny, wrapTo[0], wrapTo[1], wrapTo[2])
      nx = wrapTo[0]
      ny = wrapTo[1]
      ndir = wrapTo[2]
    }

    //if len(p.grid) <= ny {
    //  fmt.Println(nx, ny, len(p.grid))
    //} else if len(p.grid[ny]) <= nx {
    //  fmt.Println(nx, ny, len(p.grid), len(p.grid[ny]))
    //}

    if !p.grid[ny][nx] {
      break
    }

    s.x = nx
    s.y = ny
    if s.dir != ndir {
      s.dir = ndir
      s.xyFromDir()
    }
    l--
  }
}

func (p *problem_two) move(s *solution, directions string) {
  n := 0
  nc := 1

  moveInDir := func() {
    p.moveInDir(s, n)
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

func partTwo(lines []string) int {
  s := &solution{50, 0, 1, 0, 0}

  p := newProblemTwo(lines[:len(lines)-2])
  p.fillWrap()
  directions := lines[len(lines)-1]

  p.move(s, directions)
  return s.get()
}
