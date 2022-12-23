package grove_positioning_system

import (
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 3, partOne(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 27726, partOne(lines))
}

func Test_partTwo(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 1623178306, partTwo(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 4275451658004, partTwo(lines))
}
