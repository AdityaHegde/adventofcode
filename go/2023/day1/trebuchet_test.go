package day1

import (
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  require.Equal(t, 142, partOne(utils.InputLinesFromFile("sample.txt")))
  require.Equal(t, 56397, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
  require.Equal(t, 281, partTwo(utils.InputLinesFromFile("sample2.txt")))
  require.Equal(t, 55701, partTwo(utils.InputLinesFromFile("input.txt")))
}
