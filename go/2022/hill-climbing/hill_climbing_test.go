package hill_climbing

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/go/utils"
)

func Test_partOne(t *testing.T) {
  p := parse(utils.InputLinesFromFile("input_1.txt"))
  for i := 0; i < len(p.grid); i++ {
    p.startX = i
    fmt.Println(partOne(p))
  }
}
