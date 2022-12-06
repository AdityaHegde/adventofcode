package main

import (
  "AdityaHegde/adventofcode/utils"
  "fmt"
)

const window = 3

func main() {
  lines := utils.InputLines()

  var prev int64 = 0
  var c int64 = 0

  ints := make([]int64, len(lines))
  for i := 0; i < window; i++ {
    ints[i] = utils.Int(lines[i])
    prev += ints[i]
  }

  for i := window; i < len(lines); i++ {
    ints[i] = utils.Int(lines[i])
    next := prev - ints[i-window] + ints[i]
    if prev < next {
      c++
    }
    next = prev
  }

  fmt.Println(c)
}