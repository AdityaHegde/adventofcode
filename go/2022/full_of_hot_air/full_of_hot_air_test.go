package full_of_hot_air

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

var cases = []struct {
  snafu string
  num   int
}{
  {"1", 1},
  {"2", 2},
  {"1=", 3},
  {"1-", 4},
  {"10", 5},
  {"11", 6},
  {"12", 7},
  {"2=", 8},
  {"2-", 9},
  {"20", 10},
  {"21", 11},
  {"22", 12},
  {"1==", 13},
  {"1=-", 14},
  {"1=0", 15},
}

func Test_fromSNAFU(t *testing.T) {
  for _, tt := range cases {
    t.Run(fmt.Sprintf("%s => %d", tt.snafu, tt.num), func(t *testing.T) {
      require.Equal(t, tt.num, fromSNAFU(tt.snafu))
    })
  }
}

func Test_toSNAFU(t *testing.T) {
  for _, tt := range cases {
    t.Run(fmt.Sprintf("%d => %s", tt.num, tt.snafu), func(t *testing.T) {
      require.Equal(t, tt.snafu, toSNAFU(tt.num))
    })
  }
}

func Test_partOne(t *testing.T) {
  lines := utils.InputLinesFromFile("sample.txt")
  require.Equal(t, "2=-1=0", partOne(lines))

  lines = utils.InputLinesFromFile("input.txt")
  require.Equal(t, "2=-0=1-0012-=-2=0=01", partOne(lines))
}
