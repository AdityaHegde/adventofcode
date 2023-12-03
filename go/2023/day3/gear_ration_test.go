package day3

import (
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  require.Equal(t, 4361, partOne(utils.InputLinesFromFile("sample.txt")))
  require.Equal(t, 535351, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
  require.Equal(t, 467835, partTwo(utils.InputLinesFromFile("sample.txt")))
  require.Equal(t, 87287096, partTwo(utils.InputLinesFromFile("input.txt")))
}
