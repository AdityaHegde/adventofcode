package pyroclastic_flow

import "fmt"

type cache struct {
  cache      map[string]int
  lookup     int
  prevHeight int
}

func newCache(lookup int) *cache {
  return &cache{
    cache:      map[string]int{},
    lookup:     lookup,
    prevHeight: 0,
  }
}

func (c *cache) key(height, rockIdx, dirIdx int) string {
  key := fmt.Sprint(height-c.prevHeight, rockIdx, dirIdx)
  fmt.Println(key)
  c.prevHeight = height
  return key
}

func (c *cache) add(height, rockIdx, dirIdx, idx int) int {
  if height < c.lookup {
    return -1
  }

  key := c.key(height, rockIdx, dirIdx)
  if existing, ok := c.cache[key]; ok {
    return existing
  }
  c.cache[key] = idx
  return -1
}
