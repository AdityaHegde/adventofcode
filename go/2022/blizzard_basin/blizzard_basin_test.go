package blizzard_basin

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("sample.txt")
  require.Equal(t, 18, partOne(lines))

  lines = utils.InputLinesFromFile("input.txt")
  require.Equal(t, 242, partOne(lines))
}

func Test_partTwo(t *testing.T) {
  lines := utils.InputLinesFromFile("sample.txt")
  require.Equal(t, 54, partTwo(lines))

  lines = utils.InputLinesFromFile("input.txt")
  fmt.Println(partTwo(lines))
  require.Equal(t, 720, partTwo(lines))
}
