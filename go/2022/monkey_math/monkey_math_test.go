package monkey_math

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/go/utils"
)

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  fmt.Println(partOne(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  fmt.Println(partOne(lines))
}

func Test_partTwo(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  fmt.Println(partTwo(lines))

  // not 9780869100463 or 9780869100464
  // is 3327575724809
  // TODO: why does solving not work?
  lines = utils.InputLinesFromFile("input_1.txt")
  fmt.Println(partTwo(lines))
}
