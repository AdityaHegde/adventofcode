package beacon_exclusion_zone

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  p := newProblem()
  p.parse(lines)
  require.Equal(t, 26, partOne(p, 10))

  lines = utils.InputLinesFromFile("input_1.txt")
  p = newProblem()
  p.parse(lines)
  require.Equal(t, 5461729, partOne(p, 2000000))
}

func Test_partTwo(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  p := newProblem()
  p.parse(lines)
  fmt.Println(partTwo(p))

  lines = utils.InputLinesFromFile("input_1.txt")
  p = newProblem()
  p.parse(lines)
  fmt.Println(partTwo(p))
}
