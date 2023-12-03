package day2

import (
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  require.Equal(t, 8, partOne(utils.InputLinesFromFile("sample.txt")))
  require.Equal(t, 2207, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
  require.Equal(t, 2286, partTwo(utils.InputLinesFromFile("sample.txt")))
  require.Equal(t, 62241, partTwo(utils.InputLinesFromFile("input.txt")))
}
