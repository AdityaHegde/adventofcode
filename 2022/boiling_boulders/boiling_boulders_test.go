package boiling_boulders

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 64, partOne(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 4300, partOne(lines))
}

func Test_partTwo_adhoc(t *testing.T) {
  require.Equal(t, 96, partTwo(generate(4, map[string]bool{
    "1,1,1": true,
    "2,1,1": true,
    "1,2,1": true,
    "1,1,2": true,
  })))

  require.Equal(t, 96, partTwo(generate(4, map[string]bool{
    "1,1,1": true,
    "2,1,1": true,
    "1,2,1": true,
    "1,1,2": true,
    "2,2,1": true,
    "1,2,2": true,
    "2,1,2": true,
    "2,2,2": true,
  })))

  require.Equal(t, 150, partTwo(generate(5, map[string]bool{
    "2,2,2": true,
    "1,2,2": true,
    "3,2,2": true,
    "2,1,2": true,
    "2,3,2": true,
    "2,2,1": true,
    "2,2,3": true,
  })))

  require.Equal(t, 122, partTwo(generate(4, map[string]bool{
    "1,1,1": true,
    "2,1,1": true,
    "1,2,1": true,
    "1,1,2": true,
    "2,2,1": true,
    "1,2,2": true,
    "2,1,2": true,
    "2,2,2": true,
    "2,2,3": true,
  })))

  require.Equal(t, 30, partTwo([]string{
    "0,1,1",
    "2,1,1",
    "1,0,1",
    "1,2,1",
    "1,1,0",
    "1,1,2",
  }))
}

func Test_partTwo(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 58, partTwo(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 2490, partTwo(lines))
}

func generate(size int, exclude map[string]bool) []string {
  lines := make([]string, 0)
  for x := 0; x < size; x++ {
    for y := size - 1; y >= 0; y-- {
      for z := 0; z < size; z++ {
        k := fmt.Sprintf("%d,%d,%d", x, y, z)
        if _, ok := exclude[k]; ok {
          continue
        }
        lines = append(lines, k)
      }
    }
  }
  return lines
}
