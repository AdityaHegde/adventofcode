package not_enough_minerals

import (
  "testing"

  "AdityaHegde/adventofcode/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 33, partOne(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 1418, partOne(lines))
}

func Test_partTwo(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 3472, partTwo(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 4114, partTwo(lines))
}
