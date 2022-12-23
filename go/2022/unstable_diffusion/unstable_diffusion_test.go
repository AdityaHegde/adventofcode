package unstable_diffusion

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 110, partOne(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  fmt.Println(partOne(lines))
}

func Test_partTwo(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 20, partTwo(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 862, partTwo(lines))
}
