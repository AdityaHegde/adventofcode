package pyroclastic_flow

import (
  "AdityaHegde/adventofcode/go/utils"
)

type cache struct {
  cache      *utils.ThreeDGrid[int]
  lookup     int
  prevHeight int
}

func newCache(lookup int) *cache {
  return &cache{
    cache:      utils.NewThreeDGrid[int](),
    lookup:     lookup,
    prevHeight: 0,
  }
}

func (c *cache) key(p *problem, height, rockIdx, dirIdx int) (int, int, int) {
  gridKey1 := 0
  gridKey2 := 0
  for i := 0; i < 15; i++ {
    for j := 0; j < 7; j++ {
      n := 0
      if len(p.grid[j]) > height-i-1 && p.grid[j][height-i-1] {
        n = 1
      }
      if j < 4 {
        gridKey1 = gridKey1 | (n << (i * j))
      } else {
        gridKey2 = gridKey2 | (n << (i * (j - 4)))
      }
    }
  }
  return gridKey1, gridKey2, rockIdx + (dirIdx << 4)
}

func (c *cache) add(p *problem, height, rockIdx, dirIdx, idx int) int {
  if height < c.lookup {
    return -1
  }

  k1, k2, k3 := c.key(p, height, rockIdx, dirIdx)
  if c.cache.Has(k1, k2, k3) {
    return c.cache.Grid[k1][k2][k3]
  }
  c.cache.Set(k1, k2, k3, idx)
  return -1
}
