package utils

type TwoDGrid[val interface{}] struct {
  Grid map[int]map[int]val
}

func NewTwoDGrid[val interface{}]() *TwoDGrid[val] {
  return &TwoDGrid[val]{
    Grid: map[int]map[int]val{},
  }
}

func (g *TwoDGrid[val]) Set(x, y int, v val) {
  _, ok := g.Grid[x]
  if !ok {
    g.Grid[x] = make(map[int]val)
  }
  g.Grid[x][y] = v
}

func (g *TwoDGrid[val]) Has(x, y int) bool {
  _, ok := g.Grid[x]
  if !ok {
    return false
  }
  _, ok = g.Grid[x][y]
  return ok
}

func (g *TwoDGrid[val]) Delete(x, y int) {
  if !g.Has(x, y) {
    return
  }
  delete(g.Grid[x], y)
}

func (g *TwoDGrid[val]) ForEach(callback func(x, y int, v val)) {
  for x, xg := range g.Grid {
    for y, v := range xg {
      callback(x, y, v)
    }
  }
}

type ThreeDGrid[val interface{}] struct {
  Grid map[int]map[int]map[int]val
}

func NewThreeDGrid[val interface{}]() *ThreeDGrid[val] {
  return &ThreeDGrid[val]{
    Grid: map[int]map[int]map[int]val{},
  }
}

func (g *ThreeDGrid[val]) Set(x, y, z int, v val) {
  _, ok := g.Grid[x]
  if !ok {
    g.Grid[x] = make(map[int]map[int]val)
  }
  _, ok = g.Grid[x][y]
  if !ok {
    g.Grid[x][y] = make(map[int]val)
  }
  g.Grid[x][y][z] = v
}

func (g *ThreeDGrid[val]) Has(x, y, z int) bool {
  _, ok := g.Grid[x]
  if !ok {
    return false
  }
  _, ok = g.Grid[x][y]
  if !ok {
    return false
  }
  _, ok = g.Grid[x][y][z]
  return ok
}

func (g *ThreeDGrid[val]) Delete(x, y, z int) {
  if !g.Has(x, y, z) {
    return
  }
  delete(g.Grid[x][y], z)
}

func (g *ThreeDGrid[val]) ForEach(callback func(x, y, z int, v val)) {
  for x, xg := range g.Grid {
    for y, yg := range xg {
      for z, v := range yg {
        callback(x, y, z, v)
      }
    }
  }
}
