package day10

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 4, partOne(utils.InputLinesFromFile("sample1.txt"), SouthEast))
	require.Equal(t, 8, partOne(utils.InputLinesFromFile("sample2.txt"), SouthEast))
	require.Equal(t, 7102, partOne(utils.InputLinesFromFile("input.txt"), NorthEast))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 4, partTwo(utils.InputLinesFromFile("sample3.txt"), SouthEast))
	require.Equal(t, 10, partTwo(utils.InputLinesFromFile("sample4.txt"), SouthWest))
	require.Equal(t, 363, partTwo(utils.InputLinesFromFile("input.txt"), NorthEast))
}
