package regolith_reservoir

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/go/utils"
)

func Test_partOne(t *testing.T) {
  p := newProblem()
  p.parse(utils.InputLinesFromFile("input_0.txt"))
  fmt.Println(partOne(p))

  p = newProblem()
  p.parse(utils.InputLinesFromFile("input_1.txt"))
  fmt.Println(partOne(p))
}

func Test_partTwo(t *testing.T) {
  p := newProblem()
  p.parse(utils.InputLinesFromFile("input_0.txt"))
  fmt.Println(partTwo(p))

  p = newProblem()
  p.parse(utils.InputLinesFromFile("input_1.txt"))
  fmt.Println(partTwo(p))
}
