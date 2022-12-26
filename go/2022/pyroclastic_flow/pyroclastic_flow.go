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
  rocks      []*rock
  rockIdx    int
  dirs       []int
  dirIdx     int
  grid       [][]bool
  height     int
  prevHeight int
  cache      *cache
}

func newProblem(width int, jet string) *problem {
  dirs := make([]int, len(jet))
  for i, j := range jet {
    if j == '>' {
      dirs[i] = 1
    } else {
      dirs[i] = -1
    }
  }

  p := &problem{
    []*rock{
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
    0,
    dirs,
    0,
    make([][]bool, width),
    0,
    0,
    newCache(15),
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

func (p *problem) moveToDir(r *rock) {
  p.moveSide(r, p.dirs[p.dirIdx])
  p.dirIdx = (p.dirIdx + 1) % len(p.dirs)
}

func (p *problem) cycle(iters int, shouldCache bool) int {
  for i := 0; i < iters; i++ {
    p.prevHeight = p.height
    r := p.rocks[p.rockIdx].clone()
    r.offset(2, p.height+3+r.yMin)

    p.moveToDir(r)
    for p.moveDown(r) {
      p.moveToDir(r)
    }
    p.add(r)

    p.rockIdx = (p.rockIdx + 1) % len(p.rocks)
    if r.yMax+1 > p.height {
      p.height = r.yMax + 1
    }
    if shouldCache {
      cycle := p.cache.add(p, p.height, p.rockIdx, p.dirIdx, i)
      if cycle != -1 {
        return cycle
      }
    }
  }

  return iters
}

func partOne(jet string, iters int) int {
  p := newProblem(7, jet)
  p.cycle(iters, false)
  return p.height
}

func partTwo(jet string, iters int) int {
  p := newProblem(7, jet)
  loop := p.cycle(iters, true)
  height := p.prevHeight * iters / loop
  iters %= loop
  fmt.Println("Loop", loop, p.prevHeight, iters, height)
  p.cycle(iters, false)
  return height + p.height
}
