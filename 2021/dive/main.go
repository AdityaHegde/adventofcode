package main

import (
  "AdityaHegde/adventofcode/utils"
  "fmt"
  "strings"
)

func main() {
  lines := utils.InputLines()

  var x int64 = 0
  var aim int64 = 0
  var y int64 = 0

  for _, line := range lines {
    splits := strings.Split(line, " ")
    dist := utils.Int(splits[1])
    switch splits[0] {
    case "forward":
      x += dist
      y += dist * aim
    case "up":
      aim -= dist
    case "down":
      aim += dist
    }
  }

  fmt.Println(x * y)
}
