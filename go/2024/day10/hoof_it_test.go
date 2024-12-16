package day10

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 36, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 682, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 81, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 1511, partTwo(utils.InputLinesFromFile("input.txt")))
}
