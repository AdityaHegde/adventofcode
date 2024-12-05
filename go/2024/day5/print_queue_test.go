package day5

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 143, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 4281, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 123, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 5466, partTwo(utils.InputLinesFromFile("input.txt")))
}
