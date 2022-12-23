package monkey_map

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

func Test_turn(t *testing.T) {
  cases := []struct {
    dx   int
    dy   int
    dir  int
    turn int
    edx  int
    edy  int
  }{
    {1, 0, 0, 1, 0, 1},
    {0, 1, 1, 1, -1, 0},
    {-1, 0, 2, 1, 0, -1},
    {0, -1, 3, 1, 1, 0},
    {1, 0, 0, -1, 0, -1},
    {0, -1, 3, -1, -1, 0},
    {-1, 0, 2, -1, 0, 1},
    {0, 1, 1, -1, 1, 0},
  }

  for _, tt := range cases {
    t.Run(fmt.Sprintf("%d,%d =%d> %d,%d", tt.dx, tt.dy, tt.turn, tt.edx, tt.edy), func(t *testing.T) {
      s := solution{0, 0, tt.dx, tt.dy, tt.dir}
      s.turn(tt.turn)
      require.Equal(t, s.dx, tt.edx)
      require.Equal(t, s.dy, tt.edy)
    })
  }
}

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 6032, partOne(lines))

  lines = utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 189140, partOne(lines))
}

func Test_partTwo(t *testing.T) {
  // only works for this input. needs manual mapping
  lines := utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 115063, partTwo(lines))
}
