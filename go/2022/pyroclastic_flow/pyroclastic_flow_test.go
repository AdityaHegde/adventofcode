package pyroclastic_flow

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/go/utils"
  "github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
  jets := utils.InputLinesFromFile("input_0.txt")
  require.Equal(t, 3068, partOne(jets[0], 2022))

  jets = utils.InputLinesFromFile("input_1.txt")
  require.Equal(t, 3124, partOne(jets[0], 2022))
}

func Test_partTwo(t *testing.T) {
  jets := utils.InputLinesFromFile("input_0.txt")
  fmt.Println(partOne(jets[0], 1000))
  //fmt.Println(partOne(jets[0], 1000000000000))

  //jets = utils.InputLinesFromFile("input_1.txt")
  //fmt.Println(partOne(jets[0], 1000000000000))
}
