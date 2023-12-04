package day4

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 13, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 18653, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 30, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 5921508, partTwo(utils.InputLinesFromFile("input.txt")))
}
