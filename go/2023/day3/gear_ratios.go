package day3

import (
  "AdityaHegde/adventofcode/go/utils"
)

type partNum struct {
  num       int
  y, bx, ex int
}

type gear struct {
  ct   int
  nums []int
}

func parse(input []string) (*utils.TwoDGrid[bool], *utils.TwoDGrid[*gear], []partNum) {
  symbols := utils.NewTwoDGrid[bool]()
  gears := utils.NewTwoDGrid[*gear]()
  nums := make([]partNum, 0)

  for y, line := range input {
    n := 0
    bx := -1
    addNum := func(x int) {
      nums = append(nums, partNum{
        num: n,
        y:   y,
        bx:  bx,
        ex:  x - 1,
      })
      bx = -1
      n = 0
    }

    for x, c := range line {
      if c == '.' {
        if bx != -1 {
          addNum(x)
        }
        continue
      }
      if c >= '0' && c <= '9' {
        n = n*10 + int(c-'0')
        if bx == -1 {
          bx = x
        }
      } else {
        if bx != -1 {
          addNum(x)
        }
        symbols.Set(x, y, true)
        if c == '*' {
          gears.Set(x, y, &gear{0, []int{}})
        }
      }
    }
    if bx != -1 {
      addNum(len(line))
    }
  }

  return symbols, gears, nums
}

func partOne(input []string) int {
  res := 0
  symbols, _, nums := parse(input)

  for _, num := range nums {
    if symbols.Has(num.bx-1, num.y) || symbols.Has(num.ex+1, num.y) {
      res += num.num
      continue
    }
    for x := num.bx - 1; x <= num.ex+1; x++ {
      if symbols.Has(x, num.y-1) {
        res += num.num
        break
      }
      if symbols.Has(x, num.y+1) {
        res += num.num
        break
      }
    }
  }

  return res
}

func partTwo(input []string) int {
  _, gears, nums := parse(input)
  checkGear := func(x, y int, num partNum) {
    if !gears.Has(x, y) {
      return
    }
    g := gears.Grid[x][y]
    g.ct++
    g.nums = append(g.nums, num.num)
  }

  for _, num := range nums {
    checkGear(num.bx-1, num.y, num)
    checkGear(num.ex+1, num.y, num)
    for x := num.bx - 1; x <= num.ex+1; x++ {
      checkGear(x, num.y-1, num)
      checkGear(x, num.y+1, num)
    }
  }

  res := 0
  gears.ForEach(func(x, y int, g *gear) {
    if g.ct == 2 {
      res += g.nums[0] * g.nums[1]
    }
  })

  return res
}
