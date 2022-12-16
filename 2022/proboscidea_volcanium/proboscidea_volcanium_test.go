package proboscidea_volcanium

import (
  "testing"

  "AdityaHegde/adventofcode/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  g := newGraph()
  g.parse(lines)
  require.Equal(t, 1651, partOne(g, "AA", 30, 0))

  lines = utils.InputLinesFromFile("input_1.txt")
  g = newGraph()
  g.parse(lines)
  require.Equal(t, 1991, partOne(g, "AA", 30, 0))
}

func Test_partTwo(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  g := newGraph()
  g.parse(lines)
  require.Equal(t, 1707, partTwo(g, "AA", "AA", 26, 26, 0))

  lines = utils.InputLinesFromFile("input_1.txt")
  g = newGraph()
  g.parse(lines)

  // fluctuates to 2357 sometimes. TODO: Why??
  require.Equal(t, 2705, partTwo(g, "AA", "AA", 26, 26, 0))
}
