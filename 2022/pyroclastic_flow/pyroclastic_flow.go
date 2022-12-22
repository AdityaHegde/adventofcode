package pyroclastic_flow

import "fmt"

type point struct {
  x int
  y int
}

type rock struct {
  points []*point
  xMin   int
  xMax   int
  yMin   int
  yMax   int
}

func newRock(points []*point) *rock {
  r := &rock{
    points: points,
    xMin:   points[0].x,
    xMax:   points[0].x,
    yMin:   points[0].y,
    yMax:   points[0].y,
  }
  for i := 1; i < len(points); i++ {
    if points[i].x < r.xMin {
      r.xMin = points[i].x
    }
    if points[i].x > r.xMax {
      r.xMax = points[i].x
    }
    if points[i].y < r.yMin {
      r.yMin = points[i].y
    }
    if points[i].y > r.yMax {
      r.yMax = points[i].y
    }
  }
  return r
}

func (r *rock) clone() *rock {
  points := make([]*point, len(r.points))
  for i, p := range r.points {
    points[i] = &point{
      x: p.x,
      y: p.y,
    }
  }
  return newRock(points)
}

func (r *rock) offset(x, y int) {
  for _, p := range r.points {
    p.x += x
    p.y += y
  }
  r.xMin += x
  r.xMax += x
  r.yMin += y
  r.yMax += y
}

func (r *rock) print() {
  fmt.Printf("%d -- %d, %d || %d\n", r.xMin, r.xMax, r.yMin, r.yMax)
}

type problem struct {
  rocks []*rock
  grid  [][]bool
}

func newProblem(width int) *problem {
  p := &problem{
    rocks: []*rock{
      // ####
      newRock([]*point{
        {0, 0}, {1, 0}, {2, 0}, {3, 0},
      }),
      // .#.
      // ###
      // .#.
      newRock([]*point{
        {1, 2},
        {0, 1}, {1, 1}, {2, 1},
        {1, 0},
      }),
      // ..#
      // ..#
      // ###
      newRock([]*point{
        {2, 2},
        {2, 1},
        {0, 0}, {1, 0}, {2, 0},
      }),
      // #
      // #
      // #
      // #
      newRock([]*point{
        {0, 0}, {0, 1}, {0, 2}, {0, 3},
      }),
      // ##
      // ##
      newRock([]*point{
        {0, 1}, {1, 1},
        {0, 0}, {1, 0},
      }),
    },
    grid: make([][]bool, width),
  }
  for i := 0; i < width; i++ {
    p.grid[i] = make([]bool, 0)
  }
  return p
}

func (p *problem) moveSide(r *rock, dir int) {
  if r.xMin+dir < 0 || r.xMax+dir >= len(p.grid) {
    return
  }
  for _, pnt := range r.points {
    column := p.grid[pnt.x+dir]
    if pnt.y < len(column) && column[pnt.y] {
      return
    }
  }
  r.offset(dir, 0)
}

func (p *problem) moveDown(r *rock) bool {
  if r.yMin == 0 {
    return false
  }
  for _, pnt := range r.points {
    column := p.grid[pnt.x]
    if pnt.y-1 < len(column) && column[pnt.y-1] {
      return false
    }
  }
  r.offset(0, -1)
  return true
}

func (p *problem) add(r *rock) {
  for _, pnt := range r.points {
    if pnt.y >= len(p.grid[pnt.x]) {
      for pnt.y > len(p.grid[pnt.x]) {
        p.grid[pnt.x] = append(p.grid[pnt.x], false)
      }
      p.grid[pnt.x] = append(p.grid[pnt.x], true)
    } else {
      p.grid[pnt.x][pnt.y] = true
    }
  }
}

func (p *problem) print() {
  fmt.Println("----")
  for _, c := range p.grid {
    for _, p := range c {
      if p {
        fmt.Print("#")
      } else {
        fmt.Print(".")
      }
    }
    fmt.Println()
  }
}

func parseJet(jet string) []int {
  dirs := make([]int, len(jet))
  for i, j := range jet {
    if j == '>' {
      dirs[i] = 1
    } else {
      dirs[i] = -1
    }
  }
  return dirs
}

func partOne(jet string, iters int) int {
  p := newProblem(7)
  c := newCache(10)
  rockIdx := 0
  height := 0
  res := height

  dirs := parseJet(jet)
  dirIdx := 0
  moveToDir := func(r *rock) {
    p.moveSide(r, dirs[dirIdx])
    dirIdx = (dirIdx + 1) % len(dirs)
  }

  for i := 0; i < iters; i++ {
    r := p.rocks[rockIdx].clone()
    r.offset(2, height+3+r.yMin)

    moveToDir(r)
    for p.moveDown(r) {
      moveToDir(r)
    }
    p.add(r)

    rockIdx = (rockIdx + 1) % len(p.rocks)
    if r.yMax+1 > height {
      height = r.yMax + 1
    }
    cycle := c.add(height, rockIdx, dirIdx, i)
    if cycle != -1 {
      fmt.Println(cycle)
    }
  }

  return res
}
