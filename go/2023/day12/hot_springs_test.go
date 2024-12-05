package day12

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 21, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 7191, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 525152, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 6512849198636, partTwo(utils.InputLinesFromFile("input.txt")))
}
