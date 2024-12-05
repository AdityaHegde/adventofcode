package day4

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 18, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 2378, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 9, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 1796, partTwo(utils.InputLinesFromFile("input.txt")))
}
