package main

import (
  "fmt"
  "math"

  utils2 "AdityaHegde/adventofcode/go/utils"
)

func main() {
  lines := utils2.InputLines()
  fmt.Println(trackRope(lines, 2))
  fmt.Println(trackRope(lines, 10))
}

var DirToOffset = map[uint8]struct {
  x int
  y int
}{
  'L': {-1, 0},
  'U': {0, -1},
  'R': {1, 0},
  'D': {0, 1},
}

type ropeKnot struct {
  x int
  y int
}

func trackRope(lines []string, ropeSize int) int {
  visited := make(map[int]map[int]bool)
  res := 0
  tailVisited := func(x int, y int) {
    inner, ok := visited[x]
    if !ok {
      visited[x] = make(map[int]bool)
      inner = visited[x]
    }

    if inner[y] {
      return
    }
    inner[y] = true
    res++
  }

  rope := utils2.NewLinkedList[*ropeKnot]()
  for i := 0; i < ropeSize; i++ {
    rope.Push(&ropeKnot{
      x: 0,
      y: 0,
    })
  }
  moveHead := func(dir uint8) {
    offsets := DirToOffset[dir]
    knot := rope.Root
    knot.Value.x += offsets.x
    knot.Value.y += offsets.y
    next := knot.Next
    for next != nil {
      dx := knot.Value.x - next.Value.x
      dy := knot.Value.y - next.Value.y
      adx := math.Abs(float64(dx))
      ady := math.Abs(float64(dy))
      if adx > 1 || ady > 1 {
        next.Value.x += utils2.Sign(dx)
        next.Value.y += utils2.Sign(dy)
      } else {
        break
      }
      knot = next
      next = next.Next
    }
    if knot.Next == nil {
      tailVisited(knot.Value.x, knot.Value.y)
    }
  }

  tailVisited(0, 0)

  for _, line := range lines {
    dir := line[0]
    dist := int(utils2.Int64(line[2:]))
    for i := 0; i < dist; i++ {
      moveHead(dir)
    }
  }

  return res
}
