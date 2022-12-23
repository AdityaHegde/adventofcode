package main

import (
  "fmt"

  utils2 "AdityaHegde/adventofcode/go/utils"
)

func main() {
  lines := utils2.InputLines()
  grid := parseGrid(lines)
  fmt.Println(partOne(grid))
  fmt.Println(partTwo(grid))
}

func parseGrid(lines []string) [][]int {
  grid := make([][]int, len(lines))
  for i, line := range lines {
    grid[i] = make([]int, len(line))

    for j, n := range line {
      grid[i][j] = int(n - '0')
    }
  }
  return grid
}

func partOne(grid [][]int) int {
  gridLen := len(grid[0])
  gridLenOff := gridLen - 1
  visible := make([][]bool, len(grid))
  for i := 0; i < gridLen; i++ {
    visible[i] = make([]bool, gridLen)
  }
  res := 0

  isVisible := func(i int, j int) {
    if visible[i][j] {
      return
    }
    visible[i][j] = true
    res++
  }

  for i := 0; i < gridLen; i++ {
    left := -1
    right := -1
    top := -1
    bottom := -1

    for j := 0; j < gridLen; j++ {
      if left < grid[i][j] {
        left = grid[i][j]
        isVisible(i, j)
      }
      if right < grid[i][gridLenOff-j] {
        right = grid[i][gridLenOff-j]
        isVisible(i, gridLenOff-j)
      }
      if top < grid[j][i] {
        top = grid[j][i]
        isVisible(j, i)
      }
      if bottom < grid[gridLenOff-j][i] {
        bottom = grid[gridLenOff-j][i]
        isVisible(gridLenOff-j, i)
      }
    }
  }

  return res
}

type tree struct {
  height int
  index  int
}

func partTwo(grid [][]int) int {
  gridLen := len(grid[0])
  gridLenOff := gridLen - 1
  score := make([][]int, len(grid))
  for i := 0; i < gridLen; i++ {
    score[i] = make([]int, gridLen)
    for j := 0; j < gridLen; j++ {
      score[i][j] = -1
    }
  }
  res := 0

  addToScore := func(i int, j int, s int) {
    if score[i][j] == -1 {
      score[i][j] = s
    } else {
      score[i][j] *= s
    }
    fmt.Println(i, j, s, score[i][j], res)
    if score[i][j] > res {
      res = score[i][j]
    }
  }

  for i := 0; i < gridLen; i++ {
    left := utils2.NewLinkedList[tree]()
    right := utils2.NewLinkedList[tree]()
    top := utils2.NewLinkedList[tree]()
    bottom := utils2.NewLinkedList[tree]()

    loopThroughList := func(
      list *utils2.LinkedList[tree],
      height int,
      index int,
      i int,
      j int,
    ) {
      node := list.Tail
      for node != nil && node.Value.height < height {
        list.Pop()
        node = list.Tail
      }
      s := index
      if node != nil {
        s -= node.Value.index
      }
      if node != nil && node.Value.height == height {
        node.Value.index = index
      } else {
        list.Push(tree{
          height: height,
          index:  index,
        })
      }
      addToScore(i, j, s)
    }

    for j := 0; j < gridLen; j++ {
      loopThroughList(left, grid[i][j], j, i, j)
      loopThroughList(right, grid[i][gridLenOff-j], j, i, gridLenOff-j)
      loopThroughList(top, grid[j][i], j, j, i)
      loopThroughList(bottom, grid[gridLenOff-j][i], j, gridLenOff-j, i)
    }
  }

  return res
}
