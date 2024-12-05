package day1

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 11, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 1189304, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 31, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 24349736, partTwo(utils.InputLinesFromFile("input.txt")))
}
